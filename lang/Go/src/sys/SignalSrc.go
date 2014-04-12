// By: Jens Rantil
// Fm: https://gist.github.com/JensRantil/5073646

// Style: gofmt -tabs=false -tabwidth=2

package main

import (
  "fmt"
  "os"
  "os/signal"
  "time"
)

func main() {
  c := make(chan os.Signal)
  signal.Notify(c)
  for sig := range c {
    fmt.Println(sig.String())
    time.Sleep(5 * time.Second)
    return
  }
}
