package mail

// 使用示例

/*
// 示例1: 发送普通邮件（适用于大部分 SMTP 服务器）
func ExampleSendMail() {
	sender := NewMailSender(
		"smtp.example.com",    // SMTP 服务器
		587,                   // SMTP 端口
		"your@email.com",      // 发件人邮箱
		"your-password",       // 发件人密码
		"Pancake Bot",         // 发件人名称（可选）
	)

	err := sender.SendMail(
		"recipient@example.com",
		"Test Email",
		"<h1>Hello!</h1><p>This is a test email.</p>",
	)

	if err != nil {
		log.Printf("Failed to send email: %v", err)
	}
}

// 示例2: 发送邮件到多个收件人
func ExampleSendMailToMultiple() {
	sender := NewMailSender(
		"smtp.gmail.com",
		587,
		"your@gmail.com",
		"your-app-password",   // Gmail 需要使用应用专用密码
		"Pancake Strategy",
	)

	recipients := []string{
		"user1@example.com",
		"user2@example.com",
	}

	err := sender.SendMailToMultiple(
		recipients,
		"Strategy Alert",
		"<h2>Bet Opportunity Found!</h2><p>Details...</p>",
	)

	if err != nil {
		log.Printf("Failed to send email: %v", err)
	}
}

// 示例3: 使用 TLS 发送邮件（适用于 Gmail, Outlook 等）
func ExampleSendMailWithTLS() {
	sender := NewMailSender(
		"smtp.gmail.com",
		465,                   // Gmail SSL 端口
		"your@gmail.com",
		"your-app-password",
		"Pancake Bot",
	)

	err := sender.SendMailWithTLS(
		"recipient@example.com",
		"Bet Result",
		`
		<html>
		<body>
			<h1>Bet Result</h1>
			<p>Epoch: 12345</p>
			<p>Result: Won</p>
			<p>Amount: 1.5 BNB</p>
		</body>
		</html>
		`,
	)

	if err != nil {
		log.Printf("Failed to send email: %v", err)
	}
}

// 示例4: 发送纯文本邮件
func ExampleSendPlainText() {
	sender := NewMailSender(
		"smtp.example.com",
		587,
		"your@email.com",
		"your-password",
		"",
	)

	err := sender.SendPlainText(
		"recipient@example.com",
		"Plain Text Alert",
		"This is a plain text email without HTML formatting.",
	)

	if err != nil {
		log.Printf("Failed to send email: %v", err)
	}
}

// 示例5: 在策略中使用（发送下注通知）
func ExampleIntegrationWithStrategy() {
	// 创建邮件发送器
	mailSender := NewMailSender(
		"smtp.gmail.com",
		587,
		"bot@gmail.com",
		"app-password",
		"Pancake Strategy Bot",
	)

	// 下注成功后发送通知
	epoch := uint64(12345)
	betAmount := "1.5 BNB"
	ratio := 2.5

	subject := fmt.Sprintf("Bet Placed - Epoch %d", epoch)
	body := fmt.Sprintf(`
		<html>
		<body style="font-family: Arial, sans-serif;">
			<h2 style="color: #4CAF50;">Bet Placed Successfully</h2>
			<table style="border-collapse: collapse; width: 100%%;">
				<tr>
					<td style="padding: 8px; border: 1px solid #ddd;"><strong>Epoch:</strong></td>
					<td style="padding: 8px; border: 1px solid #ddd;">%d</td>
				</tr>
				<tr>
					<td style="padding: 8px; border: 1px solid #ddd;"><strong>Amount:</strong></td>
					<td style="padding: 8px; border: 1px solid #ddd;">%s</td>
				</tr>
				<tr>
					<td style="padding: 8px; border: 1px solid #ddd;"><strong>Ratio:</strong></td>
					<td style="padding: 8px; border: 1px solid #ddd;">%.2f</td>
				</tr>
			</table>
		</body>
		</html>
	`, epoch, betAmount, ratio)

	err := mailSender.SendMail("your@email.com", subject, body)
	if err != nil {
		log.Printf("Failed to send bet notification: %v", err)
	}
}

// 常见 SMTP 服务器配置:
//
// Gmail:
//   SMTP: smtp.gmail.com
//   Port: 587 (TLS) 或 465 (SSL)
//   需要开启"应用专用密码": https://myaccount.google.com/apppasswords
//
// Outlook/Hotmail:
//   SMTP: smtp-mail.outlook.com
//   Port: 587
//
// Yahoo:
//   SMTP: smtp.mail.yahoo.com
//   Port: 587 或 465
//
// QQ邮箱:
//   SMTP: smtp.qq.com
//   Port: 587 或 465
//   需要开启 SMTP 服务并使用授权码
//
// 163邮箱:
//   SMTP: smtp.163.com
//   Port: 465 或 994
//   需要开启 SMTP 服务并使用授权码
*/
