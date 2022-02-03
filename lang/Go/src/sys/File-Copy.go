// Copy a file in Go
// https://stackoverflow.com/questions/30376921/how-do-you-copy-a-file-in-go/

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	copyFile("test.txt", "test_copy.txt")
}

func copyFile(src, dst string) {
	srcFile, err := os.Open(src)
	check(err)
	defer srcFile.Close()

	destFile, err := os.Create(dst) // creates if file doesn't exist
	check(err)
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile) // check first var for number of bytes copied
	check(err)

	err = destFile.Sync()
	check(err)
}

func check(err error) {
	if err != nil {
		fmt.Println("Error : %s", err.Error())
		os.Exit(1)
	}
}

/*

echo "abc" > test.txt
ls -l
cat test_copy.txt
rm test.txt test_copy.txt

*/
