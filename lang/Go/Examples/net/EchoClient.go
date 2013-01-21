////////////////////////////////////////////////////////////////////////////
// Porgram: EchoClient
// Purpose: Client for EchoServer
// Authors: Tong Sun (c) 2013, All rights reserved
////////////////////////////////////////////////////////////////////////////

// Style: gofmt -tabs=false -tabwidth=2

package main

import (
	"net"
	"os"
	"fmt"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage:\n %s host:port\n", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]

	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	handshake(conn)

	os.Exit(0)
}

func handshake(conn net.Conn) {
	// close connection on exit
	defer conn.Close()

	var buf [512]byte
	for i := 1; i <= 5; i++ {
		daytime := time.Now().String()
		fmt.Fprintf(os.Stderr, "=> %s\n", daytime)
		conn.Write([]byte(daytime))

		// read upto 512 bytes
		n, err := conn.Read(buf[0:])
		checkError(err)

		fmt.Println(string(buf[0:n]))

		time.Sleep(time.Second)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
