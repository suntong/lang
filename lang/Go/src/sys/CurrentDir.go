package main

import (
  "fmt"
  "log"
  "os"
  "path/filepath"
)

func main() {
  fmt.Println(os.Args[0])

  dir := filepath.Dir(os.Args[0])
  fmt.Print("Dir=")
  fmt.Println(dir)

  dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
  if err != nil {
    log.Fatal(err)
  }
  fmt.Print("Abs=")
  fmt.Println(dir)

  dir, err = os.Getwd()
  fmt.Print("Wd=")
  fmt.Println(dir)
}
