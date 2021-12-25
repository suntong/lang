// https://tutorialedge.net/golang/temporary-files-directories-go-111/

package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	tempDir, err := ioutil.TempDir("", "cars-")
	if err != nil {
		fmt.Println(err)
	}
	defer os.RemoveAll(tempDir)

	file, err := ioutil.TempFile(tempDir, "car-*.png")
	if err != nil {
		fmt.Println(err)
	}
	defer os.Remove(file.Name())

	// This will print out the full name and path of our image
	fmt.Println(file.Name())

	if _, err := file.Write([]byte("hello world\n")); err != nil {
		fmt.Println(err)
	}

	data, err := ioutil.ReadFile(file.Name())
	// if our program was unable to read the file
	// print out the reason why it can't
	if err != nil {
		fmt.Println(err)
	}

	// if it was successful in reading the file then
	// print out the contents as a string
	fmt.Print(string(data))

}

/*

$ go run FileTemp.go
/tmp/cars-3532421589/car-1014067776.png
hello world

*/
