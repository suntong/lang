////////////////////////////////////////////////////////////////////////////
// Porgram: tsgraph
// Purpose: a graph class with proper accessing methods defined
// Authors: Tong Sun (c) 2013, All rights reserved
////////////////////////////////////////////////////////////////////////////

// Style: gofmt -tabs=false -tabwidth=2

/*

This is a proper wrapper for gographviz, code.google.com/p/gographviz.

The reason for this wrapper is that I believe the biggest no-no in 
programming, is to dip into other's class and directly access or 
manipulate their internal data structure. However, the author of 
gographviz has a different view on this and strongly refused to 
wrap up its internal representations.

This class will properly hide gographviz's internal data structure by
providing access methods instead. Ref:
http://en.wikipedia.org/wiki/Encapsulation_(object-oriented_programming)

*/

package cluster

import (
  "code.google.com/p/gographviz"
)

type Graph struct {
  gographviz.Graph
}

type Edge struct {
  gographviz.Edge
}

type Node struct {
  gographviz.Node
}

type Nodes struct {
  gographviz.Nodes
}

func (this *Graph) Lookup(nodeName string) *gographviz.Node {
  return this.Nodes.Lookup[nodeName];
}

func (this *Graph) EdgesToParents(nodeName string) map[string]*gographviz.Edge {
  return this.Edges.DstToSrcs[nodeName];
}

func (this *Graph) EdgesToChildren(nodeName string) map[string]*gographviz.Edge {
  return this.Edges.SrcToDsts[nodeName];
}
