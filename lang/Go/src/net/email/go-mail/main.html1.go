package main

import (
	"log"
	"strings"

	"github.com/wneessen/go-mail"
)

var (
	smtpServer = "smtp.example.com" // Port: 587 for SMTP
	smtpUser   = "bg@example.com"
	smtpName   = "Barry Gibbs"
	smtpPass   = "password*******"
	mesgTo     = "Alice <alice@example.com>, Bob <bob@example.com>, Eve <eve@example.com>"
)

func main() {
	// Create a new message
	m := mail.NewMsg()
	if err := m.FromFormat(smtpName, smtpUser); err != nil {
		log.Fatalf("failed to set From address: %s", err)
	}
	if err := m.To(strings.Split(mesgTo, ",")...); err != nil {
		log.Fatalf("failed to set To address: %s", err)
	}
	m.Subject("Test mail from go-mail")
	m.SetBodyString(mail.TypeTextHTML, `<p>Hi my friend,</p>
<p>Hello from <b>Bob</b> and <i>Cora</i>!</p>`)

	// Sending the email
	c, err := mail.NewClient(smtpServer, mail.WithPort(587),
		// mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithSMTPAuth(mail.SMTPAuthLogin), // for Outlook
		mail.WithUsername(smtpUser), mail.WithPassword(smtpPass))
	if err != nil {
		log.Fatalf("failed to create mail client: %s", err)
	}
	if err := c.DialAndSend(m); err != nil {
		log.Fatalf("failed to send mail: %s", err)
	}
}
