package processor

import (
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
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

	log.Printf("💰 BET OPPORTUNITY: epoch=%d, block=%d, side=%s, ratio=%.2f, betAmount=%s",
		round.Epoch, currentBlock, round.MinoritySide, round.Ratio, betAmount.String())

	// 构造交易
	signedTx, err := p.buildAndSignBetTx(round.Epoch, round.MinoritySide, betAmount)
	if err != nil {
		log.Printf("❌ Failed to build bet tx: %v", err)
		return
	}

	// 记录到 logger
	p.logger.Info("Bet Transaction Signed",
		zap.Uint64("epoch", round.Epoch),
		zap.String("side", round.MinoritySide),
		zap.String("betAmount", betAmount.String()),
		zap.Float64("ratio", round.Ratio),
		zap.Uint64("currentBlock", currentBlock),
		zap.String("txHash", signedTx.Hash().Hex()))

	log.Printf("✅ Signed Tx: %s (NOT SENT YET)", signedTx.Hash().Hex())

	// TODO: 这里暂时不发送交易，只打印
	// 后续需要：
	// 1. 发送交易到所有 RPC 节点
	// 2. 跟踪交易状态
	// 3. 标记该 epoch 已下注，避免重复下注
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
		betAmount,                       // value
		p.config.GasLimitBet,            // gas limit
		big.NewInt(p.config.GasPrice),   // gas price
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
