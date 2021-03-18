// Based on: http://stackoverflow.com/questions/32635943/

package main

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var (
	html = `<html><body><div class="outter-class">
        <h1 class="inner-class">
	        The string I need
	        <span class="other-class" >Some value I don't need</span>
	        <span class="other-class2" title="sometitle"></span>
        </h1>
        <h3>Some heading i don't need</h3>
        <div class="other-class3">
          Some div
        </div>
    </div></body></html>`
)

func main() {
	test1()
	fmt.Println("---\n")
	test2()
}

func test1() {
	r := strings.NewReader(html)
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		panic(err)
	}

	// h1 := doc.Find("h1")
	// h1.Children().Remove()
	// text := h1.Text()

	text := doc.Find("h1").Children().Remove().End().Text()

	text = strings.TrimSpace(text)
	fmt.Println(text)
}

func test2() {
	r := strings.NewReader(html)
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		panic(err)
	}

	b := doc.Find("body")
	b.Find("h1").Remove()
	b.Find("h3").Remove()

	fmt.Println(b.Html())
}
