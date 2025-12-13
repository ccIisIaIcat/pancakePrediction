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
}
