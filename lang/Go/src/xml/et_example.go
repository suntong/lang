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
	"fmt"
	"os"
)

import (
	//"gopkg.in/beevik/etree.v0"
	"github.com/beevik/etree"
)

////////////////////////////////////////////////////////////////////////////
// Function definitions

//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
// Function main

func main() {
	ExampleDocument_creating()
	fmt.Println()
	ExamplePath()
	fmt.Println()
	ExampleAddElement()
	fmt.Println()

	ProcessingEA()
	fmt.Println()
	PathQueries()
	fmt.Println()

	DemoRemoveElement()
	fmt.Println()
	DemoFindElements()
	fmt.Println()
}

func readXml(xml string) *etree.Document {
	doc := etree.NewDocument()
	doc.ReadFromString(xml)
	return doc
}

//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
// From https://github.com/beevik/etree/blob/master/example_test.go

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

var xml = `<bookstore><book><title>Great Expectations</title>
      <author>Charles Dickens</author></book><book><title>Ulysses</title>
      <author>James Joyce</author></book></bookstore>`

func ExamplePath() {
	doc := etree.NewDocument()
	doc.ReadFromString(xml)
	for _, e := range doc.FindElements(".//book[author='Charles Dickens']") {
		book := etree.NewDocument()
		book.SetRoot(e.Copy())
		book.Indent(2)
		book.WriteTo(os.Stdout)
	}
	// Output:
	// <book>
	//   <title>Great Expectations</title>
	//   <author>Charles Dickens</author>
	// </book>
}

func ExampleAddElement() {
	docR := readXml(bookstore)
	root := docR.SelectElement("bookstore")

	docA := etree.NewDocument()
	docA.ReadFromString(xml)
	docA.Indent(2)
	docA.WriteTo(os.Stdout)

	// To directly add docA under root
	//root.AddElement(docA.Root())

	// To add all book nodes under docA to root
	for _, e := range docA.FindElements(".//book") {
		// add (e) under root
		root.AddElement(e)
	}

	fmt.Println()
	docR.Indent(2)
	docR.WriteTo(os.Stdout)

	fmt.Println()
	docA.Indent(2)
	docA.WriteTo(os.Stdout)
}

//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
// From https://github.com/beevik/etree/

var bookstore = `
<bookstore xmlns:p="urn:schemas-books-com:prices">

  <book category="COOKING">
    <title lang="en">Everyday Italian</title>
    <author>Giada De Laurentiis</author>
    <year>2005</year>
    <p:price>30.00</p:price>
  </book>

  <book category="CHILDREN">
    <title lang="en">Harry Potter</title>
    <author>J K. Rowling</author>
    <year>2005</year>
    <p:price>29.99</p:price>
  </book>

  <book category="WEB">
    <title lang="en">XQuery Kick Start</title>
    <author>James McGovern</author>
    <author>Per Bothner</author>
    <author>Kurt Cagle</author>
    <author>James Linn</author>
    <author>Vaidyanathan Nagarajan</author>
    <year>2003</year>
    <p:price>49.99</p:price>
  </book>

  <book category="WEB">
    <title lang="en">Learning XML</title>
    <author>Erik T. Ray</author>
    <year>2003</year>
    <p:price>39.95</p:price>
  </book>

</bookstore>
`

// Processing elements and attributes
// Illustrates some ways to access elements and attributes using simple etree queries
func ProcessingEA() {
	doc := readXml(bookstore)

	root := doc.SelectElement("bookstore")
	fmt.Println("ROOT element:", root.Tag)

	for _, book := range root.SelectElements("book") {
		fmt.Println("CHILD element:", book.Tag)
		if title := book.SelectElement("title"); title != nil {
			lang := title.SelectAttrValue("lang", "unknown")
			fmt.Printf("  TITLE: %s (%s)\n", title.Text(), lang)
		}
		for _, attr := range book.Attr {
			fmt.Printf("  ATTR: %s=%s\n", attr.Key, attr.Value)
		}
	}
}

/*

Output:

ROOT element: bookstore
CHILD element: book
  TITLE: Everyday Italian (en)
  ATTR: category=COOKING
CHILD element: book
  TITLE: Harry Potter (en)
  ATTR: category=CHILDREN
CHILD element: book
  TITLE: XQuery Kick Start (en)
  ATTR: category=WEB
CHILD element: book
  TITLE: Learning XML (en)
  ATTR: category=WEB

*/

// Path queries
func PathQueries() {
	doc := readXml(bookstore)

	// select all book titles that fall into the category of 'WEB'
	for _, t := range doc.FindElements("//book[@category='WEB']/title") {
		fmt.Println("Title:", t.Text())
	}
	fmt.Println()

	// finds the first book element under the bookstore element
	// and outputs the tag and text of all of its child elements
	{
		for _, e := range doc.FindElements("./bookstore/book[1]/*") {
			fmt.Printf("%s: %s\n", e.Tag, e.Text())
		}
	}
	fmt.Println()

	// finds all books with a price of 49.99 and outputs their titles.
	{
		path := etree.MustCompilePath("./bookstore/book[p:price='49.99']/title")
		for _, e := range doc.FindElementsPath(path) {
			fmt.Println(e.Text())
		}
	}
}

