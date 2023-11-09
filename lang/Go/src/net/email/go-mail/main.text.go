package main

import (
	"log"

	"github.com/wneessen/go-mail"
)

var (
	smtpServer = "smtp.example.com" // Port: 587 for SMTP
	smtpUser   = "bg@example.com"
	smtpPass   = "password*******"
	mesgTo     = "Alice <alice@example.com>, Bob <bob@example.com>, Eve <eve@example.com>"
)

func main() {
	// Create a new message
	m := mail.NewMsg()
	if err := m.From(smtpUser); err != nil {
		log.Fatalf("failed to set From address: %s", err)
	}
	if err := m.To(mesgTo); err != nil {
		log.Fatalf("failed to set To address: %s", err)
	}
	m.Subject("This is my first mail with go-mail!")
	m.SetBodyString(mail.TypeTextPlain, "Do you like this mail? I certainly do!")

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
