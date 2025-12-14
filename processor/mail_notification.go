package processor

import (
	"fmt"
	"log"
	"math/big"
	"time"
)

// weiToBNB 将 wei 转换为 BNB（保留4位小数）
func weiToBNB(wei *big.Int) string {
	if wei == nil {
		return "0.0000"
	}

	// 1 BNB = 10^18 wei
	bnbFloat := new(big.Float).SetInt(wei)
	divisor := new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil))
	bnbFloat.Quo(bnbFloat, divisor)

	return fmt.Sprintf("%.4f BNB", bnbFloat)
}

// sendMailAsync 异步发送邮件（不阻塞）
func (p *PancakeStrategy) sendMailAsync(subject, body string) {
	if p.mailSender == nil || len(p.mailTo) == 0 {
		return // 邮件未配置，跳过
	}

	go func() {
		err := p.mailSender.SendMailToMultipleWithTLS(p.mailTo, subject, body)
		if err != nil {
			log.Printf("⚠️ Failed to send email '%s': %v", subject, err)
		} else {
			log.Printf("📧 Email sent: %s", subject)
		}
	}()
}

// notifyBetOpportunity 通知发现下注机会
func (p *PancakeStrategy) notifyBetOpportunity(epoch uint64, side string, ratio float64, calculatedAmount, finalAmount *big.Int, currentBlock uint64) {
	subject := fmt.Sprintf("💰 下注机会 - Epoch %d", epoch)
	body := fmt.Sprintf(`
		<html>
		<body style="font-family: Arial, sans-serif; padding: 20px; background-color: #f5f5f5;">
			<div style="max-width: 600px; margin: 0 auto; background-color: white; padding: 20px; border-radius: 10px; box-shadow: 0 2px 4px rgba(0,0,0,0.1);">
				<h2 style="color: #FF9800; margin-top: 0;">💰 发现下注机会</h2>

				<table style="width: 100%%; border-collapse: collapse; margin: 20px 0;">
					<tr style="background-color: #f9f9f9;">
						<td style="padding: 12px; border: 1px solid #ddd; font-weight: bold;">Epoch</td>
						<td style="padding: 12px; border: 1px solid #ddd;">%d</td>
					</tr>
					<tr>
						<td style="padding: 12px; border: 1px solid #ddd; font-weight: bold;">下注方向</td>
						<td style="padding: 12px; border: 1px solid #ddd; color: %s; font-weight: bold;">%s</td>
					</tr>
					<tr style="background-color: #f9f9f9;">
						<td style="padding: 12px; border: 1px solid #ddd; font-weight: bold;">Ratio</td>
						<td style="padding: 12px; border: 1px solid #ddd;">%.2f</td>
					</tr>
					<tr>
						<td style="padding: 12px; border: 1px solid #ddd; font-weight: bold;">计算金额</td>
						<td style="padding: 12px; border: 1px solid #ddd;">%s</td>
					</tr>
					<tr style="background-color: #f9f9f9;">
						<td style="padding: 12px; border: 1px solid #ddd; font-weight: bold;">实际下注金额</td>
						<td style="padding: 12px; border: 1px solid #ddd; font-weight: bold; color: #4CAF50;">%s</td>
					</tr>
					<tr>
						<td style="padding: 12px; border: 1px solid #ddd; font-weight: bold;">当前区块</td>
						<td style="padding: 12px; border: 1px solid #ddd;">%d</td>
					</tr>
					<tr style="background-color: #f9f9f9;">
						<td style="padding: 12px; border: 1px solid #ddd; font-weight: bold;">时间</td>
						<td style="padding: 12px; border: 1px solid #ddd;">%s</td>
					</tr>
				</table>

				<p style="margin: 20px 0; padding: 15px; background-color: #fff3cd; border-left: 4px solid #FF9800; border-radius: 4px;">
					<strong>状态:</strong> 准备发送交易...
				</p>

				<hr style="border: none; border-top: 1px solid #ddd; margin: 20px 0;">
				<p style="color: #666; font-size: 12px; text-align: center;">
					此邮件由 Pancake Strategy Bot 自动发送
				</p>
			</div>
		</body>
		</html>
	`, epoch, getBetColorHTML(side), side, ratio, weiToBNB(calculatedAmount), weiToBNB(finalAmount), currentBlock, time.Now().Format("2006-01-02 15:04:05"))

	p.sendMailAsync(subject, body)
}

