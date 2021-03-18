// From https://gist.github.com/alexmullins/ac9581e106eb6b1a33ac
// Stackoverflow Answer
// Original Question: http://stackoverflow.com/questions/32635943/get-text-from-div-without-child-elements/32640258#32640258

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
        <div class="other-class3">
            <h3>Some heading i don't need</h3>
        </div>
    </div></body></html>`
)

func main() {
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
