package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	httpClient *http.Client
	baseURL    string
}

type SlothfulMessage struct {
	Message string `json:"message"`
}

func NewClient(httpClient *http.Client, baseURL string) Client {
	return Client{
		httpClient: httpClient,
		baseURL:    baseURL,
	}
}

func (c *Client) GetSlothfulMessage() (*SlothfulMessage, error) {
	res, err := c.httpClient.Get(c.baseURL + "/sloth")
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"got status code %d", res.StatusCode,
		)
	}

	var m SlothfulMessage
	if err := json.NewDecoder(res.Body).Decode(&m); err != nil {
		return nil, err
	}
	return &m, nil
}

func handleSlothfulMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(`{"message": "Stay slothful!"}`))
}

func appRouter() http.Handler {
	rt := http.NewServeMux()
	rt.HandleFunc("/sloth", handleSlothfulMessage)
	return rt
}

func main() { http.ListenAndServe(":1123", appRouter()) }