// notifyBetSent 通知交易已发送
func (p *PancakeStrategy) notifyBetSent(epoch uint64, side string, betAmount *big.Int, txHash string) {
	subject := fmt.Sprintf("✅ 交易已发送 - Epoch %d", epoch)
	body := fmt.Sprintf(`
		<html>
		<body style="font-family: Arial, sans-serif; padding: 20px; background-color: #f5f5f5;">
			<div style="max-width: 600px; margin: 0 auto; background-color: white; padding: 20px; border-radius: 10px; box-shadow: 0 2px 4px rgba(0,0,0,0.1);">
				<h2 style="color: #4CAF50; margin-top: 0;">✅ 下注交易已发送</h2>

				<table style="width: 100%%; border-collapse: collapse; margin: 20px 0;">
					<tr style="background-color: #f9f9f9;">
						<td style="padding: 12px; border: 1px solid #ddd; font-weight: bold;">Epoch</td>
						<td style="padding: 12px; border: 1px solid #ddd;">%d</td>
					</tr>
					<tr>
						<td style="padding: 12px; border: 1px solid #ddd; font-weight: bold;">下注方向</td>
						<td style="padding: 12px; border: 1px solid #ddd; color: %s; font-weight: bold;">%s</td>
					</tr>
					<tr style="background-color: #f9f9f9;">
						<td style="padding: 12px; border: 1px solid #ddd; font-weight: bold;">下注金额</td>
						<td style="padding: 12px; border: 1px solid #ddd; font-weight: bold; color: #4CAF50;">%s</td>
					</tr>
					<tr>
						<td style="padding: 12px; border: 1px solid #ddd; font-weight: bold;">交易哈希</td>
						<td style="padding: 12px; border: 1px solid #ddd; font-family: monospace; font-size: 11px; word-break: break-all;">%s</td>
					</tr>
					<tr style="background-color: #f9f9f9;">
						<td style="padding: 12px; border: 1px solid #ddd; font-weight: bold;">时间</td>
						<td style="padding: 12px; border: 1px solid #ddd;">%s</td>
					</tr>
				</table>

				<p style="margin: 20px 0; padding: 15px; background-color: #e8f5e9; border-left: 4px solid #4CAF50; border-radius: 4px;">
					<strong>状态:</strong> 交易已发送到 RPC 节点，等待区块确认中...
				</p>

				<hr style="border: none; border-top: 1px solid #ddd; margin: 20px 0;">
				<p style="color: #666; font-size: 12px; text-align: center;">
					此邮件由 Pancake Strategy Bot 自动发送
				</p>
			</div>
		</body>
		</html>
	`, epoch, getBetColorHTML(side), side, weiToBNB(betAmount), txHash, time.Now().Format("2006-01-02 15:04:05"))

	p.sendMailAsync(subject, body)
}

// notifyBetConfirmed 通知交易已确认
func (p *PancakeStrategy) notifyBetConfirmed(epoch uint64, txHash string, blockNumber uint64, success bool) {
	var statusColor, statusIcon, statusText string
	if success {
		statusColor = "#4CAF50"
		statusIcon = "✅"
		statusText = "交易确认成功"
	} else {
		statusColor = "#F44336"
		statusIcon = "❌"
		statusText = "交易执行失败"
	}

	subject := fmt.Sprintf("%s 交易确认 - Epoch %d", statusIcon, epoch)
	body := fmt.Sprintf(`
		<html>
		<body style="font-family: Arial, sans-serif; padding: 20px; background-color: #f5f5f5;">
			<div style="max-width: 600px; margin: 0 auto; background-color: white; padding: 20px; border-radius: 10px; box-shadow: 0 2px 4px rgba(0,0,0,0.1);">
				<h2 style="color: %s; margin-top: 0;">%s %s</h2>

				<table style="width: 100%%; border-collapse: collapse; margin: 20px 0;">
					<tr style="background-color: #f9f9f9;">
						<td style="padding: 12px; border: 1px solid #ddd; font-weight: bold;">Epoch</td>
						<td style="padding: 12px; border: 1px solid #ddd;">%d</td>
					</tr>
					<tr>
						<td style="padding: 12px; border: 1px solid #ddd; font-weight: bold;">交易哈希</td>
						<td style="padding: 12px; border: 1px solid #ddd; font-family: monospace; font-size: 11px; word-break: break-all;">%s</td>
					</tr>
					<tr style="background-color: #f9f9f9;">
						<td style="padding: 12px; border: 1px solid #ddd; font-weight: bold;">确认区块</td>
						<td style="padding: 12px; border: 1px solid #ddd;">%d</td>
					</tr>
					<tr>
						<td style="padding: 12px; border: 1px solid #ddd; font-weight: bold;">时间</td>
						<td style="padding: 12px; border: 1px solid #ddd;">%s</td>
					</tr>
				</table>

				<hr style="border: none; border-top: 1px solid #ddd; margin: 20px 0;">
				<p style="color: #666; font-size: 12px; text-align: center;">
					此邮件由 Pancake Strategy Bot 自动发送
				</p>
			</div>
		</body>
		</html>
	`, statusColor, statusIcon, statusText, epoch, txHash, blockNumber, time.Now().Format("2006-01-02 15:04:05"))

	p.sendMailAsync(subject, body)
}

