/*
 * Author, Copyright: Oleg Borodin <onborodin@gmail.com>
 * https://wiki.unix7.org/go/graphql
 */

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
		"query": `
            {
                hello(id: 5) {
                }
            }
        `,
	}
	queryJson, _ := json.Marshal(query)
	request, err := http.NewRequest("POST", "http://localhost:8080/query", bytes.NewBuffer(queryJson))
	client := &http.Client{Timeout: time.Second * 10}
	response, err := client.Do(request)
	defer response.Body.Close()
	if err != nil {
		fmt.Printf("error: %s\n", err)
	}
	data, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(data))
}
