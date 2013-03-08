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

////// Debugging //////
import "github.com/davecgh/go-spew/spew"
var _ = spew.Config
//////////////////////

func parse2graph(filename string) *gographviz.Graph {
  f, err := os.Open(filename)
  checkError(err)
  all, err := ioutil.ReadAll(f)
  checkError(err)
  g, err := gographviz.Parse(all) // *ast.Graph
  checkError(err)
  //fmt.Printf("Parsed: %v\n", g)
  //ag := gographviz.NewAnalysedGraph(g) // *gographviz.Graph
  ag := gographviz.NewGraph() // *gographviz.Graph
  cluster.Read(g, ag);

  //fmt.Printf("Analysed: %v\n", ag)
  //fmt.Printf("Written: %v\n", ag.String())
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
  cg := cluster.NewGraph(g)
  //fmt.Printf("Written: %v\n", ag.String())
  cg.NodesStats()

  _ = nn
  // fmt.Printf("%#v\n", cg.Lookup(nn))
  // fmt.Printf("%#v\n", cg.EdgesToParents(nn))
  // fmt.Printf("%#v\n", cg.EdgesToChildren(nn))

  subGraphs := cg.Cluster()
  for i := range subGraphs {
    //spew.Dump(i, subGraphs[i].Graph)
    fmt.Printf("Written: %v: %v\n", i, subGraphs[i].Graph.String())
  }

  os.Exit(0)
}

func checkError(err error) {
  if err != nil {
    fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
    os.Exit(1)
  }
}
