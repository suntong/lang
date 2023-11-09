package main

import (
	"log"
	"time"

	"github.com/toorop/go-dkim"
	mail "github.com/xhit/go-simple-mail/v2"
)

const htmlBody = `<html>
	<head>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
		<title>Hello Gophers!</title>
	</head>
	<body>
		<p>This is the <b>Go gopher</b>.</p>
		<p><img src="cid:Gopher.png" alt="Go gopher" /></p>
		<p>Image created by Renee French</p>
	</body>
</html>`

const privateKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAwGrscUxi9zEa9oMOJbS0kLVHZXNIW+EBjY7KFWIZSxuGAils
wBVl+s5mMRrR5VlkyLQNulAdNemg6OSeB0R+2+8/lkHiMrimqQckZ5ig8slBoZhZ
wUoL/ZkeQa1bacbdww5TuWkiVPD9kooT/+TZW1P/ugd6oYjpOI56ZjsXzJw5pz7r
DiwcIJJaaDIqvvc5C4iW94GZjwtmP5pxhvBZ5D6Uzmh7Okvi6z4QCKzdJQLdVmC0
CMiFeh2FwqMkVpjZhNt3vtCo7Z51kwHVscel6vl51iQFq/laEzgzAWOUQ+ZEoQpL
uTaUiYzzNyEdGEzZ2CjMMoO8RgtXnUo2qX2FDQIDAQABAoIBAHWKW3kycloSMyhX
EnNSGeMz+bMtYwxNPMeebC/3xv+shoYXjAkiiTNWlfJ1MbbqjrhT1Pb1LYLbfqIF
1csWum/bjHpbMLRPO++RH1nxUJA/BMqT6HA8rWpy+JqiLW9GPf2DaP2gDYrZ0+yK
UIFG6MfzXgnju7OlkOItlvOQMY+Y501u/h6xnN2yTeRqXXJ1YlWFPRIeFdS6UOtL
J2wSxRVdymHbGwf+D7zet7ngMPwFBsbEN/83KGLRjkt8+dMQeUeob+nslsQofCZx
iokIAvByTugmqrB4JqhNkAlZhC0mqkRQh7zUFrxSj5UppMWlxLH+gPFZHKAsUJE5
mqmylcECgYEA8I/f90cpF10uH4NPBCR4+eXq1PzYoD+NdXykN65bJTEDZVEy8rBO
phXRNfw030sc3R0waQaZVhFuSgshhRuryfG9c1FP6tQhqi/jiEj9IfCW7zN9V/P2
r16pGjLuCK4SyxUC8H58Q9I0X2CQqFamtkLXC6Ogy86rZfIc8GcvZ9UCgYEAzMQZ
WAiLhRF2MEmMhKL+G3jm20r+dOzPYkfGxhIryluOXhuUhnxZWL8UZfiEqP5zH7Li
NeJvLz4pOL45rLw44qiNu6sHN0JNaKYvwNch1wPT/3/eDNZKKePqbAG4iamhjLy5
gjO1KgA5FBbcNN3R6fuJAg1e4QJCOuo55eW6vFkCgYEA7UBIV72D5joM8iFzvZcn
BPdfqh2QnELxhaye3ReFZuG3AqaZg8akWqLryb1qe8q9tclC5GIQulTInBfsQDXx
MGLNQL0x/1ylsw417kRl+qIoidMTTLocUgse5erS3haoDEg1tPBaKB1Zb7NyF8QV
+W1kX2NKg5bZbdrh9asekt0CgYA6tUam7NxDrLv8IDo/lRPSAJn/6cKG95aGERo2
k+MmQ5XP+Yxd+q0LOs24ZsZyRXHwdrNQy7khDGt5L2EN23Fb2wO3+NM6zrGu/WbX
nVbAdQKFUL3zZEUjOYtuqBemsJH27e0qHXUls6ap0dwU9DxJH6sqgXbggGtIxPsQ
pQsjEQKBgQC9gAqAj+ZtMXNG9exVPT8I15reox9kwxGuvJrRu/5eSi6jLR9z3x9P
2FrgxQ+GCB2ypoOUcliXrKesdSbolUilA8XQn/M113Lg8oA3gJXbAKqbTR/EgfUU
kvYaR/rTFnivF4SL/P4k/gABQoJuFUtSKdouELqefXB+e94g/G++Bg==
-----END RSA PRIVATE KEY-----`

func main() {
	server := mail.NewSMTPClient()

	// SMTP Server
	server.Host = "smtp.example.com"
	server.Port = 587
	server.Username = "test@example.com"
	server.Password = "examplepass"
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

	// Set TLSConfig to provide custom TLS configuration. For example,
	// to skip TLS verification (useful for testing):
	//server.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// SMTP client
	smtpClient, err := server.Connect()

	if err != nil {
		log.Fatal(err)
	}

	// New email simple html with inline and CC
	email := mail.NewMSG()
	email.SetFrom("From Example <nube@example.com>").
		AddTo("xhit@example.com").
		AddCc("otherto@example.com").
		SetSubject("New Go Email").
		SetListUnsubscribe("<mailto:unsubscribe@example.com?subject=https://example.com/unsubscribe>")

	email.SetBody(mail.TextHTML, htmlBody)

	// also you can add body from []byte with SetBodyData, example:
	// email.SetBodyData(mail.TextHTML, []byte(htmlBody))
	// or alternative part
	// email.AddAlternativeData(mail.TextHTML, []byte(htmlBody))

	// add inline
	email.Attach(&mail.File{FilePath: "/path/to/image.png", Name: "Gopher.png", Inline: true})

	// also you can set Delivery Status Notification (DSN) (only is set when server supports DSN)
	email.SetDSN([]mail.DSN{mail.SUCCESS, mail.FAILURE}, false)

	// you can add dkim signature to the email.
	// to add dkim, you need a private key already created one.
	if privateKey != "" {
		options := dkim.NewSigOptions()
		options.PrivateKey = []byte(privateKey)
		options.Domain = "example.com"
		options.Selector = "default"
		options.SignatureExpireIn = 3600
		options.Headers = []string{"from", "date", "mime-version", "received", "received"}
		options.AddSignatureTimestamp = true
		options.Canonicalization = "relaxed/relaxed"

		email.SetDkim(options)
	}

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
