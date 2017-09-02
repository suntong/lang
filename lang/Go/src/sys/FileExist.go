////////////////////////////////////////////////////////////////////////////
// Porgram: FileExist.go
// Purpose: Go check file exists demo
// Authors: Tong Sun (c) 2017, All rights reserved
// Credits: As listed below
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"os"
)

func main() {
	var fileC string
	fileC = "/etc/fstab"
	fmt.Printf("File '%s' exist: %v\n", fileC, IsExist(fileC))
	fileC = "/tmp/not-exist"
	fmt.Printf("File '%s' exist: %v\n", fileC, IsExist(fileC))
	fileC = "/usr/local/share"
	fmt.Printf("Dir '%s' exist: %v\n", fileC, IsExist(fileC))
}

/*

File '/etc/fstab' exist: true
File '/tmp/not-exist' exist: false
Dir '/usr/local/share' exist: true

*/

// IsExist checks if the given file exist
func IsExist(fileName string) bool {
	//fmt.Printf("] Checking %s\n", fileName)
	_, err := os.Stat(fileName)
	return err == nil || os.IsExist(err)
	// CAUTION! os.IsExist(err) != !os.IsNotExist(err)
	// https://gist.github.com/mastef/05f46d3ab2f5ed6a6787#file-isexist_vs_isnotexist-go-L35-L56
}
