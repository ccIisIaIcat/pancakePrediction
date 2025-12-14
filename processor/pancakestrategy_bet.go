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

// checkBetOpportunities 检查当前活跃轮次的下注时机
// 注意：调用此方法前必须已持有锁
func (p *PancakeStrategy) checkBetOpportunities(currentBlock uint64) {
	// 只检查当前 epoch（最新轮次）
	if p.currentEpoch == 0 {
		return // 还没有收到任何 StartRound
	}

	// 获取当前 epoch 的状态
	round, exists := p.rounds[p.currentEpoch]
	if !exists {
		// 当前 epoch 不存在，说明还没收到 StartRound
		return
	}

	// 检查是否到了下注时机
	if p.shouldBet(round, currentBlock) {
		// 执行下注（构造和签名交易）
		p.executeBet(round, currentBlock)
	}
}

// shouldBet 判断是否应该下注
func (p *PancakeStrategy) shouldBet(round *RoundState, currentBlock uint64) bool {
	// 条件1: StartRound 已收到
	if round.StartBlock == 0 {
		return false
	}

	// 条件2: 还没锁定
	if round.RoundLocked {
		return false
	}

	// 条件3: ratio >= minRatio
	if round.Ratio < p.config.MinRatio {
		return false
	}

	// 条件4: 区块时机判断
	// 下注窗口: startBlock + (blocksPerRound - triggerBlockDiff) ~ startBlock + blocksPerRound
	// 例如: startBlock + 501 ~ startBlock + 508
	minBetBlock := round.StartBlock + p.config.BlocksPerRound - p.config.TriggerBlockDiff
	maxBetBlock := round.StartBlock + p.config.BlocksPerRound

	if currentBlock < minBetBlock {
		return false // 还没到时机
	}

	if currentBlock >= maxBetBlock {
		return false // 已经太晚了（应该已经锁定）
	}

	// 条件5: MinoritySide 必须确定
	if round.MinoritySide == "" {
		return false
	}

	return true
}

// executeBet 执行下注（构造和签名交易）
func (p *PancakeStrategy) executeBet(round *RoundState, currentBlock uint64) {
	// 检查是否已经下注过
	if round.HasBet {
		log.Printf("⚠️ Already bet on epoch %d, skipping", round.Epoch)
		return
	}

	// 计算下注金额
	minorityAmount := round.BullAmount
	if round.MinoritySide == "Bear" {
		minorityAmount = round.BearAmount
	}

	// 下注金额 = k × 少数方池子
	betAmountFloat := new(big.Float).SetInt(minorityAmount)
	kFloat := big.NewFloat(p.config.KFactor)
	betAmountFloat.Mul(betAmountFloat, kFloat)

	betAmount, _ := betAmountFloat.Int(nil)

	log.Printf("💰 BET OPPORTUNITY: epoch=%d, block=%d, side=%s, ratio=%.2f, calculatedAmount=%s",
		round.Epoch, currentBlock, round.MinoritySide, round.Ratio, betAmount.String())

	// 风控检查和金额调整
	adjustedAmount, canBet, reason := p.riskManager.AdjustBetAmount(p, betAmount)
	if !canBet {
		log.Printf("🚫 Bet blocked by risk control: %s", reason)

		p.logger.Warn("Bet Blocked by Risk Control",
			zap.Uint64("epoch", round.Epoch),
			zap.String("reason", reason),
			zap.String("calculatedAmount", betAmount.String()))

		return
	}

	// 使用调整后的金额
	betAmount = adjustedAmount
	log.Printf("✅ Final bet amount after risk control: %s", betAmount.String())

	// 1. 发送邮件：通过风控判断，准备下注
	calculatedBetAmount := new(big.Int).Set(betAmount)
	if adjustedAmount.Cmp(betAmount) != 0 {
		calculatedBetAmount, _ = betAmountFloat.Int(nil) // 原始计算金额
	}
	p.notifyBetOpportunity(round.Epoch, round.MinoritySide, round.Ratio, calculatedBetAmount, betAmount, currentBlock)

	// 构造交易
	signedTx, err := p.buildAndSignBetTx(round.Epoch, round.MinoritySide, betAmount)
	if err != nil {
		log.Printf("❌ Failed to build bet tx: %v", err)
		return
	}

	txHash := signedTx.Hash().Hex()

	// 记录到 logger
	p.logger.Info("Bet Transaction Signed",
		zap.Uint64("epoch", round.Epoch),
		zap.String("side", round.MinoritySide),
		zap.String("betAmount", betAmount.String()),
		zap.Float64("ratio", round.Ratio),
		zap.Uint64("currentBlock", currentBlock),
		zap.String("txHash", txHash))

	// 发送交易到所有 RPC 节点
	success := p.sendBetTransaction(signedTx)
	if !success {
		log.Printf("❌ Failed to send bet transaction for epoch %d", round.Epoch)
		return
	}

	// 标记已下注
	round.HasBet = true
	round.BetTxHash = txHash
	round.BetSide = round.MinoritySide
	round.BetAmount = betAmount
	round.BetConfirmed = false

	// 递增 nonce
	p.incrementNonce()

	// 通知风控管理器
	p.riskManager.OnBetPlaced(betAmount)

	log.Printf("✅ Bet transaction sent: epoch=%d, txHash=%s", round.Epoch, txHash)

	// 2. 发送邮件：交易已发送
	p.notifyBetSent(round.Epoch, round.MinoritySide, betAmount, txHash)

	// 启动交易确认追踪
	go p.trackTransaction(round.Epoch, txHash)
}

