// https://godoc.org/github.com/jhillyerd/enmime#example-ReadEnvelope

package main

import (
	"fmt"
	"os"

	"github.com/jhillyerd/enmime"
)

func main() {
	// Open a sample message file
	r, err := os.Open("test.raw")
	if err != nil {
		fmt.Print(err)
		return
	}

	// Parse message body with enmime
	env, err := enmime.ReadEnvelope(r)
	if err != nil {
		fmt.Print(err)
		return
	}

	// Headers can be retrieved via Envelope.GetHeader(name)
	fmt.Printf("From: %v\n", env.GetHeader("From"))

	// Address-type headers can be parsed into a list of decoded mail.Address structs
	alist, _ := env.AddressList("To")
	for _, addr := range alist {
		fmt.Printf("To: %s <%s>\n", addr.Name, addr.Address)
	}

	// enmime can decode quoted-printable headers
	fmt.Printf("Subject: %v\n", env.GetHeader("Subject"))

	// The plain text body is available as mime.Text
	fmt.Printf("Text Body: %v chars\n", len(env.Text))
	fmt.Printf("Text Body: \n\t%s\n", env.Text)

	// The HTML body is stored in mime.HTML
	fmt.Printf("HTML Body: %v chars\n", len(env.HTML))
	fmt.Printf("HTML Body: \n%s\n", env.HTML)

	// mime.Inlines is a slice of inlined attacments
	fmt.Printf("Inlines: %v\n", len(env.Inlines))

	// mime.Attachments contains the non-inline attachments
	fmt.Printf("Attachments: %v\n", len(env.Attachments))
}
