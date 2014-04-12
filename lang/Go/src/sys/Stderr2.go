package main

import "fmt"
import "os"
import "log"
import "time"

func main() {
	file, err := os.Create("temp")
	if err != nil {
		panic(err)
	}
	// close file on exit and check for its returned error
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	log.Printf("Starting...\n")
	for i := 0; i < 10; i++ {
		fmt.Fprintf(file, "%02d\n", i)
		fmt.Fprintf(os.Stderr, ".")
		if i%100000 == 0 {
			fmt.Fprintf(os.Stderr, ".")
		}
	}
		time.Sleep(time.Second)
	}
	fmt.Fprintf(os.Stderr, "\n")
}
