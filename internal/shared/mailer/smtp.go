package mailer

import (
	"fmt"
	"net/smtp"
)

type SMTPConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	From     string
}

type SMTPMailer struct {
	cfg  SMTPConfig
	auth smtp.Auth
}

func NewSMTPMailer(cfg SMTPConfig) *SMTPMailer {
	auth := smtp.PlainAuth(
		"",
		cfg.Username,
		cfg.Password,
		cfg.Host,
	)

	return &SMTPMailer{
		cfg:  cfg,
		auth: auth,
	}
}

func (m *SMTPMailer) SendVerificationEmail(toEmail, verifyLink string) error {
	subject := "Verify your email"
	body := fmt.Sprintf(`
Hi 👋

Please verify your email by clicking the link below:

%s

If you did not create an account, you can safely ignore this email.

Thanks,
Ecommerce Team
`, verifyLink)

	msg := []byte(
		"From: " + m.cfg.From + "\r\n" +
			"To: " + toEmail + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"MIME-Version: 1.0\r\n" +
			"Content-Type: text/plain; charset=\"UTF-8\"\r\n" +
			"\r\n" +
			body,
	)

	addr := fmt.Sprintf("%s:%s", m.cfg.Host, m.cfg.Port)

	return smtp.SendMail(
		addr,
		m.auth,
		m.cfg.From,
		[]string{toEmail},
		msg,
	)
}
