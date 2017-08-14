////////////////////////////////////////////////////////////////////////////
// Porgram: FileTemp.go
// Purpose: Go TempFile() demo
// Authors: Tong Sun (c) 2017, All rights reserved
// Credits:
// http://golang-examples.tumblr.com/post/81467564564/create-a-temporary-file
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// Use os.TempDir() to get the name of the directory
// and ioutil.TempFile() to create a file.

func main() {
	file, err := ioutil.TempFile(os.TempDir(), "prefix.")
	_ = err
	// It’s caller’s responsibility to remove the file.
	defer os.Remove(file.Name())
	fmt.Println(os.TempDir())
	fmt.Println(file.Name())
}

/*

/tmp
/tmp/prefix.943362349

*/
