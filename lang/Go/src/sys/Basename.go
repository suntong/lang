package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println(filepath.Base("/a/b.c"))
	fmt.Println(filepath.Base(`C:\Program Files\Microsoft Visual Studio 9\Common7\IDE\devenv.exe`))
	fmt.Println(filepath.Base("."))
	abs, _ := filepath.Abs(".")
	fmt.Println(filepath.Base(abs))
}

/*

$ go run Basename.go
b.c
C:\Program Files\Microsoft Visual Studio 9\Common7\IDE\devenv.exe
.
sys

*/
