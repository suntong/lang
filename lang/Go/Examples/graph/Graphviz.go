////////////////////////////////////////////////////////////////////////////
// Porgram: Graphviz
// Purpose: gographviz demo
// Authors: Tong Sun (c) 2013, All rights reserved
////////////////////////////////////////////////////////////////////////////

// Style: gofmt -tabs=false -tabwidth=2

package main

import (
  "fmt"
  "io/ioutil"
  "os"

  "code.google.com/p/gographviz"
  "code.google.com/p/gographviz/parser"
)

// func (this *Nodes) String() string {
//   s := "Nodes:"
//   for i := range this.Nodes {
//     s += fmt.Sprintf("Node{%v}", this.Nodes[i])
//   }
//   return s + "\n"
// }

// func (this *Edges) String() string {
//   s := "Edges:"
//   for i := range this.Edges {
//     s += fmt.Sprintf("Edge{%v}", this.Edges[i])
//   }
//   return s + "\n"
// }

func openfile(filename string) *gographviz.Graph {
  f, err := os.Open(filename)
  checkError(err)
  all, err := ioutil.ReadAll(f)
  checkError(err)
  input := string(all)
  g, err := parser.ParseString(input)
  checkError(err)
  fmt.Printf("Parsed: %v\n", g)
  ag := gographviz.NewAnalysedGraph(g)
  fmt.Printf("Analysed: %v\n", ag)
  agstr := ag.String()
  fmt.Printf("Written: %v\n", agstr)
  return ag
}

func main() {
  if len(os.Args) != 2 {
    fmt.Fprintf(os.Stderr, "Usage:\n %s graphviz_file\n", os.Args[0])
    os.Exit(1)
  }
  gf := os.Args[1]

  _ = openfile(gf)
  os.Exit(0)
}

func checkError(err error) {
  if err != nil {
    fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
    os.Exit(1)
  }
}
