////////////////////////////////////////////////////////////////////////////
// Porgram: tsgraph
// Purpose: provides proper accessing methods and clustering capability 
// Authors: Tong Sun (c) 2013, All rights reserved
////////////////////////////////////////////////////////////////////////////

// Style: gofmt -tabs=false -tabwidth=2

package cluster

import (
  "fmt"

  "code.google.com/p/gographviz"
)

////// Debugging //////
import "github.com/davecgh/go-spew/spew"
var _ = spew.Config
//////////////////////

// Cluster break down the given Graph into clusters and returns all the sub-graphs.
func (this *Graph) Cluster() []*Graph {
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

  toVisit := make(nodeSet); hasVisited := make(nodeSet)
  toWalk := make(edgeSet); hasWalked := make(edgeSet)

  for starter := range *this.Starters() {
    // define an empty sub-graph and, put it into the toVisit set
    sgRet = append(sgRet, NewGraph(gographviz.NewGraph())); sgNdx++; 
    sgRet[sgNdx].Attrs = this.Attrs
    sgRet[sgNdx].SetDir(this.Directed)
    graphName := fmt.Sprintf("%s_%03d\n", this.Name, sgNdx);
    sgRet[sgNdx].SetName(graphName)
    toVisit.Add(starter)
    hubVisited := make(nodeSet)
    for len(toVisit) > 0 { for nodep := range toVisit {
      toVisit.Del(nodep);       //print("O ")
      if this.IsHub(nodep) && hasVisited.Has(nodep) && !hubVisited.Has(nodep) { 
        // add the already-visited but not-in-this-graph hub node to the sub-graph
        sgRet[sgNdx].AddNode(nodep)
        hubVisited.Add(nodep)
        continue 
      }
      if hasVisited.Has(nodep) { continue }
      //spew.Dump("toVisit", nodep)
      // add the node to the sub-graph
      sgRet[sgNdx].AddNode(nodep)
      // remove the nodes into the hasVisited node set
      hasVisited.Add(nodep)
      // stop at the hub nodes
      if this.IsHub(nodep) { continue }
      // put all its incoming and outgoing edge into the the toWalk set
      noden := nodep.Name
      for _, ep := range this.EdgesToParents(noden) {
        toWalk.Add(ep)
      }
      for _, ep := range this.EdgesToChildren(noden) {
        toWalk.Add(ep)
      }
      for edgep := range toWalk  {
        toWalk.Del(edgep);         //print("- ")
        if hasWalked.Has(edgep) { continue }
        //spew.Dump("toWalk", edgep)
        sgRet[sgNdx].Edges.Add(edgep)
        // put its connected nodes into the toVisit node set
        toVisit.Add(this.Lookup(edgep.Src))
        toVisit.Add(this.Lookup(edgep.Dst))
        // remove the edge into the hasWalked edge set
        hasWalked.Add(edgep)
      }
    }}
    //spew.Dump(sgNdx)
  }
  return sgRet
}
