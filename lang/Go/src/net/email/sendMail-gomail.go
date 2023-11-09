////////////////////////////////////////////////////////////////////////////
// Porgram: sendMail-SSL9.go
// Purpose: SSL/TLS Email Example
// Authors: Tong Sun (c) 2016, All rights reserved
// Credits: https://gist.github.com/chrisgillis/10888032
////////////////////////////////////////////////////////////////////////////

package main

import "gopkg.in/mail.v2"

var (
	smtpServer = "smtp.example.com" // Port: 587 for SMTP
	smtpUser   = "Barry Gibbs <bg@example.com>"
	smtpPass   = "password*******"
	mesgTo     = "Alice <alice@example.com>, Bob <bob@example.com>, Eve <eve@example.com>"
)

func main() {
	m := mail.NewMessage()
	m.SetHeader("From", smtpUser)
	m.SetHeader("To", mesgTo)
	//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")
	//m.Attach("/home/Alex/lolcat.jpg")

	d := mail.NewDialer(smtpServer, 587, smtpUser, smtpPass)

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
