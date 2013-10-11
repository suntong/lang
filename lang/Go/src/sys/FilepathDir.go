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
)

func isDir(path string) bool {
  info, _ := os.Stat(path)
  return info.IsDir()
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
      fmt.Println("D:", f, filepath.Dir(f), filepath.Base(f))
    } else {
      p, n := filepath.Split(f)
      fmt.Printf("F: %v='%v'+'%v'\n", f, p, n)
      fmt.Printf("  ext='%v'\n", filepath.Ext(f))
    }
  }
}
