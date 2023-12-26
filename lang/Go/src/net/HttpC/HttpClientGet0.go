package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get(os.Getenv("URL"))
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

URL=https://httpbin.org/get go run HttpClientGet0.go
URL=https://ifconfig.me/ go run HttpClientGet0.go

*/
