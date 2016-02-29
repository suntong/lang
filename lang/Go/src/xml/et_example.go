////////////////////////////////////////////////////////////////////////////
// Porgram: et_example.go
// Purpose: etree example
// Authors: Tong Sun (c) 2016, All rights reserved
// Credits: Brett Vickers (c) 2015 BSD
//          https://github.com/beevik/etree
//          https://github.com/beevik/etree/blob/master/example_test.go
////////////////////////////////////////////////////////////////////////////

package main

import (
	"os"
)

import (
	"github.com/beevik/etree"
)

func main() {
	ExampleDocument_creating()
	print("\n")
	ExamplePath()
}

// Create an etree Document, add XML entities to it, and serialize it
// to stdout.
func ExampleDocument_creating() {
	doc := etree.NewDocument()
	doc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)
	doc.CreateProcInst("xml-stylesheet", `type="text/xsl" href="style.xsl"`)

	people := doc.CreateElement("People")
	people.CreateComment("These are all known people")

	jon := people.CreateElement("Person")
	jon.CreateAttr("name", "Jon O'Reilly")

	sally := people.CreateElement("Person")
	sally.CreateAttr("name", "Sally")

	doc.Indent(2)
	doc.WriteTo(os.Stdout)
	// Output:
	// <?xml version="1.0" encoding="UTF-8"?>
	// <?xml-stylesheet type="text/xsl" href="style.xsl"?>
	// <People>
	//   <!--These are all known people-->
	//   <Person name="Jon O&apos;Reilly"/>
	//   <Person name="Sally"/>
	// </People>
}

func ExampleDocument_reading() {
	doc := etree.NewDocument()
	if err := doc.ReadFromFile("document.xml"); err != nil {
		panic(err)
	}
}

func ExamplePath() {
	xml := `<bookstore><book><title>Great Expectations</title>
      <author>Charles Dickens</author></book><book><title>Ulysses</title>
      <author>James Joyce</author></book></bookstore>`

	doc := etree.NewDocument()
	doc.ReadFromString(xml)
	for _, e := range doc.FindElements(".//book[author='Charles Dickens']") {
		book := etree.CreateDocument(e)
		book.Indent(2)
		book.WriteTo(os.Stdout)
	}
	// Output:
	// <book>
	//   <title>Great Expectations</title>
	//   <author>Charles Dickens</author>
	// </book>
}
