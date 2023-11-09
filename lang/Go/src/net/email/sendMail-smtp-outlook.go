////////////////////////////////////////////////////////////////////////////
// Porgram: sendMail-SMTP.go
// Purpose: SMTP Email Example
// Authors: Tong Sun (c) 2023, All rights reserved
// Credits: https://gist.github.com/homme/22b457eb054a07e7b2fb
////////////////////////////////////////////////////////////////////////////

package main

import (
	"errors"
	"log"
	"net/mail"
	"net/smtp"
)

type loginAuth struct {
	username, password string
}

func LoginAuth(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte(a.username), nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		default:
			return nil, errors.New("Unkown fromServer")
		}
	}
	return nil, nil
}

// usage:
// auth := LoginAuth("loginname", "password")
// err := smtp.SendMail(smtpServer + ":25", auth, fromAddress, toAddresses, []byte(message))
// or
// client, err := smtp.Dial(smtpServer)
// client.Auth(LoginAuth("loginname", "password"))

var (
	smtpServer = "smtp.example.com:465"
	smtpUser   = "Barry Gibbs <bg@example.com>"
	smtpPass   = "password*******"
	mesgTo     = "Alice <alice@example.com>, Bob <bob@example.com>, Eve <eve@example.com>"
)

func main() {
	auth := LoginAuth(smtpUser, smtpPass)

	// https://pkg.go.dev/net/smtp#SendMail
	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	to, _ := mail.ParseAddressList(mesgTo)
	_ = to
	msg := []byte("To: " + mesgTo + "\r\n" +
		"Subject: discount Gophers!\r\n" +
		"\r\n" +
		"This is the email body.\r\n")
	err := smtp.SendMail(smtpServer, auth, smtpUser, []string{}, msg)
	if err != nil {
		log.Fatal(err)
	}
}

/*

NOK, Got:
 503 5.5.1 Bad sequence of commands

*/
