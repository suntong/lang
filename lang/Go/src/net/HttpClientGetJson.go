////////////////////////////////////////////////////////////////////////////
// Porgram: CommandLineFlag
// Purpose: Go command line flags/switches/arguments demo
// Authors: Tong Sun (c) 2016, All rights reserved
// Credits: https://www.socketloop.com/tutorials/golang-unmarshal-json-from-http-response
////////////////////////////////////////////////////////////////////////////

package main

import (
	"bytes"
	"crypto/tls"
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

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get(*url)
	//resp, err := http.Get(*url)

	defer resp.Body.Close()

	if err != nil {
		panic(err)
	}

	// read json http response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf"))

	var jsonData interface{}

	err = json.Unmarshal(body, &jsonData) // here!

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

Without the bytes.TrimPrefix fix:

  $ go run HttpClientGetJson.go --url 'http://www.totoagriculture.org/weather/weather?longitude=0.2167&latitude=5.55&format=JSON&days=2&language=Portuguese'
  Get 'http://www.totoagriculture.org/weather/weather?longitude=0.2167&latitude=5.55&format=JSON&days=2&language=Portuguese'
  panic: invalid character 'i' looking for beginning of value

With the bytes.TrimPrefix fix:

  $ go run HttpClientGetJson.go --url 'http://www.totoagriculture.org/weather/weather?longitude=0.2167&latitude=5.55&format=JSON&days=2&language=Portuguese'
  Get 'http://www.totoagriculture.org/weather/weather?longitude=0.2167&latitude=5.55&format=JSON&days=2&language=Portuguese'
  map[date:2016-10-29 current-conditions:map[time:2016-10-29 summary:Quente, ...


*/

/*

Ref:

Unmarshal JSON from http response
https://www.socketloop.com/tutorials/golang-unmarshal-json-from-http-response

Parsing into an Interface
https://eager.io/blog/go-and-json/

Getting x509: certificate signed by unknown authority
https://groups.google.com/forum/#!topic/golang-nuts/v5ShM8R7Tdc

UTF8 BOM Character Confuses JSON Decoder
http://grokbase.com/t/gg/golang-nuts/137a49jmqe/go-nuts-utf8-bom-character-confuses-json-decoder

Got error "invalid character looking for beginning of value" from json.Unmarshal
http://qaoverflow.com/question/got-error-invalid-character-i-looking-for-beginning-of-value-from-json-unmarshal/

*/
