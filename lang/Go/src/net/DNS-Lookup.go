package main

import (
	"log"
	"net"
	// "net/http"
)

const (
	domain = "amazonaws.com"
)

func lookup() {
	log.Println(net.LookupHost(domain))
}

func main() {
	lookup()
}
