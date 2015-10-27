////////////////////////////////////////////////////////////////////////////
// Porgram: parser03
// Purpose: Go xml parsing demo
// Authors: Tong Sun (c) 2015, All rights reserved
// Credits: http://rosettacode.org/wiki/XML/XPath#Go
//
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"strings"
)

import "gopkg.in/xmlpath.v1"

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

func main() {
	/*
			f, err := os.Open("test3.xml")
		  defer f.Close()
			if err != nil {
				fmt.Println(err)
				return
			}
			n, err := xmlpath.Parse(f)
	*/

	n, err := xmlpath.Parse(strings.NewReader(data))
	if err != nil {
		fmt.Println(err)
		return
	}

	q1 := xmlpath.MustCompile("//item")
	if _, ok := q1.String(n); !ok {
		fmt.Println("no item")
	}
	q2 := xmlpath.MustCompile("//price")
	for it := q2.Iter(n); it.Next(); {
		fmt.Println(it.Node())
	}
	q3 := xmlpath.MustCompile("//name")
	names := []*xmlpath.Node{}
	for it := q3.Iter(n); it.Next(); {
		names = append(names, it.Node())
	}
	if len(names) == 0 {
		fmt.Println("no names")
	}
}
