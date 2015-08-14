////////////////////////////////////////////////////////////////////////////
// Porgram: FilepathWalk.go
// Purpose: filepath Walk demo
// Authors: Tong Sun (c) 2013, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func visit(path string, f os.FileInfo, err error) error {
	if f.IsDir() {
		fmt.Printf("  Directory: %s\n", path)
	} else {
		fmt.Printf("  File: %s with %d bytes\n", path, f.Size())
	}
	return nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:\n  ", os.Args[0], "root_path")
		os.Exit(1)
	}

	root := os.Args[1]
	err := filepath.Walk(root, visit)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
