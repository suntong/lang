// Create Hardlink with golang
// https://stackoverflow.com/questions/16800044/create-hardlink-with-golang

package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	err := ioutil.WriteFile("original.txt", []byte("hello world"), 0600)
	if err != nil {
		log.Fatalln(err)
	}

	err = os.Link("original.txt", "link.txt")
	if err != nil {
		log.Fatalln(err)
	}
}

/*

go run File-Hardlink.go
ls -l original.txt link.txt
stat original.txt link.txt
rm original.txt link.txt

*/
