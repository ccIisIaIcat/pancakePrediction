package processor

import (
	"log"
	"math/big"
	"strconv"
	"time"

	"github.com/ccIisIaIcat/pancakePrediction/common/types"
	"go.uber.org/zap"
)

// handleStartRound 处理 StartRound 事件
// 事件数据: epoch
func (p *PancakeStrategy) handleStartRound(logResult *types.LogResult, eventData map[string]interface{}) {
	// 提取 epoch
	epochBigInt, ok := eventData["epoch"].(*big.Int)
	if !ok {
		log.Printf("⚠️ handleStartRound: invalid epoch type")
		return
	}
	epoch := epochBigInt.Uint64()

	// 提取区块号（十六进制字符串转uint64）
	blockNumHex := logResult.BlockNumber
	blockNum, err := strconv.ParseUint(blockNumHex[2:], 16, 64) // 去掉 "0x"
	if err != nil {
		log.Printf("⚠️ handleStartRound: failed to parse block number: %v", err)
		return
	}

	p.mu.Lock()
	defer p.mu.Unlock()

	// 创建新的轮次状态
	p.rounds[epoch] = &RoundState{
		Epoch:        epoch,
		StartBlock:   blockNum,
		BullAmount:   big.NewInt(0),
		BearAmount:   big.NewInt(0),
		Ratio:        0,
		MinoritySide: "",
		CreatedAt:    time.Now(),
		RoundLocked:  false,
		RoundEnded:   false,
		HasBet:       false,
		BetTxHash:    "",
		BetSide:      "",
		BetAmount:    big.NewInt(0),
		BetConfirmed: false,
		HasClaimed:   false,
		LockPrice:    big.NewInt(0),
		ClosePrice:   big.NewInt(0),
	}

	// 更新当前 epoch
	if epoch > p.currentEpoch {
		p.currentEpoch = epoch
	}

	// 刷新 nonce（每轮开始时同步一次，保险起见）
	go func() {
		if err := p.refreshNonce(); err != nil {
			log.Printf("⚠️ Failed to refresh nonce in StartRound: %v", err)
		}
	}()

	// 记录到 zap logger
	p.logger.Info("RoundState Updated",
		zap.String("action", "StartRound"),
		zap.Uint64("epoch", epoch),
		zap.Uint64("startBlock", blockNum),
		zap.String("bullAmount", "0"),
		zap.String("bearAmount", "0"),
		zap.Float64("ratio", 0),
		zap.String("minoritySide", ""))

	log.Printf("🎯 StartRound: epoch=%d, block=%d", epoch, blockNum)
}

// handleBetBull 处理 BetBull 事件
// 事件数据: sender, epoch, amount
func (p *PancakeStrategy) handleBetBull(logResult *types.LogResult, eventData map[string]interface{}) {
	// 提取 epoch
	epochBigInt, ok := eventData["epoch"].(*big.Int)
	if !ok {
		log.Printf("⚠️ handleBetBull: invalid epoch type")
		return
	}
	epoch := epochBigInt.Uint64()

	// 提取 amount
	amount, ok := eventData["amount"].(*big.Int)
	if !ok {
		log.Printf("⚠️ handleBetBull: invalid amount type")
		return
	}

	p.mu.Lock()
	defer p.mu.Unlock()

	// 查找对应轮次
	round, exists := p.rounds[epoch]
	if !exists {
		// 如果 StartRound 还没收到，忽略这个 Bet 事件（防止中途启动导致数据不完整）
		log.Printf("⚠️ handleBetBull: round %d not found (StartRound not received yet), ignoring", epoch)
		return
	}

	// 累加 BullAmount
	round.BullAmount = new(big.Int).Add(round.BullAmount, amount)

	// 重新计算 ratio 和 MinoritySide
	p.updateRatio(round)

	// 记录到 zap logger
	p.logger.Info("RoundState Updated",
		zap.String("action", "BetBull"),
		zap.Uint64("epoch", epoch),
		zap.String("betAmount", amount.String()),
		zap.String("bullAmount", round.BullAmount.String()),
		zap.String("bearAmount", round.BearAmount.String()),
		zap.Float64("ratio", round.Ratio),
		zap.String("minoritySide", round.MinoritySide))

	log.Printf("🐂 BetBull: epoch=%d, amount=%s, totalBull=%s, ratio=%.2f, minority=%s",
		epoch, amount.String(), round.BullAmount.String(), round.Ratio, round.MinoritySide)
}

// handleBetBear 处理 BetBear 事件
// 事件数据: sender, epoch, amount
func (p *PancakeStrategy) handleBetBear(logResult *types.LogResult, eventData map[string]interface{}) {
	// 提取 epoch
	epochBigInt, ok := eventData["epoch"].(*big.Int)
	if !ok {
		log.Printf("⚠️ handleBetBear: invalid epoch type")
		return
	}
	epoch := epochBigInt.Uint64()

	// 提取 amount
	amount, ok := eventData["amount"].(*big.Int)
	if !ok {
		log.Printf("⚠️ handleBetBear: invalid amount type")
		return
	}

	p.mu.Lock()
	defer p.mu.Unlock()

	// 查找对应轮次
	round, exists := p.rounds[epoch]
	if !exists {
		// 如果 StartRound 还没收到，忽略这个 Bet 事件（防止中途启动导致数据不完整）
		log.Printf("⚠️ handleBetBear: round %d not found (StartRound not received yet), ignoring", epoch)
		return
	}

	// 累加 BearAmount
	round.BearAmount = new(big.Int).Add(round.BearAmount, amount)

	// 重新计算 ratio 和 MinoritySide
	p.updateRatio(round)

	// 记录到 zap logger
	p.logger.Info("RoundState Updated",
		zap.String("action", "BetBear"),
		zap.Uint64("epoch", epoch),
		zap.String("betAmount", amount.String()),
		zap.String("bullAmount", round.BullAmount.String()),
		zap.String("bearAmount", round.BearAmount.String()),
		zap.Float64("ratio", round.Ratio),
		zap.String("minoritySide", round.MinoritySide))

	log.Printf("🐻 BetBear: epoch=%d, amount=%s, totalBear=%s, ratio=%.2f, minority=%s",
		epoch, amount.String(), round.BearAmount.String(), round.Ratio, round.MinoritySide)
}

