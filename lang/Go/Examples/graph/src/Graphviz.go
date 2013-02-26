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
  //"code.google.com/p/gographviz/ast"
  "code.google.com/p/gographviz/parser"

  "github.com/davecgh/go-spew/spew"
)

var _ = spew.Config

func showNodes(this *gographviz.Graph) {
  fmt.Println("Nodes Lookups")
  for k, v := range this.Nodes.Lookup {
    fmt.Printf("%s -> %s\n", k, v)
  }

  fmt.Println("Nodes")
  for i := range this.Nodes.Nodes {
    fmt.Printf("%v ->: %s\n",i, this.Nodes.Nodes[i])
  }
}

func showEdges(this *gographviz.Graph) {
  fmt.Println("Edges:")
  //edges := this.Edges.Sorted()
  edges := this.Edges.Edges
  for i := range edges {
    fmt.Printf("%v ->: %s\n",i, edges[i])
  }
}

func parser2graph(filename string) *gographviz.Graph {
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

func parse2graph(filename string) *gographviz.Graph {
  f, err := os.Open(filename)
  checkError(err)
  all, err := ioutil.ReadAll(f)
  checkError(err)
  g, err := gographviz.Parse(all) // *ast.Graph
  checkError(err)
  //fmt.Printf("Parsed: %v\n", g)
  ag := gographviz.NewAnalysedGraph(g)
  //fmt.Printf("Analysed: %v\n", ag)
  fmt.Printf("Written: %v\n", ag.String())
  return ag
}

func read2graph(filename string) *gographviz.Graph {
  f, err := os.Open(filename)
  checkError(err)
  all, err := ioutil.ReadAll(f)
  checkError(err)
  g, err := gographviz.Read(all) // *gographviz.Graph
  checkError(err)
  //fmt.Printf("Written: %v\n", g.String())
  return g
}

func main() {
  if len(os.Args) != 2 {
    fmt.Fprintf(os.Stderr, "Usage:\n %s graphviz_file\n", os.Args[0])
    os.Exit(1)
  }
  gf := os.Args[1]

  //_ = parser2graph(gf)
  // g := parse2graph(gf)
  g := read2graph(gf); _ = g
  //showNodes(g)
  //spew.Dump("Nodes: ", g.Nodes)
  //spew.Dump("Edges ", g.Edges)
  //spew.Dump("Edges ", g.Edges.SrcToDsts)
  showEdges(g)
  // spew.Dump("Relations P-C: ", g.Relations.ParentToChildren)
  // spew.Dump("Relations C-T: ", g.Relations.ChildToParents)

  os.Exit(0)
}

func checkError(err error) {
  if err != nil {
    fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
    os.Exit(1)
  }
}
