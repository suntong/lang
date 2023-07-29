// Get the IP to which HTTP request was addressed
// https://stackoverflow.com/questions/62678542/
// https://go.dev/play/p/zTBqVQuZtFL

package main

import (
	"context"
	"log"
	"net"
	"net/http"
)

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	server := &http.Server{
		ConnContext: func(ctx context.Context, conn net.Conn) context.Context {
			log.Println(conn.LocalAddr())
			return ctx
		},
	}

	err = server.Serve(l)
	if err != nil {
		log.Fatal(err)
	}
}

// tmp$ go run ht.go
// 2020/07/01 17:29:46 192.168.1.26:8080
// 2020/07/01 17:29:53 127.0.0.1:8080
// 2020/07/01 17:29:59 [::1]:8080
// 2020/07/01 17:30:04 [::1]:8080
// 2020/07/01 17:30:12 127.0.1.1:8080
