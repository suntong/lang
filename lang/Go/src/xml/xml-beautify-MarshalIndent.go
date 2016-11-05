// https://golang.org/pkg/encoding/xml/#MarshalIndent
// https://play.golang.org/p/bMi4fsdWib

package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

func main() {
	type Address struct {
		City, State string
	}
	type Person struct {
		XMLName   xml.Name `xml:"person"`
		Id        int      `xml:"id,attr"`
		FirstName string   `xml:"name>first"`
		LastName  string   `xml:"name>last"`
		Age       int      `xml:"age"`
		Height    float32  `xml:"height,omitempty"`
		Married   bool
		Address
		Comment string `xml:",comment"`
	}

	v := &Person{Id: 13, FirstName: "John", LastName: "Doe", Age: 42}
	v.Comment = " Need more details. "
	v.Address = Address{"Hanga Roa", "Easter Island"}

	output, err := xml.MarshalIndent(v, "xx ", " ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	os.Stdout.Write(output)
}

/*

xx <person id="13">
xx  <name>
xx   <first>John</first>
xx   <last>Doe</last>
xx  </name>
xx  <age>42</age>
xx  <Married>false</Married>
xx  <City>Hanga Roa</City>
xx  <State>Easter Island</State>
xx  <!-- Need more details. -->
xx </person>

*/