// handleLockRound 处理 LockRound 事件
// 事件数据: epoch, price, roundId
func (p *PancakeStrategy) handleLockRound(logResult *types.LogResult, eventData map[string]interface{}) {
	// 提取 epoch
	epochBigInt, ok := eventData["epoch"].(*big.Int)
	if !ok {
		return
	}
	epoch := epochBigInt.Uint64()

	// 提取 price
	price, ok := eventData["price"].(*big.Int)
	if !ok {
		log.Printf("⚠️ handleLockRound: invalid price type")
		price = big.NewInt(0)
	}

	p.mu.Lock()
	defer p.mu.Unlock()

	// 查找对应轮次
	round, exists := p.rounds[epoch]
	if !exists {
		log.Printf("⚠️ handleLockRound: round %d not found", epoch)
		return
	}

	// 标记轮次已锁定，不能再下注
	round.RoundLocked = true
	round.LockPrice = price

	// 记录到 zap logger
	p.logger.Info("RoundState Updated",
		zap.String("action", "LockRound"),
		zap.Uint64("epoch", epoch),
		zap.String("lockPrice", price.String()),
		zap.String("bullAmount", round.BullAmount.String()),
		zap.String("bearAmount", round.BearAmount.String()),
		zap.Float64("ratio", round.Ratio),
		zap.String("minoritySide", round.MinoritySide),
		zap.Bool("roundLocked", true))

	log.Printf("🔒 LockRound: epoch=%d (locked, no more bets), lockPrice=%s", epoch, price.String())
}

// handleEndRound 处理 EndRound 事件
// 事件数据: epoch, price, roundId
func (p *PancakeStrategy) handleEndRound(logResult *types.LogResult, eventData map[string]interface{}) {
	// 提取 epoch
	epochBigInt, ok := eventData["epoch"].(*big.Int)
	if !ok {
		log.Printf("⚠️ handleEndRound: invalid epoch type")
		return
	}
	epoch := epochBigInt.Uint64()

	// 提取 price (closePrice)
	price, ok := eventData["price"].(*big.Int)
	if !ok {
		log.Printf("⚠️ handleEndRound: invalid price type")
		price = big.NewInt(0)
	}

	p.mu.Lock()
	defer p.mu.Unlock()

	// 查找对应轮次
	round, exists := p.rounds[epoch]
	if !exists {
		log.Printf("⚠️ handleEndRound: round %d not found", epoch)
		return
	}

	// 标记轮次结束
	round.RoundEnded = true
	round.ClosePrice = price

	// 记录到 zap logger
	p.logger.Info("RoundState Updated",
		zap.String("action", "EndRound"),
		zap.Uint64("epoch", epoch),
		zap.String("lockPrice", round.LockPrice.String()),
		zap.String("closePrice", price.String()),
		zap.String("bullAmount", round.BullAmount.String()),
		zap.String("bearAmount", round.BearAmount.String()),
		zap.Float64("ratio", round.Ratio),
		zap.String("minoritySide", round.MinoritySide),
		zap.Bool("roundEnded", true))

	log.Printf("🏁 EndRound: epoch=%d, bull=%s, bear=%s, ratio=%.2f, lockPrice=%s, closePrice=%s",
		epoch, round.BullAmount.String(), round.BearAmount.String(), round.Ratio, round.LockPrice.String(), price.String())

	// 检查所有缓存的轮次，尝试 claim 还没 claim 的
	p.checkAllRoundsForClaim()
}

// handleClaim 处理 Claim 事件
// 事件数据: sender, epoch, amount
func (p *PancakeStrategy) handleClaim(logResult *types.LogResult, eventData map[string]interface{}) {
	// 暂时只记录日志
	epochBigInt, ok := eventData["epoch"].(*big.Int)
	if !ok {
		return
	}
	epoch := epochBigInt.Uint64()

	amount, ok := eventData["amount"].(*big.Int)
	if !ok {
		return
	}

	log.Printf("💰 Claim: epoch=%d, amount=%s", epoch, amount.String())
}

// updateRatio 更新轮次的 ratio 和 MinoritySide
// 注意：调用此方法前必须已持有锁
func (p *PancakeStrategy) updateRatio(round *RoundState) {
	// 如果任意一方为0，ratio无意义
	if round.BullAmount.Cmp(big.NewInt(0)) == 0 || round.BearAmount.Cmp(big.NewInt(0)) == 0 {
		round.Ratio = 0
		round.MinoritySide = ""
		return
	}

	// 转换为 float64 计算 ratio
	bullFloat := new(big.Float).SetInt(round.BullAmount)
	bearFloat := new(big.Float).SetInt(round.BearAmount)

	// ratio = max / min
	if round.BullAmount.Cmp(round.BearAmount) > 0 {
		// Bull > Bear, Bear 是少数方
		ratio := new(big.Float).Quo(bullFloat, bearFloat)
		round.Ratio, _ = ratio.Float64()
		round.MinoritySide = "Bear"
	} else {
		// Bear >= Bull, Bull 是少数方
		ratio := new(big.Float).Quo(bearFloat, bullFloat)
		round.Ratio, _ = ratio.Float64()
		round.MinoritySide = "Bull"
	}
}
