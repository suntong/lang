////////////////////////////////////////////////////////////////////////////
// Porgram: jsonfile-http.go
// Purpose: GO jsonfile package HTTP JSON handling demo
// Authors: Tong Sun (c) 2016, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"

	// json http stream
	jhs "github.com/go-jsonfile/jsonfile/ext"
)

func main() {
	var data interface{}
	jhs.GetJSON("http://headers.jsontest.com/", &data)

	fmt.Printf("%+v\n", data)
}

/*

map[X-Cloud-Trace-Context:ae6b0b67215738632b6ee06849e85eb4/10686516943854567290 Host:headers.jsontest.com X-Forwarded-For:192.168.2.102 Via:1.1 mypc (squid/3.5.12) User-Agent:Go-http-client/1.1 Cache-Control:max-age=259200]

*/
