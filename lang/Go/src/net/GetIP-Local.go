// How to get the local IP address in Go
// https://stackoverflow.com/questions/23558425/

package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	fmt.Println(GetOutboundIP().String())
}

// Get preferred outbound ip of this machine
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
