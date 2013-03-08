////////////////////////////////////////////////////////////////////////////
// Porgram: tsgraph
// Purpose: provides proper accessing methods and clustering capability 
// Authors: Tong Sun (c) 2013, All rights reserved
////////////////////////////////////////////////////////////////////////////

// Style: gofmt -tabs=false -tabwidth=2

package cluster

import (
  "code.google.com/p/gographviz"
)

////// Debugging //////
import "github.com/davecgh/go-spew/spew"
var _ = spew.Config
//////////////////////

// Cluster break down the given Graph into clusters and returns all the sub-graphs.
func (this *Graph) Cluster() []*gographviz.Graph {
  /*

   Algorithm synopsis:

     Loop over the Starters, for each unvisited Starter,
     define an empty sub-graph and, put it into the toVisit set

     Loop over the toVisit node set, for each node in it, 
       skip if already visited
       add the node to the sub-graph
       remove the nodes into the hasVisited node set
       put all its incoming and outgoing edge into the the toWalk set while
       stop at the hub nodes (edges from the hub nodes are not put in the toWalk set)
       then iterate through the toWalk edge set 
         skip if already walked
         add the edge to the sub-graph
         put its connected nodes into the toVisit node set
         remove the edge from the toWalk edge set into the hasWalked edge set

  */
  
  // sub-graph index
  sgNdx := -1
  sgRet := make([]*Graph,0)

  toVisit := make(set); hasVisited := make(set)
  toWalk := make(set); hasWalked := make(set)

  for starter := range *this.Starters() {
    // define an empty sub-graph and, put it into the toVisit set
    sgRet = append(sgRet, NewGraph(gographviz.NewGraph())); sgNdx++; 
    sgRet[sgNdx].Attrs = this.Attrs
    toVisit.Add(&Node{starter})
    for len(toVisit) > 0 { for nodep := range toVisit {
      if hasVisited.Has(nodep) { continue }
      spew.Dump("toVisit", nodep)
      // add the node to the sub-graph
      sgRet[sgNdx].AddNode(nodep.(*Node))
      // remove the nodes into the hasVisited node set
      toVisit.Del(nodep); hasVisited.Add(nodep)
      // stop at the hub nodes
      if this.IsHub(nodep.(*Node)) { continue }
      // put all its incoming and outgoing edge into the the toWalk set
      noden := nodep.(*Node).String()
      for _, ep := range this.EdgesToParents(noden) {
        toWalk.Add(&Edge{ep})
      }
      for _, ep := range this.EdgesToChildren(noden) {
        toWalk.Add(&Edge{ep})
      }
      for edgep := range toWalk  {
        if hasWalked.Has(edgep) { continue }
        spew.Dump("toWalk", edgep)
        sgRet[sgNdx].Edges.Add(edgep.(*Edge).Edge)
        // put its connected nodes into the toVisit node set
        toVisit.Add(&Node{this.Lookup(edgep.(*Edge).Edge.Src)})
        toVisit.Add(&Node{this.Lookup(edgep.(*Edge).Edge.Dst)})
        // remove the edge from the toWalk edge set into the hasWalked edge set
        toWalk.Del(edgep); hasWalked.Add(edgep)
      }
      //spew.Dump(toVisit)
      //spew.Dump(sgRet)
    }}
    spew.Dump(sgNdx)
  }
  return nil
}
