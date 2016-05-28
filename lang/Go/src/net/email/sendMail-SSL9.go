package main

import (
	"crypto/tls"
	//"fmt"
	"log"
	"net"
	"net/mail"
	"net/smtp"
	"strings"
)

// SSL/TLS Email Example

var (
	smtpServer = "smtp.example.tld:465"
	smtpUser   = "username@example.tld"
	smtpUName  = "Real User Name"
	smtpPass   = "password*******"
	mesgTo     = "username@anotherexample.tld"
)

func main() {

	SendEmail([]string{mesgTo},
		"This is the email subject",
		"<html><body><h1>Hello World!</h1><p>This is an example body.<p>With two lines.</body></html>")

}

func SendEmail(_to []string, _subject string, body string) (err error) {
	servername := smtpServer
	userName := smtpUser
	password := smtpPass

	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n"
	parameters := struct {
		From    string
		To      string
		Subject string
		Body    string
	}{
		userName,
		strings.Join([]string(_to), ","),
		_subject,
		body,
	}

	from := mail.Address{smtpUName, parameters.From}
	to := mail.Address{"", parameters.To}
	message := mime +
		"From: " + from.String() + "\r\n" +
		"To: " + to.String() + "\r\n" +
		"Subject: " + parameters.Subject + "\r\n" +
		"\r\n" + parameters.Body

	// Connect to the SMTP Server
	host, _, _ := net.SplitHostPort(servername)
	auth := smtp.PlainAuth("", userName, password, host)

	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	// Here is the key, you need to call tls.Dial instead of smtp.Dial
	// for smtp servers running on 465 that require an ssl connection
	// from the very beginning (no starttls)
	conn, err := tls.Dial("tcp", servername, tlsconfig)
	checkError(err)

	c, err := smtp.NewClient(conn, host)
	checkError(err)

	// Auth
	err = c.Auth(auth)
	checkError(err)

	// To && From
	err = c.Mail(from.Address)
	checkError(err)
	err = c.Rcpt(to.Address)
	checkError(err)

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
