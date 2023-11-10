package main

import (
	ht "html/template"
	"log"

	"github.com/wneessen/go-mail"
)

// User is a simple type allowing us to set a firstname, lastname and mail address
type User struct {
	Firstname string
	Lastname  string
	EmailAddr string
}

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
	if err := m.To(mesgTo); err != nil {
		log.Fatalf("failed to set To address: %s", err)
	}
	m.Subject("Test mail from go-mail")

	htmlBodyTemplate := `<p>Hi {{.Firstname}},</p>
<p>Hello from <b>Bob</b> and <i>Cora</i>!</p>`
	u := User{"Alice", "", ""}
	htpl, err := ht.New("htmltpl").Parse(htmlBodyTemplate)
	if err != nil {
		log.Fatalf("failed to parse text template: %s", err)
	}
	if err := m.SetBodyHTMLTemplate(htpl, u); err != nil {
		log.Fatalf("failed to set HTML template as HTML body: %s", err)
	}

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