// buildAndSignBetTx 构造并签名下注交易
func (p *PancakeStrategy) buildAndSignBetTx(epoch uint64, side string, betAmount *big.Int) (*ethtypes.Transaction, error) {
	// 构造交易数据
	var data []byte
	var err error

	if side == "Bull" {
		// betBull(uint256 epoch)
		data, err = p.contractABI.Pack("betBull", new(big.Int).SetUint64(epoch))
	} else {
		// betBear(uint256 epoch)
		data, err = p.contractABI.Pack("betBear", new(big.Int).SetUint64(epoch))
	}

	if err != nil {
		return nil, fmt.Errorf("failed to pack data: %w", err)
	}

	// 使用缓存的 nonce
	nonce := p.getCurrentNonce()

	tx := ethtypes.NewTransaction(
		nonce,
		common.HexToAddress(p.config.ContractAddress),
		betAmount,                     // value
		p.config.GasLimitBet,          // gas limit
		big.NewInt(p.config.GasPrice), // gas price
		data,
	)

	log.Printf("📝 Building tx with nonce=%d", nonce)

	// 签名交易
	chainID := big.NewInt(56) // BSC mainnet
	signedTx, err := ethtypes.SignTx(tx, ethtypes.NewEIP155Signer(chainID), p.privateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to sign tx: %w", err)
	}

	return signedTx, nil
}

// sendBetTransaction 并发发送交易到所有 RPC 节点
func (p *PancakeStrategy) sendBetTransaction(signedTx *ethtypes.Transaction) bool {
	// 获取所有 RPC URL
	if len(p.rpcList) == 0 {
		log.Printf("❌ No RPC endpoints configured")
		return false
	}

	// 使用 WaitGroup 和 channel 来并发发送
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
				log.Printf("⚠️ RPC #%d [%s] send failed: %v", index, url, err)
				successChan <- false
				return
			}

			log.Printf("✅ RPC #%d [%s] sent successfully", index, url)
			successChan <- true
		}(i, rpcURL)
	}

	// 等待所有 goroutine 完成
	wg.Wait()
	close(successChan)

	// 只要有一个成功就算成功
	for success := range successChan {
		if success {
			return true
		}
	}

	return false
}

// trackTransaction 追踪交易确认状态
func (p *PancakeStrategy) trackTransaction(epoch uint64, txHash string) {
	log.Printf("🔍 Starting to track transaction: epoch=%d, txHash=%s", epoch, txHash)

	// 最多追踪 5 分钟
	timeout := time.After(5 * time.Minute)
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-timeout:
			log.Printf("⏰ Transaction tracking timeout: epoch=%d, txHash=%s", epoch, txHash)
			return

		case <-ticker.C:
			// 尝试从任意 RPC 节点获取交易收据
			receipt, err := p.getTransactionReceipt(txHash)
			if err != nil {
				// 还没上链，继续等待
				continue
			}

			// 检查交易状态
			if receipt.Status == 1 {
				log.Printf("✅ Transaction confirmed: epoch=%d, txHash=%s, blockNumber=%d",
					epoch, txHash, receipt.BlockNumber.Uint64())

				// 更新 RoundState
				p.mu.Lock()
				if round, exists := p.rounds[epoch]; exists {
					round.BetConfirmed = true

					p.logger.Info("Bet Transaction Confirmed",
						zap.Uint64("epoch", epoch),
						zap.String("txHash", txHash),
						zap.Uint64("blockNumber", receipt.BlockNumber.Uint64()),
						zap.Uint64("gasUsed", receipt.GasUsed))
				}
				p.mu.Unlock()

				// 3. 发送邮件：交易确认成功
				p.notifyBetConfirmed(epoch, txHash, receipt.BlockNumber.Uint64(), true)

				return

			} else {
				log.Printf("❌ Transaction failed: epoch=%d, txHash=%s", epoch, txHash)

				p.logger.Error("Bet Transaction Failed",
					zap.Uint64("epoch", epoch),
					zap.String("txHash", txHash))

				// 3. 发送邮件：交易确认失败
				p.notifyBetConfirmed(epoch, txHash, receipt.BlockNumber.Uint64(), false)

				// 交易失败，刷新 nonce
				go p.refreshNonce()

				return
			}
		}
	}
}

// getTransactionReceipt 从任意可用的 RPC 节点获取交易收据
func (p *PancakeStrategy) getTransactionReceipt(txHash string) (*ethtypes.Receipt, error) {
	for _, rpcURL := range p.rpcList {
		client, err := ethclient.Dial(rpcURL)
		if err != nil {
			continue
		}
		defer client.Close()

		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		receipt, err := client.TransactionReceipt(ctx, common.HexToHash(txHash))
		cancel()

		if err == nil {
			return receipt, nil
		}
	}

	return nil, fmt.Errorf("no receipt found from any RPC")
}
