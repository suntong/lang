////////////////////////////////////////////////////////////////////////////
// Porgram: htmlParser.go
// Purpose: Parser HTML with Go x/net/html lib
// Authors: Tong Sun (c) 2017, All rights reserved
// Credits: https://godoc.org/golang.org/x/net/html#example-Parse
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"log"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	s := `<p>Links:</p><ul><li><a href="foo">Foo</a><li><a href="/bar/baz">BarBaz</a></ul>`
	doc, err := html.Parse(strings.NewReader(s))
	if err != nil {
		log.Fatal(err)
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					fmt.Println(a.Val)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
}
