////////////////////////////////////////////////////////////////////////////
// Porgram: htmlParserTokens.go
// Purpose: Go html token parsing demo
// Authors: Tong Sun (c) 2017, All rights reserved
// Credits: https://godoc.org/golang.org/x/net/html
//          https://play.golang.org/p/0MRSefJ_-E
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

type TokenVisitor interface {
	VisitToken()
}

func WalkBody(of TokenVisitor) {
	// here you call of.VisitToken()
}

type MyTokenizer1 struct {
	*html.Tokenizer
}

func (the MyTokenizer1) VisitToken() {

}

type MyTokenizer2 struct {
	*html.Tokenizer
}

func (the MyTokenizer2) VisitToken() {

}

func main() {

	HTMLString := `<!DOCTYPE html>
  <html itemscope itemtype="http://schema.org/QAPage">
  <head>
  <title>go - Golang parse HTML, extract all content with &lt;body&gt; &lt;/body&gt; tags - Stack Overflow</title>
    <link rel="shortcut icon" href="//cdn.sstatic.net/Sites/stackoverflow/img/favicon.ico?v=4f32ecc8f43d">
    <link rel="apple-touch-icon image_src" href="//cdn.sstatic.net/Sites/stackoverflow/img/apple-touch-icon.png?v=c78bd457575a">
    <link rel="search" type="application/opensearchdescription+xml" title="Stack Overflow" href="/opensearch.xml">
    <meta name="twitter:card" content="summary">
    <meta name="twitter:domain" content="stackoverflow.com"/>
    <meta property="og:type" content="website" />
    </head>
<body class="template-blog">
<nav class="navigation">
<div class="navigation__container container">
<a class="navigation__logo" href="/">
<h1>Foobar</h1>
</a>
<ul class="navigation__menu">
<li><a href="/tags/">Topics</a></li>
<li><a href="/about">About</a></li>
</ul>
</div>`

	// response, err := http.Get("https://www.coastal.edu/scs/employee")
	// checkError(err)
	// defer resp.Body.Close()
	// z := html.NewTokenizer(response.Body)
	r := strings.NewReader(HTMLString)
	// z := html.NewTokenizer(r)
	z := &MyTokenizer1{html.NewTokenizer(r)}

	htmlWalk(z)
}

func htmlWalk(z TokenVisitor) error {
	depth := 0
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			return z.Err()
		case html.TextToken:
			if depth > 0 {
				// emitBytes should copy the []byte it receives,
				// if it doesn't process it immediately.
				// emitBytes(z.Text())

				text := strings.TrimSpace(string(z.Text()))
				if text != "" {
					z.printElmt(depth, text)
				}
			}
		case html.StartTagToken, html.EndTagToken:
			tn, _ := z.TagName()
			tag := strings.ToLower(string(tn))
			if tt == html.StartTagToken {
				if tag == "body" {
					depth = 0
				}
				z.printElmt(depth, tag)
				depth++
			} else {
				depth--
			}
		}
	}
}

func (z MyTokenizer1) printElmt(depth int, s string) {
	for n := 0; n < depth; n++ {
		fmt.Print("  ")
	}
	fmt.Println(s)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}

/*

$ go run htmlParserTokens.go
html
  head
    title
      go - Golang parse HTML, extract all content with <body> </body> tags - Stack Overflow
    link
      link
        link
          meta
body
  nav
    div
      a
        h1
          Foobar
      ul
        li
          a
            Topics
        li
          a
            About

*/
