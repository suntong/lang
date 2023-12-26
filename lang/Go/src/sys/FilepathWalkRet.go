package main

import (
	"fmt"
	"os"
	"path/filepath"
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

	fmt.Println(files)
}

// TO_LIST=. go run FilepathWalkRet.go
