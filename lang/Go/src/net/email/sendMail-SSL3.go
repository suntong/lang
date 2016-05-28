package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/mail"
	"net/smtp"
)

sss 
func main() {
	from := mail.Address{"", "receiver's email address"}
	to := mail.Address{"", "sender's email address "}
	subj := "sample mail "
	body := "sample body <html> your things :)</html> "
	fmt.Println(body)

	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subj
	headers["Content-type"] = "text/html" //this will allow to send html body.if you
	//remove this part what ever the html body you try to send in this email will show as
	//a normal string(text )

	// Setup message
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	// Connect to the SMTP Server
	servername := "smtp.gmail.com:465"
	host, _, _ := net.SplitHostPort(servername)
	auth := smtp.PlainAuth("", from, "Sender gamil account password ", host)
	// TLS config
	//fmt.Print("7")
	tlsconfig := &tls.Config{InsecureSkipVerify: true, ServerName: host}

	// Here is the key, you need to call tls.Dial instead of smtp.Dial
	// for smtp servers running on 465 that require an ssl connection
	// from the very beginning (no starttls)
	//fmt.Print("8")
	conn, err := tls.Dial("tcp", servername, tlsconfig)
	if err != nil {
		log.Panic(err)
	}
	c, err := smtp.NewClient(conn, host)
	if err != nil {
		log.Panic(err)
	}

	// Auth
	if err = c.Auth(auth); err != nil {
		log.Panic(err)
	}
	// To && From
	if err = c.Mail(from.Address); err != nil {
		log.Panic(err)
	}
	if err = c.Rcpt(to.Address); err != nil {
		log.Panic(err)
	}
	// Data
	w, err := c.Data()
	if err != nil {
		log.Panic(err)
	}
	_, err = w.Write([]byte(message))
	if err != nil {
		log.Panic(err)
	}
	err = w.Close()
	if err != nil {
		log.Panic(err)
	} else {
		fmt.Println("\nMail sent sucessfully....")
	}
	c.Quit()
}
