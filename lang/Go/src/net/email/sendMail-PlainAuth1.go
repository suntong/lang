package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/mail"
	"net/smtp"
	"strings"
)

var (
	smtpServer = "smtp.163.com"
	smtpPort   = ":25"
	smtpUser   = "fledna@163.com"
	smtpPass   = "password*******"
	mesgTo     = "name@139.com"
)

func encodeRFC2047(String string) string {
	// use mail's rfc2047 to encode any string
	addr := mail.Address{String, ""}
	return strings.Trim(addr.String(), " <>")
}

func main() {
	//log.Fatal(smtpUser, smtpPass, smtpServer, mesgTo)

	// Set up authentication information.
	auth := smtp.PlainAuth(
		"",
		smtpUser,
		smtpPass,
		smtpServer,
	)

	from := mail.Address{"监控中心", smtpUser}
	to := mail.Address{"收件人", mesgTo}
	title := "当前时段统计报表"

	body := "报表内容一切正常"

	header := make(map[string]string)
	header["From"] = from.String()
	header["To"] = to.String()
	header["Subject"] = encodeRFC2047(title)
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/plain; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(body))

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	err := smtp.SendMail(
		smtpServer+smtpPort,
		auth,
		from.Address,
		[]string{to.Address},
		[]byte(message),
		//[]byte("This is the email body."),
	)
	if err != nil {
		log.Fatal(err)
	}
}
