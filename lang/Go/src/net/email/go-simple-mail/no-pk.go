package main

import (
	"log"
	"time"

	mail "github.com/xhit/go-simple-mail/v2"
)

var (
	smtpServer = "smtp.example.com" // Port: 587 for SMTP
	smtpUser   = "bg@example.com"
	smtpPass   = "password*******"
	mesgTo     = "Alice <alice@example.com>, Bob <bob@example.com>, Eve <eve@example.com>"
)

const htmlBody = `<html>
	<head>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
		<title>Hello!</title>
	</head>
	<body>
		<p>This is <i>the</i> <b>gopher</b>.</p>
	</body>
</html>`

func main() {
	server := mail.NewSMTPClient()

	// SMTP Server
	server.Host = smtpServer
	server.Port = 587
	server.Username = smtpUser
	server.Password = smtpPass
	server.Encryption = mail.EncryptionSTARTTLS

	// You can specified authentication type:
	// - AUTO (default)
	// - PLAIN
	// - LOGIN
	// - CRAM-MD5
	// - None
	// server.Authentication = mail.AuthAuto

	// Variable to keep alive connection
	server.KeepAlive = false

	// Timeout for connect to SMTP Server
	server.ConnectTimeout = 10 * time.Second

	// Timeout for send the data and wait respond
	server.SendTimeout = 10 * time.Second

	// SMTP client
	smtpClient, err := server.Connect()

	if err != nil {
		log.Fatal(err)
	}

	// New email simple html with inline and CC
	email := mail.NewMSG()
	email.SetFrom(smtpUser).
		AddTo(mesgTo).
		//AddCc("otherto@example.com").
		SetSubject("Test Email") //.
		//SetListUnsubscribe("<mailto:unsubscribe@example.com?subject=https://example.com/unsubscribe>")

	email.SetBody(mail.TextHTML, htmlBody)

	// also you can add body from []byte with SetBodyData, example:
	// email.SetBodyData(mail.TextHTML, []byte(htmlBody))
	// or alternative part
	// email.AddAlternativeData(mail.TextHTML, []byte(htmlBody))

	// add inline
	//email.Attach(&mail.File{FilePath: "/path/to/image.png", Name: "Gopher.png", Inline: true})

	// also you can set Delivery Status Notification (DSN) (only is set when server supports DSN)
	//email.SetDSN([]mail.DSN{mail.SUCCESS, mail.FAILURE}, false)

	// always check error after send
	if email.Error != nil {
		log.Fatal(email.Error)
	}

	// Call Send and pass the client
	err = email.Send(smtpClient)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Email Sent")
	}
}
