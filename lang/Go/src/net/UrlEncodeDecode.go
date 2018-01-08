////////////////////////////////////////////////////////////////////////////
// Porgram: UrlEncodeDecode.go
// Purpose: Go URL encoding and decoding demo
// Authors: Tong Sun (c) 2018, All rights reserved
////////////////////////////////////////////////////////////////////////////

// https://play.golang.org/p/Fb_qXktITKY

package main

import (
	"fmt"
	"html"
	"net/url"
)

func main() {

	// https://stackoverflow.com/a/13826910/2125837
	var Url *url.URL
	Url, err := url.Parse("http://www.example.com")
	if err != nil {
		panic("boom")
	}

	Url.Path += "/some/path/with/funny_characters?_or_not/中文测试/"
	parameters := url.Values{}
	parameters.Add("hello", "42")
	parameters.Add("hello", "54")
	parameters.Add("vegetable", "potato++tomato")
	Url.RawQuery = parameters.Encode()

	es := Url.String()
	fmt.Printf("Encoded URL: %s\n", es)

	// decode URLs
	fmt.Printf("\nDecoded URL: ")
	fmt.Println(url.QueryUnescape(es))
	fmt.Printf("\nHtml Decode:\n")
	fmt.Println(html.UnescapeString("http&#x3a;&#x2f;&#x2f;www.test.com/potato%2B%2Btomato"))
	fmt.Println(html.UnescapeString(es))

	Url, _ = url.Parse("http://www.example.com/some/path/with%2B%2B/%E4%B8%AD%E6%96%87%E6%B5%8B%E8%AF%95/")
	fmt.Printf("\nDecoded: %s", Url.Path)
}

/*

Encoded URL: http://www.example.com/some/path/with/funny_characters%3F_or_not/%E4%B8%AD%E6%96%87%E6%B5%8B%E8%AF%95/?hello=42&hello=54&vegetable=potato%2B%2Btomato

Decoded URL: http://www.example.com/some/path/with/funny_characters?_or_not/中文测试/?hello=42&hello=54&vegetable=potato++tomato <nil>

Html Decode:
http://www.test.com/potato%2B%2Btomato
http://www.example.com/some/path/with/funny_characters%3F_or_not/%E4%B8%AD%E6%96%87%E6%B5%8B%E8%AF%95/?hello=42&hello=54&vegetable=potato%2B%2Btomato

Decoded: /some/path/with++/中文测试/

*/
