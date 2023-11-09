////////////////////////////////////////////////////////////////////////////
// Porgram: sendMail-SMTP.go
// Purpose: SMTP Email Example
// Authors: Tong Sun (c) 2023, All rights reserved
// Credits: https://pkg.go.dev/net/smtp#example-package
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"log"
	"net/smtp"
)

var (
	smtpServer = "smtp.example.com:465"
	smtpUser   = "Barry Gibbs <bg@example.com>"
	smtpPass   = "password*******"
	mesgTo     = "Alice <alice@example.com>, Bob <bob@example.com>, Eve <eve@example.com>"
)

// NOK: No auth setup
func main() {
	// Connect to the remote SMTP server.
	c, err := smtp.Dial(smtpServer)
	if err != nil {
		log.Fatal(err)
	}

	// Set the sender and recipient first
	if err := c.Mail(smtpUser); err != nil {
		log.Fatal(err)
	}
	if err := c.Rcpt(mesgTo); err != nil {
		log.Fatal(err)
	}

	// Send the email body.
	wc, err := c.Data()
	if err != nil {
		log.Fatal(err)
	}
	_, err = fmt.Fprintf(wc, "This is the email body")
	if err != nil {
		log.Fatal(err)
	}
	err = wc.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Send the QUIT command and close the connection.
	err = c.Quit()
	if err != nil {
		log.Fatal(err)
	}
}
