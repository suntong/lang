package main

import "fmt"
import "os"
import "time"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Fprintf(os.Stderr, ".")
		time.Sleep(time.Second)
	}
	fmt.Fprintf(os.Stderr, "\n")
}
