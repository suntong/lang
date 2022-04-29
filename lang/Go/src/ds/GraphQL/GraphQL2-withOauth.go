// https://kyrcha.info/2019/10/15/sending-graphql-queries-using-http-client-in-go

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http/httputil"
	"os"
	"strings"

	"golang.org/x/oauth2"
)

type graphQLRequest struct {
	Query     string `json:"query"`
	Variables string `json:"variables"`
}

func main() {
	client := oauth2.NewClient(
		context.TODO(),
		oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
		))

	query := `query {
			repository(owner:"octocat", name:"Hello-World") {
			  issues(last:20, states:CLOSED) {
				edges {
				  node {
					title
					url
					labels(first:5) {
					  edges {
						node {
						  name
						}
					  }
					}
				  }
				}
			  }
			}
		  }`

	gqlMarshalled, err := json.Marshal(graphQLRequest{Query: query})
	if err != nil {
		panic(err)
	}

	resp, err := client.Post("https://api.github.com/graphql", "application/json", strings.NewReader(string(gqlMarshalled)))
	if err != nil {
		panic(err)
	}
	b, _ := httputil.DumpResponse(resp, true)
	fmt.Println(string(b))

	query = `query($number_of_repos:Int!) {
		viewer {
		  name
		   repositories(last: $number_of_repos) {
			 nodes {
			   name
			 }
		   }
		 }
	  }`

	variables := `variables {
		"number_of_repos": 3
	 }`

	gqlMarshalled, err = json.Marshal(graphQLRequest{Query: query, Variables: variables})
	if err != nil {
		panic(err)
	}

	b, _ = httputil.DumpResponse(resp, true)
	fmt.Println(string(b))

}
