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

	"github.com/beevik/etree"
	"github.com/suntong/testing"
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

	ProcessingEA()
	fmt.Println()
	PathQueries()
	fmt.Println()

	DemoFindElements()
	fmt.Println()

	TestOrgTests()
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

	book := etree.NewDocument()
	for _, e := range doc.FindElements(".//book[author='Charles Dickens']") {
		book.AddChild(e)
	}
	book.Indent(2)
	book.WriteTo(os.Stdout)
	// Output:
	// <book>
	//   <title>Great Expectations</title>
	//   <author>Charles Dickens</author>
	// </book>
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

func DemoFindElements() {

	doc := readXml(bookstore)
	doc.Indent(2)
	doc.WriteTo(os.Stdout)
	fmt.Println()

	for _, e := range doc.FindElements(".//book") {
		p := e.FindElement(".//p:price")
		e.RemoveChild(p)
		for _, a := range e.FindElements(".//author") {
			e.RemoveChild(a)
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

//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
func TestOrgTests() {
	var t *testing.T = testing.NewT()
	TestCopy(t)
	t.Report()
	TestInsertChild(t)
	t.Report()
	TestAddChild(t)
	t.Report()
}

func checkEq(t *testing.T, got, want string) {
	if got == want {
		return
	}
	t.Errorf(
		"etree: unexpected result.\nGot:\n%s\nWanted:\n%s\n",
		got, want)
}

//--------------------------------------------------------------------------
func TestCopy(t *testing.T) {
	s := `<store>
	<book lang="en">
		<title>Great Expectations</title>
		<author>Charles Dickens</author>
	</book>
</store>`

	doc := etree.NewDocument()
	err := doc.ReadFromString(s)
	if err != nil {
		t.Fatal("etree: incorrect ReadFromString result")
	}

	s1, err := doc.WriteToString()
	if err != nil {
		t.Error("etree: incorrect WriteToString result")
	}

	doc2 := doc.Copy()
	s2, err := doc2.WriteToString()
	if err != nil {
		t.Error("etree: incorrect Copy result")
	}

	if s1 != s2 {
		t.Error("etree: mismatched Copy result")
		t.Error("wanted:\n" + s1)
		t.Error("got:\n" + s2)
	}

	e1 := doc.FindElement("./store/book/title")
	e2 := doc2.FindElement("./store/book/title")
	if e1 == nil || e2 == nil {
		t.Error("etree: incorrect FindElement result")
	}
	if e1 == e2 {
		t.Error("etree: incorrect FindElement result")
	}

	e1.Parent().RemoveChild(e1)
	s1, _ = doc.WriteToString()
	s2, _ = doc2.WriteToString()
	if s1 == s2 {
		t.Error("etree: incorrect result after RemoveElement")
	}
}

//--------------------------------------------------------------------------
func TestInsertChild(t *testing.T) {
	testdoc := `<book lang="en">
  <t:title>Great Expectations</t:title>
  <author>Charles Dickens</author>
</book>
`

	doc := etree.NewDocument()
	err := doc.ReadFromString(testdoc)
	if err != nil {
		t.Fatal("etree ReadFromString: " + err.Error())
	}

	year := etree.NewElement("year")
	year.SetText("1861")

	book := doc.FindElement("//book")
	book.InsertChild(book.SelectElement("t:title"), year)

	expected1 := `<book lang="en">
  <year>1861</year>
  <t:title>Great Expectations</t:title>
  <author>Charles Dickens</author>
</book>
`
	doc.Indent(2)
	s1, _ := doc.WriteToString()
	checkEq(t, s1, expected1)

	book.RemoveChild(year)
	book.InsertChild(book.SelectElement("author"), year)

	expected2 := `<book lang="en">
  <t:title>Great Expectations</t:title>
  <year>1861</year>
  <author>Charles Dickens</author>
</book>
`
	doc.Indent(2)
	s2, _ := doc.WriteToString()
	checkEq(t, s2, expected2)

	book.RemoveChild(year)
	book.InsertChild(book.SelectElement("UNKNOWN"), year)

	expected3 := `<book lang="en">
  <t:title>Great Expectations</t:title>
  <author>Charles Dickens</author>
  <year>1861</year>
</book>
`
	doc.Indent(2)
	s3, _ := doc.WriteToString()
	checkEq(t, s3, expected3)

	book.RemoveChild(year)
	book.InsertChild(nil, year)

	expected4 := `<book lang="en">
  <t:title>Great Expectations</t:title>
  <author>Charles Dickens</author>
  <year>1861</year>
</book>
`
	doc.Indent(2)
	s4, _ := doc.WriteToString()
	checkEq(t, s4, expected4)
}

//--------------------------------------------------------------------------
func TestAddChild(t *testing.T) {
	testdoc := `<book lang="en">
  <t:title>Great Expectations</t:title>
  <author>Charles Dickens</author>
`
	doc1 := etree.NewDocument()
	err := doc1.ReadFromString(testdoc)
	if err != nil {
		t.Fatal("etree ReadFromString: " + err.Error())
	}

	doc2 := etree.NewDocument()
	root := doc2.CreateElement("root")

	for _, e := range doc1.FindElements("//book/*") {
		root.AddChild(e)
	}

	expected1 := `<book lang="en"/>
`
	doc1.Indent(2)
	s1, _ := doc1.WriteToString()
	checkEq(t, s1, expected1)

	expected2 := `<root>
  <t:title>Great Expectations</t:title>
  <author>Charles Dickens</author>
</root>
`
	doc2.Indent(2)
	s2, _ := doc2.WriteToString()
	checkEq(t, s2, expected2)
}

/*

Output:


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
