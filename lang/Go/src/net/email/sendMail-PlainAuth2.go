package main

import (
	"bytes"
	"fmt"
	"net/smtp"
	"runtime"
	"strings"
	"text/template"
)

func main() {
	SendEmail(
		"smtp.1and1.com",
		587,
		"username@domain.com",
		"password",
		[]string{"me@domain.com"},
		"testing subject",
		"<html><body>Exception 1</body></html>Exception 1")
}

func catchPanic(err *error, functionName string) {
	if r := recover(); r != nil {
		fmt.Printf("%s : PANIC Defered : %v\n", functionName, r)

		// Capture the stack trace
		buf := make([]byte, 10000)
		runtime.Stack(buf, false)

		fmt.Printf("%s : Stack Trace : %s", functionName, string(buf))

		if err != nil {
			*err = fmt.Errorf("%v", r)
		}
	} else if err != nil && *err != nil {
		fmt.Printf("%s : ERROR : %v\n", functionName, *err)

		// Capture the stack trace
		buf := make([]byte, 10000)
		runtime.Stack(buf, false)

		fmt.Printf("%s : Stack Trace : %s", functionName, string(buf))
	}
}

func SendEmail(host string, port int, userName string, password string, to []string, subject string, message string) (err error) {
	defer catchPanic(&err, "SendEmail")

	parameters := struct {
		From    string
		To      string
		Subject string
		Message string
	}{
		userName,
		strings.Join([]string(to), ","),
		subject,
		message,
	}

	buffer := new(bytes.Buffer)

	template := template.Must(template.New("emailTemplate").Parse(emailScript()))
	template.Execute(buffer, &parameters)

	auth := smtp.PlainAuth("", userName, password, host)

	err = smtp.SendMail(
		fmt.Sprintf("%s:%d", host, port),
		auth,
		userName,
		to,
		buffer.Bytes())

	return err
}

func emailScript() (script string) {
	return "From: {{.From}}<br /> To: {{.To}}<br /> Subject: {{.Subject}}<br /> MIME-version: 1.0<br /> Content-Type: text/html; charset=&quot;UTF-8&quot;<br /> <br /> {{.Message}}"
}

/*

From
https://www.goinggo.net/2013/06/send-email-in-go-with-smtpsendmail.html

The auth variable does not have to be recreated on every call. That can be
created once and reused. I added my catchPanic function so you can see any
exceptions that are returned while testing the code.

*/
