package main

import (
	"fmt"
	"log"
	"net/mail"
)

func main() {
	ParseAddress()
	ParseAddressList()
}

func ParseAddress() {
	e, err := mail.ParseAddress("Alice <alice@example.com>")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(e.Name, e.Address)

}

func ParseAddressList() {
	const list = `bg@example.com, Alice <alice@example.com>, "Bob Uncle" <bob@example.com>, "Yang-Tan, Eve" <eve@example.com>`
	emails, err := mail.ParseAddressList(list)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range emails {
		fmt.Println(v.Name, "=>", v.Address)
	}

}

/*

Output:

Alice alice@example.com
 => bg@example.com
Alice => alice@example.com
Bob Uncle => bob@example.com
Yang-Tan, Eve => eve@example.com

*/
