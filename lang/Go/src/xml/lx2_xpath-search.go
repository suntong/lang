////////////////////////////////////////////////////////////////////////////
// Porgram: lx2_element-names.go
// Purpose: gokogiri xpath demo
// Authors: Tong Sun (c) 2016, All rights reserved
// Credits: ArunL, JimB
//          http://stackoverflow.com/a/25387687/2125837
////////////////////////////////////////////////////////////////////////////

/*

Goal:

Do a two-level xpath query, to get all divs with class="area" first, then
recursively get divs inside it with class="value"

*/

package main

import (
	"fmt"
	//"io/ioutil"

	"github.com/moovweb/gokogiri"
	"github.com/moovweb/gokogiri/xpath"
)

func main() {
	// content, _ := ioutil.ReadFile("index.html")
	//doc, _ := gokogiri.ParseHtml(content)
	content := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Document</title>
</head>
<body>
    <div class="head">
        <div class="area">
            <div class="value">10</div>
        </div>
        <div class="area">
            <div class="value">20</div>
        </div>
        <div class="area">
            <div class="value">30</div>
        </div>
    </div>
</body>
</html>`
	doc, _ := gokogiri.ParseHtml([]byte(content))
	defer doc.Free()

	{
		fmt.Println("Method 1, wrong results")
		xps := xpath.Compile("//div[@class='head']/div[@class='area']")
		xpw := xpath.Compile("//div[@class='value']")
		ss, _ := doc.Root().Search(xps)
		for _, s := range ss {
			ww, _ := s.Search(xpw)
			for _, w := range ww {
				fmt.Println(w.InnerHtml())
			}
		}
		/*
			10
			20
			30
			10
			20
			30
			10
			20
			30
		*/
	}

	{
		fmt.Println("\nMethod 2, correct results")
		xps := xpath.Compile("//div[@class='head']/div[@class='area']")
		xpw := xpath.Compile(".//div[@class='value']")

		// this works in your example case
		// xpw := xpath.Compile("div[@class='value']")
		// as does this
		// xpw := xpath.Compile("./div[@class='value']")

		ss, _ := doc.Root().Search(xps)
		for _, s := range ss {
			ww, _ := s.Search(xpw)
			for _, w := range ww {
				fmt.Println(w.InnerHtml())
			}
		}
		/*
			10
			20
			30
		*/
	}

}

/*

Explain:

An XPath search from any node ('//') can still search the entire tree.

If you want to search just the subtree, you can start the expression with a .
(assuming you still want descendant-or-self), otherwise use a exact path.

- // is short for /descendant-or-self::node()/, so on it's own the first slash indicated the tree root.
- The . is an alias for self::node() XPath.

*/
