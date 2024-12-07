package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

/*

  With recent improvements in Go 1.22, the net/http.ServeMux now supports
  enhanced routing capabilities, including path patterns with
  wildcards. While it may not match the raw speed of specialized routers
  like fasthttp, its performance has improved, and it integrates seamlessly
  with the Go standard library.

*/

func main() {
	// Root router
	mux := http.NewServeMux()

	// Middleware
	logger := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next.ServeHTTP(w, r)
			log.Printf("Request: %s %s | Duration: %v", r.Method, r.URL.Path, time.Since(start))
		})
	}

	// Handlers
	mux.HandleFunc("GET /tasks", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Fetching all tasks")
	})
	mux.HandleFunc("GET /tasks/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Fprintf(w, "Fetching task with ID: %s", id)
	})
	mux.HandleFunc("POST /tasks", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Creating a new task")
	})

	// == more comprehensive demo

	// Simple text response
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintln(w, "Welcome to the ServeMux demo!")
		w.Header().Add("Hello", "world1")
		w.Header().Add("Hello", "world2")
		http.SetCookie(w, &http.Cookie{Name: "yummy", Value: "cookie"})
		http.SetCookie(w, &http.Cookie{Name: "yummy", Value: "cake"})
		http.SetCookie(w, &http.Cookie{Name: "fruit", Value: "banana", Expires: time.Date(2000, time.January, 0, 0, 0, 0, 0, time.UTC)})
		for _, c := range r.Cookies() {
			http.SetCookie(w, c)
		}
		w.WriteHeader(200)

		encoder := json.NewEncoder(w)
		encoder.Encode(struct{ RequestQueryParams, Method any }{r.URL.Query(), r.Method})
	})

	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set("Content-Length", "12") // len("Hello World!"))
		w.WriteHeader(200)
		w.Write([]byte("Hello World!"))

	})

	// Handling query parameters
	mux.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		msg := r.URL.Query().Get("msg")
		if msg == "" {
			msg = "Default message"
		}
		fmt.Fprintf(w, "Echo: %s\n", msg)
	})

	// JSON response
	mux.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := map[string]string{"message": "Hello, JSON!"}
		json.NewEncoder(w).Encode(response)
	})

	// Extracting path variables (simulating dynamic routing)
	mux.HandleFunc("/square/", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Path[len("/square/"):]
		num, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "Invalid number", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "Square of %d is %d\n", num, num*num)
	})

	// Handling POST requests
	mux.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}
		var data map[string]string
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
			return
		}
		response := map[string]string{"received": data["message"]}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	// Start the server on port 8080.
	log.Println("Server started at :8080")
	// Using middleware
	log.Fatal(http.ListenAndServe(":8080", logger(mux)))
}

/*

$ curl localhost:8080/hello
2024/12/07 09:39:54 Request: GET /hello | Duration: 33.56µs
Hello World!

$ curl localhost:8080/tasks
2024/12/07 09:34:24 Request: GET /tasks | Duration: 53.267µs
Fetching all tasks

$ curl localhost:8080/tasks/123
2024/12/07 09:34:45 Request: GET /tasks/123 | Duration: 6.697µs
Fetching task with ID: 123

$ curl -d '' localhost:8080/tasks
2024/12/07 09:35:06 Request: POST /tasks | Duration: 31.096µs
Creating a new task

$ curl localhost:8080/json -i
2024/12/07 09:40:43 Request: GET /json | Duration: 32.383µs
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sat, 07 Dec 2024 01:40:43 GMT
Content-Length: 27

{"message":"Hello, JSON!"}

$ curl localhost:8080/square/9
2024/12/07 09:41:25 Request: GET /square/9 | Duration: 82.921µs
Square of 9 is 81

$ curl -d '{"message":"Hello, JSON!"}' localhost:8080/post
2024/12/07 09:42:14 Request: POST /post | Duration: 46.324µs
{"received":"Hello, JSON!"}


$ curl localhost:8080/
2024/12/07 09:39:42 Request: GET / | Duration: 65.494µs
{"RequestQueryParams":{},"Method":"GET"}

$ curl -i localhost:8080/'?foo=bar&hello=world&hello=universe'
2024/12/07 09:50:19 Request: GET / | Duration: 44.634µs
HTTP/1.1 200 OK
Hello: world1
Hello: world2
Set-Cookie: yummy=cookie
Set-Cookie: yummy=cake
Set-Cookie: fruit=banana; Expires=Fri, 31 Dec 1999 00:00:00 GMT
Date: Sat, 07 Dec 2024 01:50:19 GMT
Content-Length: 83
Content-Type: text/plain; charset=utf-8

{"RequestQueryParams":{"foo":["bar"],"hello":["world","universe"]},"Method":"GET"}

*/