/*

Output:

Title: XQuery Kick Start
Title: Learning XML

title: Everyday Italian
author: Giada De Laurentiis
year: 2005
price: 30.00

XQuery Kick Start

*/

//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
// More examples, according to https://godoc.org/github.com/beevik/etree

func DemoRemoveElement() {
	// doc := readXml(xml)
	// for _, e := range doc.FindElements(".//book") {

	doc := readXml(bookstore)
	for _, e := range doc.FindElements(".//book[@category='WEB']") {

		// FindElement returns the first element matched by the XPath string
		a := e.FindElement(".//author")

		// RemoveElement removes the given child element
		e.RemoveElement(a)

		an := e.CreateElement("author")
		an.CreateComment("removed")

		book := etree.NewDocument()
		book.SetRoot(e.Copy())
		book.Indent(2)
		book.WriteTo(os.Stdout)
	}
}

/*

Output:

<book category="WEB">
  <title lang="en">XQuery Kick Start</title>
  <author>Per Bothner</author>
  <author>Kurt Cagle</author>
  <author>James Linn</author>
  <author>Vaidyanathan Nagarajan</author>
  <year>2003</year>
  <p:price>49.99</p:price>
  <author>
    <!--removed-->
  </author>
</book>
<book category="WEB">
  <title lang="en">Learning XML</title>
  <year>2003</year>
  <p:price>39.95</p:price>
  <author>
    <!--removed-->
  </author>
</book>

With,

	doc := readXml(xml)
	for _, e := range doc.FindElements(".//book") {

Output:

<book>
  <title>Great Expectations</title>
  <author>
    <!--removed-->
  </author>
</book>
<book>
  <title>Ulysses</title>
  <author>
    <!--removed-->
  </author>
</book>

*/

func DemoFindElements() {

	doc := readXml(bookstore)
	doc.Indent(2)
	doc.WriteTo(os.Stdout)
	fmt.Println()

	for _, e := range doc.FindElements(".//book") {
		p := e.FindElement(".//p:price")
		e.RemoveElement(p)
		for _, a := range e.FindElements(".//author") {
			e.RemoveElement(a)
		}
		for _, t := range e.FindElements(".//title") {
			// creates an attribute and adds it to the receiving element
			// may be prefixed by a namespace and a colon.
			t.CreateAttr("bk:version", "released")
			// If an attribute with the key already exists, its value is replaced
			t.CreateAttr("lang", "en_US")

			// SetText replaces an element's subsidiary CharData text with a new string
			t.SetText(t.FindElement("..//[@category]").
				SelectAttrValue("category", "unknown") + ": " + t.Text())
		}
		// he tag may be prefixed by a namespace and a colon.
		pub := e.CreateElement("bk:publisher")
		pub.SetText("Unspecified")
	}
	doc.WriteTo(os.Stdout)
}

/*

Output:

<bookstore xmlns:p="urn:schemas-books-com:prices">
  <book category="COOKING">
    <title lang="en">Everyday Italian</title>
    <author>Giada De Laurentiis</author>
    <year>2005</year>
    <p:price>30.00</p:price>
  </book>
  <book category="CHILDREN">
    <title lang="en">Harry Potter</title>
    <author>J K. Rowling</author>
    <year>2005</year>
    <p:price>29.99</p:price>
  </book>
  <book category="WEB">
    <title lang="en">XQuery Kick Start</title>
    <author>James McGovern</author>
    <author>Per Bothner</author>
    <author>Kurt Cagle</author>
    <author>James Linn</author>
    <author>Vaidyanathan Nagarajan</author>
    <year>2003</year>
    <p:price>49.99</p:price>
  </book>
  <book category="WEB">
    <title lang="en">Learning XML</title>
    <author>Erik T. Ray</author>
    <year>2003</year>
    <p:price>39.95</p:price>
  </book>
</bookstore>

<bookstore xmlns:p="urn:schemas-books-com:prices">
  <book category="COOKING">
    <title lang="en_US" bk:version="released">COOKING: Everyday Italian</title>

    <year>2005</year>

  <bk:publisher>Unspecified</bk:publisher></book>
  <book category="CHILDREN">
    <title lang="en_US" bk:version="released">CHILDREN: Harry Potter</title>

    <year>2005</year>

  <bk:publisher>Unspecified</bk:publisher></book>
  <book category="WEB">
    <title lang="en_US" bk:version="released">WEB: XQuery Kick Start</title>





    <year>2003</year>

  <bk:publisher>Unspecified</bk:publisher></book>
  <book category="WEB">
    <title lang="en_US" bk:version="released">WEB: Learning XML</title>

    <year>2003</year>

  <bk:publisher>Unspecified</bk:publisher></book>
</bookstore>


*/

/*

//
//
func () {
	doc := readXml()

}

/*

Output:


*/
