////////////////////////////////////////////////////////////////////////////
// Porgram: unmarshal01.go
// Purpose: Go xml unmarshal demo
// Authors: Tong Sun (c) 2015, All rights reserved
// Credits: https://golang.org/pkg/encoding/xml/
////////////////////////////////////////////////////////////////////////////

package main

import (
	"encoding/xml"
	"fmt"
)

func main() {
	type Email struct {
		Where string `xml:"where,attr"`
		Addr  string
	}
	type Address struct {
		City, State string
	}
	type Result struct {
		XMLName xml.Name `xml:"Person"`
		Name    string   `xml:"FullName"`
		Company string   `xml:"Company"`
		Phone   string
		Email   []Email
		Groups  []string `xml:"Group>Value"`
		Address
	}
	v := Result{Name: "none", Phone: "none"}

	data := `
		<Person>
			<FullName>Grace R. Emlin</FullName>
			<Company>
			 <Name>Example Inc.</Name>
			 <Addr>Example work place</Addr>
			</Company>
			<Email where="home">
				<Addr>gre@example.com</Addr>
			</Email>
			<Email where='work'>
				<Addr>gre@work.com</Addr>
			</Email>
			<Group>
				<Value>Friends</Value>
				<Value>Squash</Value>
			</Group>
			<City>Hanga Roa</City>
			<State>Easter Island</State>
		</Person>
	`
	err := xml.Unmarshal([]byte(data), &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	//printXML("As XML:", v)
	printXML("Name As XML:", v.Name)
	printXML("Email As XML:", v.Email)
	printXML("Groups As XML:", v.Groups)
	printXML("Company As XML:", v.Company)

	fmt.Printf("XMLName: %#v\n", v.XMLName)
	fmt.Printf("Name: %q\n", v.Name)
	fmt.Printf("Phone: %q\n", v.Phone)
	fmt.Printf("Email: %v\n", v.Email)
	fmt.Printf("Groups: %v\n", v.Groups)
	fmt.Printf("Address: %v\n", v.Address)
}

func printXML(s string, v interface{}) {
	fmt.Println(s)
	b, _ := xml.MarshalIndent(v, "", "\t")
	fmt.Println(string(b))
	fmt.Println()
}
