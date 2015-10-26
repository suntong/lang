////////////////////////////////////////////////////////////////////////////
// Porgram: parser03
// Purpose: Go xml parsing demo
// Authors: Tong Sun (c) 2015, All rights reserved
// Credits: https://github.com/dps/go-xml-parse/blob/master/go-xml-parse.go
//          https://en.wikipedia.org/wiki/Help:Export
////////////////////////////////////////////////////////////////////////////

package main

// An example streaming XML parser.

import (
	"encoding/xml"
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

var filter, _ = regexp.Compile("^file:.*|^talk:.*|^special:.*|^wikipedia:.*|^wiktionary:.*|^user:.*|^user_talk:.*")

// Here is an example article from the Wikipedia XML dump
//
// <page>
// 	<title>Apollo 11</title>
//      <redirect title="Foo bar" />
// 	...
// 	<revision>
// 	...
// 	  <text xml:space="preserve">
// 	  {{Infobox Space mission
// 	  |mission_name=&lt;!--See above--&gt;
// 	  |insignia=Apollo_11_insignia.png
// 	...
// 	  </text>
// 	</revision>
// </page>
//
// Note how the tags on the fields of Page and Redirect below
// describe the XML schema structure.
var input = `
<mediawiki xml:lang="en">
  <page>
    <title>Page title</title>
    <restrictions>edit=sysop:move=sysop</restrictions>
    <revision>
      <timestamp>2001-01-15T13:15:00Z</timestamp>
      <contributor><username>Foobar</username></contributor>
      <comment>I have just one thing to say!</comment>
      <text>A bunch of [[text]] here.</text>
      <minor />
    </revision>
    <revision>
      <timestamp>2001-01-15T13:10:27Z</timestamp>
      <contributor><ip>10.0.0.2</ip></contributor>
      <comment>new!</comment>
      <text>An earlier [[revision]].</text>
    </revision>
  </page>
  
  <page>
    <title>Talk:Page title</title>
    <revision>
      <timestamp>2001-01-15T14:03:00Z</timestamp>
      <contributor><ip>10.0.0.2</ip></contributor>
      <comment>hey</comment>
      <text>WHYD YOU LOCK PAGE??!!! i was editing that jerk</text>
    </revision>
  </page>
</mediawiki>`

type Redirect struct {
	Title string `xml:"title,attr"`
}

type Page struct {
	Title string   `xml:"title"`
	Redir Redirect `xml:"redirect"`
	Text  string   `xml:"revision>text"`
}

func CanonicalizeTitle(title string) string {
	can := strings.ToLower(title)
	can = strings.Replace(can, " ", "_", -1)
	can = url.QueryEscape(can)
	return can
}

func WritePage(title string, text string) {
	fmt.Printf("Title: %s\nText: %s\n\n", title, text)
}

func main() {
	inputReader := strings.NewReader(input)

	decoder := xml.NewDecoder(inputReader)
	total := 0
	var inElement string
	for {
		// Read tokens from the XML document in a stream.
		t, _ := decoder.Token()
		if t == nil {
			break
		}
		// Inspect the type of the token just read.
		switch se := t.(type) {
		case xml.StartElement:
			// If we just read a StartElement token
			inElement = se.Name.Local
			// ...and its name is "page"
			if inElement == "page" {
				var p Page
				// decode a whole chunk of following XML into the
				// variable p which is a Page (se above)
				decoder.DecodeElement(&p, &se)

				// Do some stuff with the page.
				p.Title = CanonicalizeTitle(p.Title)
				m := filter.MatchString(p.Title)
				if !m && p.Redir.Title == "" {
					WritePage(p.Title, p.Text)
					total++
				}
			}
		default:
		}

	}

	fmt.Printf("Total articles: %d \n", total)
}
