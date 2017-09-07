// https://groups.google.com/d/msg/golang-nuts/12RJ7p_BV-g/EfOaSAiZAQAJ

package main

import (
	"log"
	"plugin"
	"time"
)

func main() {
	log.Println("start load plugin.")
	time.Sleep(time.Second)
	p, err := plugin.Open("say_hello.so")
	if err != nil {
		log.Println(err)
	}

	sym, err := p.Lookup("Hello")
	if err != nil {
		log.Println(err)
	}

	hello, ok := sym.(func())
	if !ok {
		log.Println(err)
	}

	for i := 0; i < 10; i++ {
		hello()
		time.Sleep(time.Second)
	}
}
