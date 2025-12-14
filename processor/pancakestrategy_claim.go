package processor

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
)

// checkAllRoundsForClaim 检查所有缓存的轮次，尝试 claim 还没 claim 的
// 注意：调用此方法前必须已持有锁
func (p *PancakeStrategy) checkAllRoundsForClaim() {
	claimedCount := 0

	// 遍历所有轮次
	for epoch, round := range p.rounds {
		// 检查是否可以 claim
		if round.RoundEnded && round.HasBet && round.BetConfirmed && !round.HasClaimed {
			log.Printf("🔍 Checking claim for epoch %d...", epoch)
			p.checkAndClaim(round)
			claimedCount++
		}
	}

	if claimedCount > 0 {
		log.Printf("✅ Checked %d rounds for claiming", claimedCount)
	}
}

// checkAndClaim 检查是否需要 claim 并执行
// 注意：调用此方法前必须已持有锁
func (p *PancakeStrategy) checkAndClaim(round *RoundState) {
	// 条件1: 已下注
	if !round.HasBet {
		return
	}

	// 条件2: 交易已确认
	if !round.BetConfirmed {
		log.Printf("⚠️ Bet not confirmed yet for epoch %d, cannot claim", round.Epoch)
		return
	}

	// 条件3: 还没 claim 过
	if round.HasClaimed {
		log.Printf("⚠️ Already claimed for epoch %d", round.Epoch)
		return
	}

	// 条件4: 价格数据完整
	if round.LockPrice == nil || round.ClosePrice == nil {
		log.Printf("⚠️ Missing price data for epoch %d", round.Epoch)
		return
	}

	if round.LockPrice.Cmp(big.NewInt(0)) == 0 || round.ClosePrice.Cmp(big.NewInt(0)) == 0 {
		log.Printf("⚠️ Invalid price data for epoch %d", round.Epoch)
		return
	}

	// 判断是否赢了
	won := p.didWin(round)
	if !won {
		log.Printf("❌ Lost bet on epoch %d: betSide=%s, lockPrice=%s, closePrice=%s",
			round.Epoch, round.BetSide, round.LockPrice.String(), round.ClosePrice.String())

		p.logger.Info("Bet Lost",
			zap.Uint64("epoch", round.Epoch),
			zap.String("betSide", round.BetSide),
			zap.String("lockPrice", round.LockPrice.String()),
			zap.String("closePrice", round.ClosePrice.String()))

		return
	}

	log.Printf("🎉 Won bet on epoch %d! betSide=%s, lockPrice=%s, closePrice=%s",
		round.Epoch, round.BetSide, round.LockPrice.String(), round.ClosePrice.String())

	// 4. 发送邮件：轮次结束，赢了
	p.notifyRoundResult(round, true)

	// 执行 claim
	p.executeClaim(round)
}

// checkAndNotifyLoss 检查是否输了并发送通知
func (p *PancakeStrategy) checkAndNotifyLoss(round *RoundState) {
	// 条件：已下注，已确认，轮次已结束，还没通知过
	if !round.HasBet || !round.BetConfirmed || !round.RoundEnded {
		return
	}

	// 判断是否输了
	won := p.didWin(round)
	if !won {
		// 4. 发送邮件：轮次结束，输了
		p.notifyRoundResult(round, false)
	}
}

// didWin 判断是否赢了
func (p *PancakeStrategy) didWin(round *RoundState) bool {
	// Bull: closePrice > lockPrice
	// Bear: closePrice < lockPrice
	// 相等的话是平局，按照合约逻辑应该也算输

	if round.BetSide == "Bull" {
		return round.ClosePrice.Cmp(round.LockPrice) > 0
	} else { // Bear
		return round.ClosePrice.Cmp(round.LockPrice) < 0
	}
}

// executeClaim 执行 claim
func (p *PancakeStrategy) executeClaim(round *RoundState) {
	log.Printf("💰 Executing claim for epoch %d", round.Epoch)

	// 构造 claim 交易
	signedTx, err := p.buildAndSignClaimTx(round.Epoch)
	if err != nil {
		log.Printf("❌ Failed to build claim tx for epoch %d: %v", round.Epoch, err)
		return
	}

	txHash := signedTx.Hash().Hex()

	// 记录到 logger
	p.logger.Info("Claim Transaction Signed",
		zap.Uint64("epoch", round.Epoch),
		zap.String("betSide", round.BetSide),
		zap.String("betAmount", round.BetAmount.String()),
		zap.String("txHash", txHash))

	log.Printf("✅ Signed Claim Tx: %s", txHash)

	// 发送交易到所有 RPC 节点
	success := p.sendClaimTransaction(signedTx)
	if !success {
		log.Printf("❌ Failed to send claim transaction for epoch %d", round.Epoch)
		return
	}

	// 标记已 claim
	round.HasClaimed = true

	// 递增 nonce
	p.incrementNonce()

	log.Printf("✅ Claim transaction sent: epoch=%d, txHash=%s", round.Epoch, txHash)

	// 5. 发送邮件：Claim 交易已发送
	p.notifyClaimSent(round.Epoch, txHash)

	// 启动交易确认追踪
	go p.trackClaimTransaction(round.Epoch, txHash)
}

