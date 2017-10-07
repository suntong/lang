package main

import (
	"log"

	"github.com/miekg/dns"
)

func main() {

	target := "microsoft.com"
	server := "8.8.8.8"

	c := dns.Client{}
	m := dns.Msg{}
	m.SetQuestion(target+".", dns.TypeA)
	r, t, err := c.Exchange(&m, server+":53")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Took %v", t)
	if len(r.Answer) == 0 {
		log.Fatal("No results")
	}
	for _, ans := range r.Answer {
		Arecord := ans.(*dns.A)
		log.Printf("%s", Arecord.A)
	}
}
