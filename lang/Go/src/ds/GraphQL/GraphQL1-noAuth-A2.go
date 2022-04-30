// Based on GraphQL1-noAuth-A.go, for
// https://studio.apollographql.com/public/SpaceX-pxxbxen/explorer

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
query Launches {
  launches {
    mission_name
    mission_id
    rocket {
      rocket_name
      rocket {
        company
        name
        mass {
          kg
        }
      }
    }
    launch_site {
      site_name
    }
    launch_date_local
  }
}
`,
	}
	queryJson, _ := json.Marshal(query)
	request, err := http.NewRequest("POST", "https://api.spacex.land/graphql/", bytes.NewBuffer(queryJson))
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
