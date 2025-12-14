package mail

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"strings"
)

// MailSender 邮件发送器
type MailSender struct {
	SMTPHost     string // SMTP 服务器地址 (例如: smtp.gmail.com)
	SMTPPort     int    // SMTP 端口 (例如: 587)
	FromEmail    string // 发件人邮箱
	FromPassword string // 发件人密码或应用专用密码
	FromName     string // 发件人名称（可选）
}

// NewMailSender 创建邮件发送器
func NewMailSender(smtpHost string, smtpPort int, fromEmail, fromPassword, fromName string) *MailSender {
	return &MailSender{
		SMTPHost:     smtpHost,
		SMTPPort:     smtpPort,
		FromEmail:    fromEmail,
		FromPassword: fromPassword,
		FromName:     fromName,
	}
}

// SendMail 发送邮件
// to: 收件人邮箱
// subject: 邮件标题
// body: 邮件内容（支持 HTML）
func (m *MailSender) SendMail(to, subject, body string) error {
	return m.SendMailToMultiple([]string{to}, subject, body)
}

// SendMailToMultiple 发送邮件到多个收件人
// toList: 收件人邮箱列表
// subject: 邮件标题
// body: 邮件内容（支持 HTML）
func (m *MailSender) SendMailToMultiple(toList []string, subject, body string) error {
	// 构建邮件头
	from := m.FromEmail
	if m.FromName != "" {
		from = fmt.Sprintf("%s <%s>", m.FromName, m.FromEmail)
	}

	headers := make(map[string]string)
	headers["From"] = from
	headers["To"] = strings.Join(toList, ", ")
	headers["Subject"] = subject
	headers["MIME-Version"] = "1.0"
	headers["Content-Type"] = "text/html; charset=UTF-8"

	// 构建邮件消息
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	// SMTP 认证
	auth := smtp.PlainAuth("", m.FromEmail, m.FromPassword, m.SMTPHost)

	// 连接地址
	addr := fmt.Sprintf("%s:%d", m.SMTPHost, m.SMTPPort)

	// 发送邮件
	err := smtp.SendMail(addr, auth, m.FromEmail, toList, []byte(message))
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}

// SendMailWithTLS 使用 TLS 发送邮件（适用于 Gmail 等）
func (m *MailSender) SendMailWithTLS(to, subject, body string) error {
	return m.SendMailToMultipleWithTLS([]string{to}, subject, body)
}

// SendMailToMultipleWithTLS 使用 TLS 发送邮件到多个收件人
func (m *MailSender) SendMailToMultipleWithTLS(toList []string, subject, body string) error {
	// 构建邮件头
	from := m.FromEmail
	if m.FromName != "" {
		from = fmt.Sprintf("%s <%s>", m.FromName, m.FromEmail)
	}

	headers := make(map[string]string)
	headers["From"] = from
	headers["To"] = strings.Join(toList, ", ")
	headers["Subject"] = subject
	headers["MIME-Version"] = "1.0"
	headers["Content-Type"] = "text/html; charset=UTF-8"

	// 构建邮件消息
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	// 连接地址
	addr := fmt.Sprintf("%s:%d", m.SMTPHost, m.SMTPPort)

	// TLS 配置
	tlsConfig := &tls.Config{
		ServerName: m.SMTPHost,
	}

	// 建立 TLS 连接
	conn, err := tls.Dial("tcp", addr, tlsConfig)
	if err != nil {
		return fmt.Errorf("failed to connect with TLS: %w", err)
	}
	defer conn.Close()

	// 创建 SMTP 客户端
	client, err := smtp.NewClient(conn, m.SMTPHost)
	if err != nil {
		return fmt.Errorf("failed to create SMTP client: %w", err)
	}
	defer client.Quit()

	// SMTP 认证
	auth := smtp.PlainAuth("", m.FromEmail, m.FromPassword, m.SMTPHost)
	if err := client.Auth(auth); err != nil {
		return fmt.Errorf("SMTP auth failed: %w", err)
	}

	// 设置发件人
	if err := client.Mail(m.FromEmail); err != nil {
		return fmt.Errorf("failed to set sender: %w", err)
	}

	// 设置收件人
	for _, to := range toList {
		if err := client.Rcpt(to); err != nil {
			return fmt.Errorf("failed to set recipient %s: %w", to, err)
		}
	}

	// 发送邮件内容
	w, err := client.Data()
	if err != nil {
		return fmt.Errorf("failed to get data writer: %w", err)
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		return fmt.Errorf("failed to write message: %w", err)
	}

	err = w.Close()
	if err != nil {
		return fmt.Errorf("failed to close writer: %w", err)
	}

	return nil
}

// SendPlainText 发送纯文本邮件
func (m *MailSender) SendPlainText(to, subject, body string) error {
	// 构建邮件头
	from := m.FromEmail
	if m.FromName != "" {
		from = fmt.Sprintf("%s <%s>", m.FromName, m.FromEmail)
	}

	headers := make(map[string]string)
	headers["From"] = from
	headers["To"] = to
	headers["Subject"] = subject
	headers["MIME-Version"] = "1.0"
	headers["Content-Type"] = "text/plain; charset=UTF-8"

	// 构建邮件消息
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	// SMTP 认证
	auth := smtp.PlainAuth("", m.FromEmail, m.FromPassword, m.SMTPHost)

	// 连接地址
	addr := fmt.Sprintf("%s:%d", m.SMTPHost, m.SMTPPort)

	// 发送邮件
	err := smtp.SendMail(addr, auth, m.FromEmail, []string{to}, []byte(message))
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
