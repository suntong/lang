////////////////////////////////////////////////////////////////////////////
// Porgram: parser03B
// Purpose: Go full-xml-parsing demo
// Authors: Tong Sun (c) 2015, All rights reserved
// Credits: Gulácsi Tamás http://play.golang.org/p/2oHEoq_PcH
//          https://groups.google.com/d/msg/golang-nuts/tf4aDQ1Hn_c/fVejKHUjBQAJ
////////////////////////////////////////////////////////////////////////////

package main

// An example streaming XML parser.

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"
	"strings"
)

// Here is an example article from the Wikipedia XML dump
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

func main() {
	inputReader := strings.NewReader(input)

	decoder := xml.NewDecoder(inputReader)
	for {
		// Read tokens from the XML document in a stream.
		t, _ := decoder.Token()
		if t == nil {
			break
		}
		// Inspect the type of the token just read.
		//fmt.Printf("\n<!--%#v-->\n", t)
		switch x := t.(type) {
		case xml.StartElement:
			fmt.Printf("<%s", x.Name.Local)
			for _, attr := range x.Attr {
				fmt.Printf(" %s=\"", attr.Name.Local)
				xml.EscapeText(os.Stdout, []byte(attr.Value))
				os.Stdout.Write([]byte{'"'})
			}
			fmt.Printf(">")
		case xml.CharData:
			xml.EscapeText(os.Stdout, bytes.TrimSpace(x))
		case xml.EndElement:
			fmt.Printf("</%s>\n", x.Name.Local)
		}

	}
}
