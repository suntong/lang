package main

import "fmt"
import "os"
import "log"

func main() {

	log.Printf("Starting...\n")
	for i := 1; i < 1000000000000; i++ {
		if i%100000 == 0 {
			fmt.Fprintf(os.Stderr, ".")
		}
		if i%5000000 == 0 {
			log.Printf("--")
		}
	}
	fmt.Fprintf(os.Stderr, "\n")
}
