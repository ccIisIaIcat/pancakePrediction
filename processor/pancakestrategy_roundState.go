package processor

import (
	"math/big"
	"time"
)

// RoundState 轮次状态
type RoundState struct {
	Epoch        uint64    // 轮次号
	StartBlock   uint64    // StartRound 事件的区块号
	BullAmount   *big.Int  // Bull池总金额
	BearAmount   *big.Int  // Bear池总金额
	Ratio        float64   // 赔率 max/min
	MinoritySide string    // 少数方 "Bull" or "Bear"
	CreatedAt    time.Time // 创建时间(用于清理)
	RoundLocked  bool      // 轮次是否已锁定(LockRound收到,不能再下注)
	RoundEnded   bool      // 轮次是否已结束(EndRound收到,可以claim)

	// 下注跟踪
	HasBet       bool      // 是否已下注
	BetTxHash    string    // 下注交易哈希
	BetSide      string    // 下注方向 "Bull" or "Bear"
	BetAmount    *big.Int  // 下注金额
	BetConfirmed bool      // 交易是否已确认
	HasClaimed   bool      // 是否已领取奖励

	// 价格信息（用于判断输赢）
	LockPrice  *big.Int // LockRound 时的价格
	ClosePrice *big.Int // EndRound 时的价格
}
