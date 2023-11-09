package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
)

/*

   Demo using net/http/httputil.

   Note that /net/http/httptrace/ is more for time/tracing,
    start from DNS lookup and go all the way to receiving the response.
   Final result can be in JSON format. See
   https://www.inanzzz.com/index.php/post/pzas/tracing-and-debugging-http-client-requests-within-golang

*/

const serverAddr = "https://httpbin.org/get"

func main() {
	testGet()
	testPut()
}

func testGet() {
	// https://gosamples.dev/print-http-request-response/
	fmt.Println("== http get test")
	req, err := http.NewRequest(http.MethodGet, serverAddr, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
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

func testPut() {
	fmt.Println("\n== http post test")
	jsonStr := []byte(`{
  "mock_data": "true",
  "ip_address": "92.188.61.181",
  "email": "user@example.com",
  "user_agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_6_4) AppleWebKit/534.30 (KHTML, like Gecko) Chrome/12.0.742.100 Safari/534.30",
  "url": "http://example.com/"
}`)

	req, err := http.NewRequest("POST", "https://httpbin.org/post", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("test-header", "test-header-value")

	reqDump, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("REQUEST:\n%s", string(reqDump))

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	log.Printf(sb)
}
