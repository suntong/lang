package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

var serverAddr = "http://example.com/"

func main() {
	serverAddr = "https://top.baidu.com/board?tab=realtime"

	req, err := http.NewRequest(http.MethodGet, serverAddr, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("test-header", "test-header-value")

	reqDump, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("REQUEST:\n%s", string(reqDump))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	respDump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("RESPONSE:\n%s", string(respDump))
}