// notifyRoundResult 通知轮次结果
func (p *PancakeStrategy) notifyRoundResult(round *RoundState, won bool) {
	var resultColor, resultIcon, resultText string
	if won {
		resultColor = "#4CAF50"
		resultIcon = "🎉"
		resultText = "赢了！"
	} else {
		resultColor = "#F44336"
		resultIcon = "😞"
		resultText = "输了"
	}

	subject := fmt.Sprintf("%s 轮次结束 - Epoch %d (%s)", resultIcon, round.Epoch, resultText)
	body := fmt.Sprintf(`
		<html>
		<body style="font-family: Arial, sans-serif; padding: 20px; background-color: #f5f5f5;">
			<div style="max-width: 600px; margin: 0 auto; background-color: white; padding: 20px; border-radius: 10px; box-shadow: 0 2px 4px rgba(0,0,0,0.1);">
				<h2 style="color: %s; margin-top: 0;">%s 轮次结束 - %s</h2>

				<table style="width: 100%%; border-collapse: collapse; margin: 20px 0;">
					<tr style="background-color: #f9f9f9;">
						<td style="padding: 12px; border: 1px solid #ddd; font-weight: bold;">Epoch</td>
						<td style="padding: 12px; border: 1px solid #ddd;">%d</td>
					</tr>
					<tr>
						<td style="padding: 12px; border: 1px solid #ddd; font-weight: bold;">下注方向</td>
						<td style="padding: 12px; border: 1px solid #ddd; color: %s; font-weight: bold;">%s</td>
					</tr>
					<tr style="background-color: #f9f9f9;">
						<td style="padding: 12px; border: 1px solid #ddd; font-weight: bold;">下注金额</td>
						<td style="padding: 12px; border: 1px solid #ddd;">%s</td>
					</tr>
					<tr>
						<td style="padding: 12px; border: 1px solid #ddd; font-weight: bold;">锁定价格</td>
						<td style="padding: 12px; border: 1px solid #ddd;">%s</td>
					</tr>
					<tr style="background-color: #f9f9f9;">
						<td style="padding: 12px; border: 1px solid #ddd; font-weight: bold;">结算价格</td>
						<td style="padding: 12px; border: 1px solid #ddd;">%s</td>
					</tr>
					<tr>
						<td style="padding: 12px; border: 1px solid #ddd; font-weight: bold;">结果</td>
						<td style="padding: 12px; border: 1px solid #ddd; font-weight: bold; color: %s; font-size: 18px;">%s</td>
					</tr>
					<tr style="background-color: #f9f9f9;">
						<td style="padding: 12px; border: 1px solid #ddd; font-weight: bold;">时间</td>
						<td style="padding: 12px; border: 1px solid #ddd;">%s</td>
					</tr>
				</table>

				<hr style="border: none; border-top: 1px solid #ddd; margin: 20px 0;">
				<p style="color: #666; font-size: 12px; text-align: center;">
					此邮件由 Pancake Strategy Bot 自动发送
				</p>
			</div>
		</body>
		</html>
	`, resultColor, resultIcon, resultText, round.Epoch, getBetColorHTML(round.BetSide), round.BetSide,
		weiToBNB(round.BetAmount), round.LockPrice.String(), round.ClosePrice.String(),
		resultColor, resultText, time.Now().Format("2006-01-02 15:04:05"))

	p.sendMailAsync(subject, body)
}

