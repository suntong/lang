package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Directory to walk through
	dirPath := os.Getenv("TO_LIST")

	// Slice to hold file names
	var files []string

	// Walk through the directory and collect file names
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error walking through directory:", err)
		return
	}

	b := &strings.Builder{}
	fmt.Fprintf(b, "[")
	for i, f := range files {
		if i != 0 {
			fmt.Fprintf(b, ",")
		}
		fmt.Fprintf(b, "{src: '%s'}", f)
	}
	fmt.Fprintf(b, "]")
	fmt.Println(b.String())
}

// TO_LIST=. go run FilepathWalkRet.go
