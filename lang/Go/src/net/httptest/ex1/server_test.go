package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUpperCaseHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/upper?word=abc", nil)
	w := httptest.NewRecorder()
	upperCaseHandler(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if string(data) != "ABC" {
		t.Errorf("expected ABC got %v", string(data))
	}
}

/*

 https://golang.cafe/blog/golang-httptest-example.html

 The idea behind the httptest package is that it’ pretty easy to mock an HTTP server or to mock a response writer to feed into our server handlers. This makes it extremely easy to test http servers and clients.

# No need to: go run server.go &

$ go test -v server_test.go server.go
=== RUN   TestUpperCaseHandler
--- PASS: TestUpperCaseHandler (0.00s)
PASS
ok   command-line-arguments     0.003s

*/
