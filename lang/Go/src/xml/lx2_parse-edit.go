////////////////////////////////////////////////////////////////////////////
// Porgram: lx2_parse-namespace.go
// Purpose: parse xml with a namespace using gokogiri (libxml2)
// Authors: Tong Sun (c) 2016, All rights reserved
// Credits: Zhigang
//          http://www.moovweb.com/blog/gokogiri-the-best-way-to-parse-xml-in-go/
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"github.com/moovweb/gokogiri"
)

func main() {
	// Parse even this bad bit of HTML and make it valid
	html := "<h2>I am so malformatted</h2>"
	doc, _ := gokogiri.ParseHtml([]byte(html))
	defer doc.Free()

	header := doc.Root().FirstChild().FirstChild()
	header.SetName("h1")

	fmt.Println(doc.String())
}

/*

From

Gokogiri – the best way to parse XML in Go
http://corpsite-prod.w62h-y2vd.accessdomain.com/blog/tag/go-lang/
http://www.moovweb.com/blog/gokogiri-the-best-way-to-parse-xml-in-go/

The HTML parser fixes up the document and makes it valid. So, the output looks like this:

        <!DOCTYPE html PUBLIC "-//W3C//DTD HTML 4.0 Transitional//EN" "http://www.w3.org/TR/REC-html40/loose.dtd">
        <html><body><h1>I am so malformatted</h1></body></html>

The h2 changed to an h1 and this is now a valid document!


From

xpath like selectors
https://groups.google.com/d/msg/golang-nuts/JdTrcAtJ39U/gMI9_eCNjHwJ

... to quickly select and match (nested) XML content:

func Parse(src string) {
	doc, err := htmlp.Parse([]byte(src), htmlp.DefaultEncodingBytes, nil, htmlp.DefaultParseOption, htmlp.DefaultEncodingBytes)
	if err != nil { return }
	defer doc.Free()

	n,_ := doc.Root().Search(`//a[starts-with(@href,"/recipe-")]`); if len(n)<1 { return }
	for k, v := range n { ... }

}



*/
