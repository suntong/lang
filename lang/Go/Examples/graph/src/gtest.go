////////////////////////////////////////////////////////////////////////////
// Porgram: gtest
// Purpose: tester program for cluster/cluster.go
// Authors: Tong Sun (c) 2013, All rights reserved
////////////////////////////////////////////////////////////////////////////

// Style: gofmt -tabs=false -tabwidth=2

package main

import (
  "fmt"
  "io/ioutil"
  "os"

  "code.google.com/p/gographviz"

  "cluster"
)

func parse2graph(filename string) *gographviz.Graph {
  f, err := os.Open(filename)
  checkError(err)
  all, err := ioutil.ReadAll(f)
  checkError(err)
  g, err := gographviz.Parse(all) // *ast.Graph
  checkError(err)
  //fmt.Printf("Parsed: %v\n", g)
  //ag := gographviz.NewAnalysedGraph(g) // *Graph
  ag := gographviz.NewGraph() // *Graph
  cluster.Read(g, ag);

  //fmt.Printf("Analysed: %v\n", ag)
  fmt.Printf("Written: %v\n", ag.String())
  return ag
}

func main() {
  if len(os.Args) < 3 {
    fmt.Fprintf(os.Stderr, "Usage:\n %s graphviz_file nodeName\n", os.Args[0])
    os.Exit(1)
  }
  gf := os.Args[1]
  nn := os.Args[2]

  //_ = parser2graph(gf)
  g := parse2graph(gf)
  cg := cluster.Graph{*g}
  
  fmt.Printf("%#v\n", cg.Lookup(nn))
  fmt.Printf("%#v\n", cg.EdgesToParents(nn))
  fmt.Printf("%#v\n", cg.EdgesToChildren(nn))

  os.Exit(0)
}

func checkError(err error) {
  if err != nil {
    fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
    os.Exit(1)
  }
}
