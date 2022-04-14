package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandleSlothfulMessage(t *testing.T) {
	wr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/sloth", nil)

	handleSlothfulMessage(wr, req)
	if wr.Code != http.StatusOK {
		t.Errorf("got HTTP status code %d, expected 200", wr.Code)
	}

	if !strings.Contains(wr.Body.String(), "Stay slothful!") {
		t.Errorf(
			`response body "%s" does not contain "Stay slothful!"`,
			wr.Body.String(),
		)
	}
}

func TestGetSlothfulMessage(t *testing.T) {
	router := http.NewServeMux()
	router.HandleFunc("/sloth", handleSlothfulMessage)

	svr := httptest.NewServer(router)
	defer svr.Close()

	c := NewClient(http.DefaultClient, svr.URL)
	m, err := c.GetSlothfulMessage()
	if err != nil {
		t.Fatalf("error in GetSlothfulMessage: %v", err)
	}
	if m.Message != "Stay slothful!" {
		t.Errorf(
			`message %s should contain string "Sloth"`,
			m.Message,
		)
	}
}

/*

Test your Go web apps with httptest
https://dev.to/salesforceeng/test-your-go-web-apps-with-httptest-26mc

   Not only does httptest provide us a way to test our handlers with requests and responses, it even provides ways to test your code with a real HTTP server!

   - The client we're making has a GetSlothfulMessage that sends an HTTP request to the /sloth of its baseURL.
   - Using Go's awesome encoding/json package, the HTTP response body is converted to a SlothfulMessage struct, which is returned if the request and JSON deserialization are successful. We are using json.NewDecoder(res.Body).Decode for reading the response body into our SlothfulMessage struct.

   As you can see, with the httptest package's ResponseRecorder and Server objects, we've got the ability to take the concepts we were already working with for writing tests using the testing package, and then start using functionality to test both receiving and sending HTTP requests. Definitely a must-know package in a Go web developer's toolbelt!

$ go test -v server_test.go server.go
=== RUN   TestHandleSlothfulMessage
--- PASS: TestHandleSlothfulMessage (0.00s)
=== RUN   TestGetSlothfulMessage
--- PASS: TestGetSlothfulMessage (0.00s)
PASS
ok   command-line-arguments     0.007s

*/
