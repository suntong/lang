////////////////////////////////////////////////////////////////////////////
// Porgram: embedFS.go
// Purpose: Go:embed FS demo
// Authors: Tong Sun (c) 2022, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"embed"
	"fmt"
)

//go:embed hello.txt
// see https://pkg.go.dev/embed
var f embed.FS

//go:embed html/index.html
var content embed.FS

func main() {
	data, _ := f.ReadFile("hello.txt")
	fmt.Printf("'%s'\n", string(data))

	data, _ = content.ReadFile("html/index.html")
	fmt.Printf("'%s'\n", string(data))
}

/*

$ go run embedFS.go
'hello world!
'
'html
'

*/