// buildAndSignClaimTx 构造并签名 claim 交易
func (p *PancakeStrategy) buildAndSignClaimTx(epoch uint64) (*ethtypes.Transaction, error) {
	// 构造交易数据: claim(uint256[] epochs)
	epochs := []*big.Int{new(big.Int).SetUint64(epoch)}
	data, err := p.contractABI.Pack("claim", epochs)
	if err != nil {
		return nil, fmt.Errorf("failed to pack data: %w", err)
	}

	// 使用缓存的 nonce
	nonce := p.getCurrentNonce()

	tx := ethtypes.NewTransaction(
		nonce,
		common.HexToAddress(p.config.ContractAddress),
		big.NewInt(0),                   // value = 0 (claim 不需要发送 BNB)
		p.config.GasLimitClaim,          // gas limit
		big.NewInt(p.config.GasPrice),   // gas price
		data,
	)

	log.Printf("📝 Building claim tx with nonce=%d", nonce)

	// 签名交易
	chainID := big.NewInt(56) // BSC mainnet
	signedTx, err := ethtypes.SignTx(tx, ethtypes.NewEIP155Signer(chainID), p.privateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to sign tx: %w", err)
	}

	return signedTx, nil
}

// sendClaimTransaction 并发发送 claim 交易到所有 RPC 节点
func (p *PancakeStrategy) sendClaimTransaction(signedTx *ethtypes.Transaction) bool {
	if len(p.rpcList) == 0 {
		log.Printf("❌ No RPC endpoints configured")
		return false
	}

	var wg sync.WaitGroup
	successChan := make(chan bool, len(p.rpcList))

	for i, rpcURL := range p.rpcList {
		wg.Add(1)
		go func(index int, url string) {
			defer wg.Done()

			client, err := ethclient.Dial(url)
			if err != nil {
				log.Printf("⚠️ RPC #%d [%s] dial failed: %v", index, url, err)
				successChan <- false
				return
			}
			defer client.Close()

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			err = client.SendTransaction(ctx, signedTx)
			if err != nil {
				log.Printf("⚠️ RPC #%d [%s] send claim failed: %v", index, url, err)
				successChan <- false
				return
			}

			log.Printf("✅ RPC #%d [%s] claim sent successfully", index, url)
			successChan <- true
		}(i, rpcURL)
	}

	wg.Wait()
	close(successChan)

	for success := range successChan {
		if success {
			return true
		}
	}

	return false
}

// trackClaimTransaction 追踪 claim 交易确认状态
func (p *PancakeStrategy) trackClaimTransaction(epoch uint64, txHash string) {
	log.Printf("🔍 Starting to track claim transaction: epoch=%d, txHash=%s", epoch, txHash)

	timeout := time.After(5 * time.Minute)
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-timeout:
			log.Printf("⏰ Claim transaction tracking timeout: epoch=%d, txHash=%s", epoch, txHash)
			return

		case <-ticker.C:
			receipt, err := p.getTransactionReceipt(txHash)
			if err != nil {
				continue
			}

			if receipt.Status == 1 {
				log.Printf("✅ Claim confirmed: epoch=%d, txHash=%s, blockNumber=%d",
					epoch, txHash, receipt.BlockNumber.Uint64())

				p.logger.Info("Claim Transaction Confirmed",
					zap.Uint64("epoch", epoch),
					zap.String("txHash", txHash),
					zap.Uint64("blockNumber", receipt.BlockNumber.Uint64()),
					zap.Uint64("gasUsed", receipt.GasUsed))

				// 6. 发送邮件：Claim 交易确认成功
				p.notifyClaimConfirmed(epoch, txHash, receipt.BlockNumber.Uint64(), true)

				return

			} else {
				log.Printf("❌ Claim transaction failed: epoch=%d, txHash=%s", epoch, txHash)

				p.logger.Error("Claim Transaction Failed",
					zap.Uint64("epoch", epoch),
					zap.String("txHash", txHash))

				// 6. 发送邮件：Claim 交易确认失败
				p.notifyClaimConfirmed(epoch, txHash, receipt.BlockNumber.Uint64(), false)

				// 交易失败，刷新 nonce
				go p.refreshNonce()

				return
			}
		}
	}
}
