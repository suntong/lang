package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
)

var auth smtp.Auth

func main() {
	auth = smtp.PlainAuth("", "dhanush@geektrust.in", "password", "smtp.gmail.com")
	templateData := struct {
		Name string
		URL  string
	}{
		Name: "Dhanush",
		URL:  "http://geektrust.in",
	}
	r := NewRequest([]string{"junk@junk.com"}, "Hello Junk!", "Hello, World!")
	err := r.ParseTemplate("template.html", templateData)
	if err != nil {
		ok, _ := r.SendEmail()
		fmt.Println(ok)
	}
}

//Request struct
type Request struct {
	from    string
	to      []string
	subject string
	body    string
}

func NewRequest(to []string, subject, body string) *Request {
	return &Request{
		to:      to,
		subject: subject,
		body:    body,
	}
}

func (r *Request) SendEmail() (bool, error) {
	mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	subject := "Subject: " + r.subject + "!\n"
	msg := []byte(subject + mime + "\n" + r.body)
	addr := "smtp.gmail.com:587"

	if err := smtp.SendMail(addr, auth, "dhanush@geektrust.in", r.to, msg); err != nil {
		return false, err
	}
	return true, nil
}

func (r *Request) ParseTemplate(templateFileName string, data interface{}) error {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}
	r.body = buf.String()
	return nil
}

/*

From
https://medium.com/@dhanushgopinath/sending-html-emails-using-templates-in-golang-9e953ca32f3d#.ua4ydv9vj

send emails that are auto-generated at runtime using HTML Templates.
sending email using Go’s html/template and net/smtp packages.

I encapsulate the smtp request in a struct Request. It contains basic things like To, Subject, Body. This template file has the data placeholders which will be replaced with actual data values by Golang’s html/template package Execute method.

<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN"
        "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html>

</head>

<body>
<p>
    Hello {{.Name}}
    <a href="{{.URL}}">Confirm email address</a>
</p>

</body>

</html>

The template data is a struct which has Name & Url as the field values. It
is initialised with values and passed to the NewRequest method to create an
object of Request.

Once the instance of Request is created, then ParseTemplate method is called
on it. This method receives a template filename and template data. It parses
the template file and executes it with the template data supplied. A
bytes.Buffer is passed to the Execute method as io.Writer so that we get
back the HTML string with the data replaced. This HTML string is then set as
the Email Body.

The SendEmail method sets the MIME encoding as text/html and calls smtp
package’s SendMail method to send the email. When the email is sent
successfully the SendEmail method returns a true value.

*/