// notifyClaimSent 通知 Claim 交易已发送
func (p *PancakeStrategy) notifyClaimSent(epoch uint64, txHash string) {
	subject := fmt.Sprintf("💰 Claim 交易已发送 - Epoch %d", epoch)
	body := fmt.Sprintf(`
		<html>
		<body style="font-family: Arial, sans-serif; padding: 20px; background-color: #f5f5f5;">
			<div style="max-width: 600px; margin: 0 auto; background-color: white; padding: 20px; border-radius: 10px; box-shadow: 0 2px 4px rgba(0,0,0,0.1);">
				<h2 style="color: #4CAF50; margin-top: 0;">💰 Claim 交易已发送</h2>

				<table style="width: 100%%; border-collapse: collapse; margin: 20px 0;">
					<tr style="background-color: #f9f9f9;">
						<td style="padding: 12px; border: 1px solid #ddd; font-weight: bold;">Epoch</td>
						<td style="padding: 12px; border: 1px solid #ddd;">%d</td>
					</tr>
					<tr>
						<td style="padding: 12px; border: 1px solid #ddd; font-weight: bold;">交易哈希</td>
						<td style="padding: 12px; border: 1px solid #ddd; font-family: monospace; font-size: 11px; word-break: break-all;">%s</td>
					</tr>
					<tr style="background-color: #f9f9f9;">
						<td style="padding: 12px; border: 1px solid #ddd; font-weight: bold;">时间</td>
						<td style="padding: 12px; border: 1px solid #ddd;">%s</td>
					</tr>
				</table>

				<p style="margin: 20px 0; padding: 15px; background-color: #e8f5e9; border-left: 4px solid #4CAF50; border-radius: 4px;">
					<strong>状态:</strong> Claim 交易已发送，等待区块确认中...
				</p>

				<hr style="border: none; border-top: 1px solid #ddd; margin: 20px 0;">
				<p style="color: #666; font-size: 12px; text-align: center;">
					此邮件由 Pancake Strategy Bot 自动发送
				</p>
			</div>
		</body>
		</html>
	`, epoch, txHash, time.Now().Format("2006-01-02 15:04:05"))

	p.sendMailAsync(subject, body)
}

// notifyClaimConfirmed 通知 Claim 交易已确认
func (p *PancakeStrategy) notifyClaimConfirmed(epoch uint64, txHash string, blockNumber uint64, success bool) {
	var statusColor, statusIcon, statusText string
	if success {
		statusColor = "#4CAF50"
		statusIcon = "✅"
		statusText = "Claim 成功"
	} else {
		statusColor = "#F44336"
		statusIcon = "❌"
		statusText = "Claim 失败"
	}

	subject := fmt.Sprintf("%s Claim 确认 - Epoch %d", statusIcon, epoch)
	body := fmt.Sprintf(`
		<html>
		<body style="font-family: Arial, sans-serif; padding: 20px; background-color: #f5f5f5;">
			<div style="max-width: 600px; margin: 0 auto; background-color: white; padding: 20px; border-radius: 10px; box-shadow: 0 2px 4px rgba(0,0,0,0.1);">
				<h2 style="color: %s; margin-top: 0;">%s %s</h2>

				<table style="width: 100%%; border-collapse: collapse; margin: 20px 0;">
					<tr style="background-color: #f9f9f9;">
						<td style="padding: 12px; border: 1px solid #ddd; font-weight: bold;">Epoch</td>
						<td style="padding: 12px; border: 1px solid #ddd;">%d</td>
					</tr>
					<tr>
						<td style="padding: 12px; border: 1px solid #ddd; font-weight: bold;">交易哈希</td>
						<td style="padding: 12px; border: 1px solid #ddd; font-family: monospace; font-size: 11px; word-break: break-all;">%s</td>
					</tr>
					<tr style="background-color: #f9f9f9;">
						<td style="padding: 12px; border: 1px solid #ddd; font-weight: bold;">确认区块</td>
						<td style="padding: 12px; border: 1px solid #ddd;">%d</td>
					</tr>
					<tr>
						<td style="padding: 12px; border: 1px solid #ddd; font-weight: bold;">时间</td>
						<td style="padding: 12px; border: 1px solid #ddd;">%s</td>
					</tr>
				</table>

				<hr style="border: none; border-top: 1px solid #ddd; margin: 20px 0;">
				<p style="color: #666; font-size: 12px; text-align: center;">
					此邮件由 Pancake Strategy Bot 自动发送
				</p>
			</div>
		</body>
		</html>
	`, statusColor, statusIcon, statusText, epoch, txHash, blockNumber, time.Now().Format("2006-01-02 15:04:05"))

	p.sendMailAsync(subject, body)
}

// getBetColorHTML 根据下注方向返回 HTML 颜色
func getBetColorHTML(side string) string {
	if side == "Bull" {
		return "#4CAF50" // 绿色
	}
	return "#F44336" // 红色
}
