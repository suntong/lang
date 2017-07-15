////////////////////////////////////////////////////////////////////////////
// Porgram: FilepathDir
// Purpose: filepath dir (Base/Split/Glob) demo
// Authors: Tong Sun (c) 2013, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func isDir(path string) bool {
	info, _ := os.Stat(path)
	return info.IsDir()
}

func Basename(s string) string {
	n := strings.LastIndexByte(s, '.')
	if n > 0 {
		return s[:n]
	}
	return s
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:\n  ", os.Args[0], "root_path")
		os.Exit(1)
	}

	root := os.Args[1]

	files, _ := filepath.Glob(root + "/*")

	for _, f := range files {
		if isDir(f) {
			f += "/"
			fmt.Println("D:", f, filepath.Dir(f), filepath.Base(f),
				Basename(filepath.Base(f)))
		} else {
			p, n := filepath.Split(f)
			fmt.Printf("F: %v='%v'+'%v'\n", f, p, n)
			fmt.Printf("  name='%v', ext='%v'\n", Basename(filepath.Base(f)), filepath.Ext(f))
		}
	}
}

/*

$ go run FilepathDir.go /etc/ssl/
D: /etc/ssl/certs/ /etc/ssl/certs certs certs
F: /etc/ssl/openssl.cnf='/etc/ssl/'+'openssl.cnf'
  name='openssl', ext='.cnf'
D: /etc/ssl/private/ /etc/ssl/private private private

*/
