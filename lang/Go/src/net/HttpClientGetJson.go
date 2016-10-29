////////////////////////////////////////////////////////////////////////////
// Porgram: CommandLineFlag
// Purpose: Go command line flags/switches/arguments demo
// Authors: Tong Sun (c) 2016, All rights reserved
// Credits: https://www.socketloop.com/tutorials/golang-unmarshal-json-from-http-response
////////////////////////////////////////////////////////////////////////////

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	url = flag.String("url", "https://api.bitfinex.com/v1/ticker/btcusd", "url that returns json")
)

func main() {
	flag.Parse()

	fmt.Printf("Get '%s'\n", *url)

	resp, err := http.Get(*url)

	defer resp.Body.Close()

	if err != nil {
		panic(err)
	}

	// read json http response
	jsonDataFromHttp, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	var jsonData interface{}

	err = json.Unmarshal([]byte(jsonDataFromHttp), &jsonData) // here!

	if err != nil {
		panic(err)
	}

	// test struct data
	fmt.Println(jsonData)
}

/*

$ go run HttpClientGetJson.go
Get 'https://api.bitfinex.com/v1/ticker/btcusd'
map[mid:715.165 bid:715.15 ask:715.18 last_price:715.15 timestamp:1477750218.066481764]

*/

/*

Ref:

Unmarshal JSON from http response
https://www.socketloop.com/tutorials/golang-unmarshal-json-from-http-response

Parsing into an Interface
https://eager.io/blog/go-and-json/

Getting x509: certificate signed by unknown authority
https://groups.google.com/forum/#!topic/golang-nuts/v5ShM8R7Tdc

*/
