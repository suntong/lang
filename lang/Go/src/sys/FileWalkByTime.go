////////////////////////////////////////////////////////////////////////////
// Porgram: FileWalkByTime.go
// Purpose: Go file walk by time demo
// Authors: Tong Sun (c) 2017, All rights reserved
// Credits: berserkk
//          https://stackoverflow.com/a/44385800/2125837
////////////////////////////////////////////////////////////////////////////

/*

Q: How to Iterate through directory, ordered based on the file time
https://stackoverflow.com/questions/44380054/

Go provides a directory iteration functionality out of the box, with
filepath.Walk in the path/filepath package.

However, filepath.Walk walks the file tree in lexical order. How can I walks
the file tree in the order of last-modified date?

A: you should implement it by yourself, because filepath.Walk doesn't allow
you to set order.

Here's how you get files in the order of last-modified date (note, that I'm
ignoring errors):

*/

package main

import (
	"fmt"
	"os"
	"sort"
)

type ByModTime []os.FileInfo

func (fis ByModTime) Len() int {
	return len(fis)
}

func (fis ByModTime) Swap(i, j int) {
	fis[i], fis[j] = fis[j], fis[i]
}

func (fis ByModTime) Less(i, j int) bool {
	return fis[i].ModTime().Before(fis[j].ModTime())
}

func main() {
	f, _ := os.Open("/etc")
	fis, _ := f.Readdir(-1)
	f.Close()
	sort.Sort(ByModTime(fis))

	for _, fi := range fis {
		fmt.Println(fi.Name())
	}
}
