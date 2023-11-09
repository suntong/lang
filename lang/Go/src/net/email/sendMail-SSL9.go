////////////////////////////////////////////////////////////////////////////
// Porgram: sendMail-SSL9.go
// Purpose: SSL/TLS Email Example
// Authors: Tong Sun (c) 2016, All rights reserved
// Credits: https://gist.github.com/chrisgillis/10888032
////////////////////////////////////////////////////////////////////////////

package main

import (
	"crypto/tls"
	"log"
	"net"
	"net/mail"
	"net/smtp"
	"os"
	"strconv"
	"strings"
)

var (
	smtpServer = "smtp.example.com:465"
	smtpUser   = "Barry Gibbs <bg@example.com>"
	smtpPass   = "password*******"
	mesgTo     = "Alice <alice@example.com>, Bob <bob@example.com>, Eve <eve@example.com>"
)

func main() {
	a, r := "", 0
	if len(os.Args) > 1 {
		a = os.Args[1]
		println("Provided command line argument for string repeat (in 1K):", a)
		r, _ = strconv.Atoi(a)
		r *= 1024
	}

	SendEmail(mesgTo,
		"Test email subject "+a,
		"<html><body><h1>Hello World!</h1>"+
			"<p>This is an example body.<p>With two lines.<p>\n"+
			strings.Repeat(`&nbsp;`, r)+
			"\n.\n</body></html>")

}

func SendEmail(to string, subject string, body string) (err error) {

	mime := "MIME-version: 1.0;\r\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	parameters := struct {
		From    *mail.Address
		To      []*mail.Address
		Subject string
		Body    string
	}{
		&mail.Address{},
		[]*mail.Address{},
		subject,
		body,
	}

	parameters.From, err = mail.ParseAddress(smtpUser)
	parameters.To, err = mail.ParseAddressList(to)
	message := mime +
		"From: " + smtpUser + "\r\n" +
		"To: " + to + "\r\n" +
		"Subject: " + parameters.Subject + "\r\n" +
		"\r\n" + parameters.Body

	// Connect to the SMTP Server
	host, _, _ := net.SplitHostPort(smtpServer)
	auth := smtp.PlainAuth("", parameters.From.Address, smtpPass, host)

	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	// Here is the key, you need to call tls.Dial instead of smtp.Dial
	// for smtp servers running on 465 that require an ssl connection
	// from the very beginning (no starttls)
	conn, err := tls.Dial("tcp", smtpServer, tlsconfig)
	checkError(err)

	c, err := smtp.NewClient(conn, host)
	checkError(err)

	// Auth
	err = c.Auth(auth)
	checkError(err)

	// To && From
	err = c.Mail(parameters.From.Address)
	checkError(err)
	for _, addr := range parameters.To {
		err = c.Rcpt(addr.Address)
		checkError(err)
	}

	// Data
	w, err := c.Data()
	checkError(err)
	_, err = w.Write([]byte(message))
	checkError(err)
	err = w.Close()
	checkError(err)

	c.Quit()
	return err
}

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

/*

Refs,

- https://github.com/suntong/lang/blob/master/lang/Go/src/net/email/ParseAddress.go
  for how to specify `mesgTo`

- https://github.com/suntong/lang/blob/master/lang/Go/src/sys/VarOverride.go
  for how to run/test this without modifying the code (for e.g.,
  to use your own email address and smtp passwerod for the testing). I.e.,

  go run -ldflags="-X main.smtpUser=me@myorg.com -X main.smtpPass=mysecret -X main.smtpServer=smtp.myorg.com:465 -X main.mesgTo=me2@myorg.com" sendMail-SSL9.go

*/
