package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	// Create an IPv4 connection to the server, following `curl -v4 ifconfig.me`
	conn, err := net.Dial("tcp4", "ifconfig.me:80")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	// Form the HTTP GET request
	//request := "GET / HTTP/1.1\r\nHost: ifconfig.me\r\nAccept: */*\r\n\r\n"

	// Send the request ("Host:" is needed besides GET)
	_, err = fmt.Fprintf(conn, "GET / HTTP/1.1\r\nHost: ifconfig.me\r\n\r\n")
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}

	// fmt.Println("Read the response")
	// Read the response
	// X: resp, err := ioutil.ReadAll(conn)
	buffer := make([]byte, 4096)
	// respBytes := []byte{}
	// for {
	n, err := conn.Read(buffer)
	// 	fmt.Println(n, buffer[:n])
	// 	if err != nil || n == 0 {
	// 		break
	// 	}
	// 	respBytes = append(respBytes, buffer[:n]...)
	// }
	response := string(buffer[:n])

	// Splitting the headers and body
	headersAndBody := strings.SplitN(response, "\r\n\r\n", 2)
	headers := headersAndBody[0]
	body := headersAndBody[1]
	fmt.Println("Response:\n", response, headers)
	fmt.Println("Body:", body)

	// Handle the body as needed
}
