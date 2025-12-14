package processor

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
)

// RiskManager 风控管理器
type RiskManager struct {
	maxBetAmount      *big.Int // 单笔最大下注金额
	minBetAmount      *big.Int // 单笔最小下注金额
	maxTotalBets      int      // 最大下注次数
	maxConcurrentBets int      // 最大并发下注数
	minBalance        *big.Int // 最小余额保护
	stopLoss          *big.Int // 止损金额
	dailyLossLimit    *big.Int // 每日亏损限制

	// 统计数据
	totalBets      int       // 总下注次数
	dailyLoss      *big.Int  // 当日亏损
	dailyResetTime time.Time // 每日重置时间
}

// NewRiskManager 创建风控管理器
func NewRiskManager(
	maxBetAmount string,
	minBetAmount string,
	maxTotalBets int,
	maxConcurrentBets int,
	minBalance string,
	stopLoss string,
	dailyLossLimit string,
) (*RiskManager, error) {
	maxBet, ok := new(big.Int).SetString(maxBetAmount, 10)
	if !ok {
		return nil, fmt.Errorf("invalid maxBetAmount: %s", maxBetAmount)
	}

	minBet, ok := new(big.Int).SetString(minBetAmount, 10)
	if !ok {
		return nil, fmt.Errorf("invalid minBetAmount: %s", minBetAmount)
	}

	minBal, ok := new(big.Int).SetString(minBalance, 10)
	if !ok {
		return nil, fmt.Errorf("invalid minBalance: %s", minBalance)
	}

	stopLossVal, ok := new(big.Int).SetString(stopLoss, 10)
	if !ok {
		return nil, fmt.Errorf("invalid stopLoss: %s", stopLoss)
	}

	dailyLoss, ok := new(big.Int).SetString(dailyLossLimit, 10)
	if !ok {
		return nil, fmt.Errorf("invalid dailyLossLimit: %s", dailyLossLimit)
	}

	rm := &RiskManager{
		maxBetAmount:      maxBet,
		minBetAmount:      minBet,
		maxTotalBets:      maxTotalBets,
		maxConcurrentBets: maxConcurrentBets,
		minBalance:        minBal,
		stopLoss:          stopLossVal,
		dailyLossLimit:    dailyLoss,
		totalBets:         0,
		dailyLoss:         big.NewInt(0),
		dailyResetTime:    getNextDayStart(),
	}

	log.Printf("📊 Risk Control Initialized:")
	log.Printf("   Max Bet Amount: %s wei", maxBet.String())
	log.Printf("   Min Bet Amount: %s wei", minBet.String())
	log.Printf("   Max Total Bets: %d (0=unlimited)", maxTotalBets)
	log.Printf("   Max Concurrent Bets: %d (0=unlimited)", maxConcurrentBets)
	log.Printf("   Min Balance: %s wei", minBal.String())
	log.Printf("   Stop Loss: %s wei (0=disabled)", stopLossVal.String())
	log.Printf("   Daily Loss Limit: %s wei (0=disabled)", dailyLoss.String())

	return rm, nil
}

// AdjustBetAmount 调整下注金额并检查风控条件
// 返回: (调整后的金额, 是否可以下注, 拒绝原因)
func (rm *RiskManager) AdjustBetAmount(p *PancakeStrategy, betAmount *big.Int) (*big.Int, bool, string) {
	// 检查每日重置
	rm.checkDailyReset()

	// 1. 调整下注金额到合法范围
	adjustedAmount := new(big.Int).Set(betAmount)

	// 如果超过最大值，调整为最大值
	if adjustedAmount.Cmp(rm.maxBetAmount) > 0 {
		log.Printf("⚠️ Bet amount %s exceeds max %s, adjusted to max",
			betAmount.String(), rm.maxBetAmount.String())
		adjustedAmount = new(big.Int).Set(rm.maxBetAmount)
	}

	// 如果低于最小值，调整为最小值
	if adjustedAmount.Cmp(rm.minBetAmount) < 0 {
		log.Printf("⚠️ Bet amount %s below min %s, adjusted to min",
			betAmount.String(), rm.minBetAmount.String())
		adjustedAmount = new(big.Int).Set(rm.minBetAmount)
	}

	// 2. 检查总下注次数
	if rm.maxTotalBets > 0 && rm.totalBets >= rm.maxTotalBets {
		reason := fmt.Sprintf("Total bets %d reached max %d", rm.totalBets, rm.maxTotalBets)
		log.Printf("🚫 Risk Control: %s", reason)
		return nil, false, reason
	}

	// 3. 检查并发下注数
	concurrentBets := p.getConcurrentBets()
	if rm.maxConcurrentBets > 0 && concurrentBets >= rm.maxConcurrentBets {
		reason := fmt.Sprintf("Concurrent bets %d reached max %d", concurrentBets, rm.maxConcurrentBets)
		log.Printf("🚫 Risk Control: %s", reason)
		return nil, false, reason
	}

	// 4. 检查余额
	balance, err := p.getBalance()
	if err != nil {
		reason := fmt.Sprintf("Failed to get balance: %v", err)
		log.Printf("⚠️ Risk Control: %s", reason)
		return nil, false, reason
	}

	if balance.Cmp(rm.minBalance) < 0 {
		reason := fmt.Sprintf("Balance %s below minimum %s", balance.String(), rm.minBalance.String())
		log.Printf("🚫 Risk Control: %s", reason)
		return nil, false, reason
	}

	// 确保余额足够下注 + 保留最小余额（使用调整后的金额）
	required := new(big.Int).Add(adjustedAmount, rm.minBalance)
	if balance.Cmp(required) < 0 {
		reason := fmt.Sprintf("Insufficient balance: have %s, need %s (bet=%s + minBalance=%s)",
			balance.String(), required.String(), adjustedAmount.String(), rm.minBalance.String())
		log.Printf("🚫 Risk Control: %s", reason)
		return nil, false, reason
	}

	// 5. 检查止损
	if rm.stopLoss.Cmp(big.NewInt(0)) > 0 {
		totalLoss := p.getTotalLoss()
		if totalLoss.Cmp(rm.stopLoss) >= 0 {
			reason := fmt.Sprintf("Stop loss triggered: loss %s >= limit %s", totalLoss.String(), rm.stopLoss.String())
			log.Printf("🚫 Risk Control: %s", reason)
			return nil, false, reason
		}
	}

	// 6. 检查每日亏损限制
	if rm.dailyLossLimit.Cmp(big.NewInt(0)) > 0 {
		if rm.dailyLoss.Cmp(rm.dailyLossLimit) >= 0 {
			reason := fmt.Sprintf("Daily loss limit reached: %s >= %s", rm.dailyLoss.String(), rm.dailyLossLimit.String())
			log.Printf("🚫 Risk Control: %s", reason)
			return nil, false, reason
		}
	}

	// 通过所有检查，返回调整后的金额
	if adjustedAmount.Cmp(betAmount) != 0 {
		log.Printf("✅ Bet amount adjusted: %s -> %s", betAmount.String(), adjustedAmount.String())
	}

	return adjustedAmount, true, ""
}

