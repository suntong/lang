// https://www.thepolyglotdeveloper.com/2020/02/interacting-with-a-graphql-api-with-golang/

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
	"time"
)

func main() {

	query := map[string]string{
		"query": `{ countries { code name } }`,
	}

	queryJson, err := json.Marshal(query)
	if err != nil {
		panic(err)
	}

	request, err := http.NewRequest("POST", "https://countries.trevorblades.com/", bytes.NewBuffer(queryJson))
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(request)
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}
	b, _ := httputil.DumpResponse(resp, true)
	fmt.Println(string(b))
}
