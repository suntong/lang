// https://stackoverflow.com/questions/24455147/how-do-i-send-a-json-string-in-a-post-request-in-go

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	query := map[string]string{
		"query": `{ countries { code name } }`,
	}
	queryJson, _ := json.Marshal(query)
	request, err := http.NewRequest("POST", "https://countries.trevorblades.com/", bytes.NewBuffer(queryJson))
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: time.Second * 10}
	response, err := client.Do(request)
	defer response.Body.Close()
	if err != nil {
		fmt.Printf("error: %s\n", err)
	}
	data, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(data))
}
