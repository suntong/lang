package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("FileSize.go")
	checkError(err)
	defer file.Close()

	// get the file size
	stat, err := file.Stat()
	checkError(err)
	fmt.Println("File size is ", stat.Size())
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
