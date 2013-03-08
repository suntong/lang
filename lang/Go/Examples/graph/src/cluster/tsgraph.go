////////////////////////////////////////////////////////////////////////////
// Porgram: tsgraph
// Purpose: provides proper accessing methods and clustering capability 
// Authors: Tong Sun (c) 2013, All rights reserved
////////////////////////////////////////////////////////////////////////////

// Style: gofmt -tabs=false -tabwidth=2

/*
Package cluster provides a graph class with clustering capability for gographviz.Graph class
as well as proper accessing methods defined.
*/
package cluster

import (
  "code.google.com/p/gographviz"
)

////// Debugging //////
import "github.com/davecgh/go-spew/spew"
var _ = spew.Config

func debugLn(x string)  { println("DEBUG: " + x) }
func debugNop(x string) {}

var debug func(x string) = debugLn
//////////////////////

/*

Type Graph gives proper wrapping for gographviz.Graph, code.google.com/p/gographviz.

The reason for this wrapper is that I believe the biggest no-no in 
programming, is to dip into other's class and directly access or 
manipulate their internal data structure. However, the author of 
gographviz has a different view on this and strongly refused to 
wrap up its internal representations.

This class will properly hide gographviz's internal data structure by providing 
access methods instead, a practice known as encapsulation in software engineering
(http://en.wikipedia.org/wiki/Encapsulation_(object-oriented_programming).

*/
type Graph struct {
  *gographviz.Graph
  nodesStats map[*gographviz.Node]nodeStat
  Starter nodeSet
  Hub nodeSet
}

type nodeStat struct {
  CntI int
  CntO int
}

// Type NodeSet represents a generic SET data structure of type *gographviz.Node
type NodeSet interface {
  Add(*gographviz.Node)
  Has(*gographviz.Node) bool
  Del(*gographviz.Node)
}
type nodeSet map[*gographviz.Node]struct{}

func (s nodeSet) Add(a *gographviz.Node)   { s[a] = struct{}{} }
func (s nodeSet) Del(a *gographviz.Node)   { delete(s, a) }
func (s nodeSet) Has(a *gographviz.Node) bool { _, ok := s[a]; return ok }

// Type EdgeSet represents a generic SET data structure of type *gographviz.Edge
type EdgeSet interface {
  Add(*gographviz.Edge)
  Has(*gographviz.Edge) bool
  Del(*gographviz.Edge)
}
type edgeSet map[*gographviz.Edge]struct{}

func (s edgeSet) Add(a *gographviz.Edge)   { s[a] = struct{}{} }
func (s edgeSet) Del(a *gographviz.Edge)   { delete(s, a) }
func (s edgeSet) Has(a *gographviz.Edge) bool { _, ok := s[a]; return ok }


func NewGraph(g *gographviz.Graph) *Graph {
	return &Graph{g, 
    make(map[*gographviz.Node]nodeStat),
    make(nodeSet),
    make(nodeSet),
  }
}

// Add an existing node to a graph
func (this *Graph) AddNode(np *gographviz.Node) {
  this.Graph.AddNode(this.Name, np.Name, np.Attrs);
}

// Lookup lookups the given nodeName within the Graph.
func (this *Graph) Lookup(nodeName string) *gographviz.Node {
  return this.Nodes.Lookup[nodeName];
}

// EdgesToParents returns all the edges linked to its parents of the given nodeName.
func (this *Graph) EdgesToParents(nodeName string) map[string]*gographviz.Edge {
  return this.Edges.DstToSrcs[nodeName];
}

// EdgesToChildren returns all the edges linked to its children of the given nodeName.
func (this *Graph) EdgesToChildren(nodeName string) map[string]*gographviz.Edge {
  return this.Edges.SrcToDsts[nodeName];
}

// NodesAll returns all the nodes within the given Graph.
func (this *Graph) NodesAll() []*gographviz.Node {
  return this.Nodes.Nodes;
}

// NodesStatGet performs statistics of all nodes within the given Graph.
func (this *Graph) NodesStats() {
  nodes := this.NodesAll();
  for i := range nodes {
    node := nodes[i]; nodeName := node.Name
    this.nodesStats[node] = nodeStat{ 
      len(this.EdgesToParents(nodeName)), 
      len(this.EdgesToChildren(nodeName)) }
    //spew.Dump(i, node, this.nodesStats[node]) 
  }
  for node, nodeStat := range this.nodesStats {
    //spew.Dump(node, nodeStat) 
    if nodeStat.CntI == 0 { this.Starter.Add(node) } 
    if nodeStat.CntI >= 3 { this.Hub.Add(node)  } 
  }
  //spew.Dump(this.Starter, this.Hub) 
}

// Starters returns all the starter nodes within the given Graph.
// Starters are those nodes that has only outgoing edges but no incoming ones.
func (this *Graph) Starters() *nodeSet {
  return &this.Starter;
}

// Hubs returns all the hub nodes within the given Graph.
// Hubs are those nodes with the number of incoming edges over the given threshold.
func (this *Graph) Hubs() *nodeSet {
  return &this.Hub;
}

// Hubs returns all the hub nodes within the given Graph.
// Hubs are those nodes with the number of incoming edges over the given threshold.
func (this *Graph) IsHub(n *gographviz.Node) bool {
  _, ok := this.Hub[n]; return ok
}