// OnBetPlaced 下注成功后调用
func (rm *RiskManager) OnBetPlaced(betAmount *big.Int) {
	rm.totalBets++
	log.Printf("📈 Risk Stats: Total Bets=%d, Daily Loss=%s", rm.totalBets, rm.dailyLoss.String())
}

// OnBetResult 下注结果后调用（更新盈亏）
func (rm *RiskManager) OnBetResult(won bool, betAmount *big.Int, payout *big.Int) {
	if won {
		// 盈利 = payout - betAmount
		profit := new(big.Int).Sub(payout, betAmount)
		log.Printf("✅ Bet Won: betAmount=%s, payout=%s, profit=%s", betAmount.String(), payout.String(), profit.String())

		// 盈利可以减少当日亏损
		rm.dailyLoss = new(big.Int).Sub(rm.dailyLoss, profit)
		if rm.dailyLoss.Cmp(big.NewInt(0)) < 0 {
			rm.dailyLoss = big.NewInt(0)
		}
	} else {
		// 亏损 = betAmount
		log.Printf("❌ Bet Lost: betAmount=%s", betAmount.String())
		rm.dailyLoss = new(big.Int).Add(rm.dailyLoss, betAmount)
	}

	log.Printf("📊 Risk Stats: Total Bets=%d, Daily Loss=%s", rm.totalBets, rm.dailyLoss.String())
}

// checkDailyReset 检查是否需要每日重置
func (rm *RiskManager) checkDailyReset() {
	now := time.Now()
	if now.After(rm.dailyResetTime) {
		log.Printf("🔄 Daily risk stats reset")
		rm.dailyLoss = big.NewInt(0)
		rm.dailyResetTime = getNextDayStart()
	}
}

// getNextDayStart 获取明天0点的时间
func getNextDayStart() time.Time {
	now := time.Now()
	tomorrow := now.AddDate(0, 0, 1)
	return time.Date(tomorrow.Year(), tomorrow.Month(), tomorrow.Day(), 0, 0, 0, 0, tomorrow.Location())
}

// getConcurrentBets 获取当前并发下注数（已下注但未结束的轮次）
func (p *PancakeStrategy) getConcurrentBets() int {
	p.mu.RLock()
	defer p.mu.RUnlock()

	count := 0
	for _, round := range p.rounds {
		if round.HasBet && !round.RoundEnded {
			count++
		}
	}
	return count
}

// getBalance 获取钱包余额
func (p *PancakeStrategy) getBalance() (*big.Int, error) {
	// 尝试从任意 RPC 获取余额
	for _, rpcURL := range p.rpcList {
		client, err := ethclient.Dial(rpcURL)
		if err != nil {
			continue
		}
		defer client.Close()

		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		balance, err := client.BalanceAt(ctx, p.myAddress, nil)
		cancel()

		if err == nil {
			return balance, nil
		}
	}

	return nil, fmt.Errorf("failed to get balance from any RPC")
}

// getTotalLoss 获取总亏损（所有已结束的轮次）
func (p *PancakeStrategy) getTotalLoss() *big.Int {
	p.mu.RLock()
	defer p.mu.RUnlock()

	totalLoss := big.NewInt(0)

	for _, round := range p.rounds {
		if round.HasBet && round.RoundEnded {
			// 判断是否输了
			won := p.riskManager.didWinForRisk(round)
			if !won {
				// 输了，累加亏损
				totalLoss = new(big.Int).Add(totalLoss, round.BetAmount)
			}
		}
	}

	return totalLoss
}

// didWinForRisk 判断是否赢了（风控模块使用）
func (rm *RiskManager) didWinForRisk(round *RoundState) bool {
	if round.LockPrice == nil || round.ClosePrice == nil {
		return false
	}

	if round.BetSide == "Bull" {
		return round.ClosePrice.Cmp(round.LockPrice) > 0
	} else {
		return round.ClosePrice.Cmp(round.LockPrice) < 0
	}
}

// GetStats 获取风控统计信息
func (rm *RiskManager) GetStats() map[string]interface{} {
	return map[string]interface{}{
		"totalBets":      rm.totalBets,
		"dailyLoss":      rm.dailyLoss.String(),
		"dailyResetTime": rm.dailyResetTime.Format("2006-01-02 15:04:05"),
	}
}
