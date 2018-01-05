////////////////////////////////////////////////////////////////////////////
// Porgram: parser01
// Purpose: Go xml parsing demo
// Authors: Tong Sun (c) 2018, All rights reserved
// Credits: http://blog.studygolang.com/tag/xml/
////////////////////////////////////////////////////////////////////////////

package main

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type Tr struct {
	Td      []*Td    `xml:"http://www.w3.org/1999/xhtml td,omitempty" json:"td,omitempty"`
	XMLName xml.Name `xml:"http://www.w3.org/1999/xhtml tr,omitempty" json:"tr,omitempty"`
}

type Td struct {
	Colspan string   `xml:"colspan,attr"  json:",omitempty"` // maxLength=1
	Rowspan string   `xml:"rowspan,attr"  json:",omitempty"` // maxLength=1
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://www.w3.org/1999/xhtml td,omitempty" json:"td,omitempty"`
}

func main() {
	input := `
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">
<html xmlns="http://www.w3.org/1999/xhtml" xml:lang="en" lang="en">
<table cellpadding="1" cellspacing="1" border="1">
<tr>
	<td>open</td>
	<td>/</td>
	<td></td>
</tr>
<tr><td>1</td><td>2</td><td></td></tr>
<tr><td>a</td><td>b</td><td>c</td></tr>
</table>
</html>`

	inputReader := strings.NewReader(input)

	decoder := xml.NewDecoder(inputReader)
	for {
		token, _ := decoder.Token()
		if token == nil {
			break
		}
		switch se := token.(type) {
		case xml.StartElement:
			if se.Name.Local == "tr" {
				var item Tr
				decoder.DecodeElement(&item, &se)
				tr := item
				fmt.Printf("%#v %#v\n", tr, tr.Td)
				fmt.Printf("%#v %#v\n", tr.Td[1], tr.Td[2])
			}
		}
	}
}
