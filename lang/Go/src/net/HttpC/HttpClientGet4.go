package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	// Enforce IPv4, Create a transport object
	transport := &http.Transport{
		Dial: (&net.Dialer{
			DualStack: false, // This ensures only IPv4 is used
		}).Dial,
	}

	// Create an HTTP client with the custom transport
	client := &http.Client{
		Transport: transport,
	}

	// Create an HTTP GET request
	req, err := http.NewRequest("GET", os.Getenv("URL"), nil)
	if err != nil {
		log.Fatal(err)
	}
	// Send the request using the client
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}

/*

URL=https://httpbin.org/get go run HttpClientGet4.go
URL=https://ifconfig.me/ go run HttpClientGet4.go # still NOK

*/
