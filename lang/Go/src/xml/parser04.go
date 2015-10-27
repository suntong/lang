////////////////////////////////////////////////////////////////////////////
// Porgram: parser03
// Purpose: Go xml parsing demo
// Authors: Tong Sun (c) 2015, All rights reserved
// Credits: http://rosettacode.org/wiki/XML/XPath#Go
//
////////////////////////////////////////////////////////////////////////////

package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"strings"
)

var data = `
<inventory title="OmniCorp Store #45x10^3">
  <section name="health">
    <item upc="123456789" stock="12">
      <name>Invisibility Cream</name>
      <price>14.50</price>
      <description>Makes you invisible</description>
    </item>
    <item upc="445322344" stock="18">
      <name>Levitation Salve</name>
      <price>23.99</price>
      <description>Levitate yourself for up to 3 hours per application</description>
    </item>
  </section>
  <section name="food">
    <item upc="485672034" stock="653">
      <name>Blork and Freen Instameal</name>
      <price>4.95</price>
      <description>A tasty meal in a tablet; just add water</description>
    </item>
    <item upc="132957764" stock="44">
      <name>Grob winglets</name>
      <price>3.56</price>
      <description>Tender winglets of Grob. Just add water</description>
    </item>
  </section>
</inventory>
`

type Inventory struct {
	XMLName  xml.Name `xml:"inventory"`
	Title    string   `xml:"title,attr"`
	Sections []struct {
		XMLName xml.Name `xml:"section"`
		Name    string   `xml:"name,attr"`
		Items   []struct {
			XMLName     xml.Name `xml:"item"`
			Name        string   `xml:"name"`
			UPC         string   `xml:"upc,attr"`
			Stock       int      `xml:"stock,attr"`
			Price       float64  `xml:"price"`
			Description string   `xml:"description"`
		} `xml:"item"`
	} `xml:"section"`
}

// To simplify main's error handling
func printXML(s string, v interface{}) {
	fmt.Println(s)
	b, err := xml.MarshalIndent(v, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
	fmt.Println()
}

func main() {

	fmt.Println("Reading XML")

	var inv Inventory
	dec := xml.NewDecoder(strings.NewReader(data))
	if err := dec.Decode(&inv); err != nil {
		log.Fatal(err)
	}

	// At this point, inv is Go struct with all the fields filled
	// in from the XML data. Well-formed XML input that doesn't
	// match the specification of the fields in the Go struct are
	// discarded without error.

	// We can reformat the parts we parsed:
	//printXML("Got:", inv)

	// 1. Retrieve first item:
	item := inv.Sections[0].Items[0]
	fmt.Println("item variable:", item)
	printXML("As XML:", item)

	// 2. Action on each price:
	fmt.Println("Prices:")
	var totalValue float64
	for _, s := range inv.Sections {
		for _, i := range s.Items {
			fmt.Println(i.Price)
			totalValue += i.Price * float64(i.Stock)
		}
	}
	fmt.Println("Total inventory value:", totalValue)
	fmt.Println()

	// 3. Slice of all the names:
	var names []string
	for _, s := range inv.Sections {
		for _, i := range s.Items {
			names = append(names, i.Name)
		}
	}
	fmt.Printf("names: %q\n", names)
}
