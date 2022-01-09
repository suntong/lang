////////////////////////////////////////////////////////////////////////////
// Porgram: tsgread
// Purpose: Read (analyse) ast.Graph as gographviz.Graph, with SubGraphs flattened out
// Authors: Tong Sun (c) 2013, All rights reserved
////////////////////////////////////////////////////////////////////////////

// Style: gofmt -tabs=false -tabwidth=2

/////////////////////////////////////////////////////////////////
// Based on gographviz.analyse.go by Walter Schulze
// http://code.google.com/p/gographviz/source/browse/analyse.go
//Copyright 2013 Vastech SA (PTY) LTD
//Licensed under the Apache License, Version 2.0 (the "License");
/////////////////////////////////////////////////////////////////

package cluster

import (
  "code.google.com/p/gographviz"
  "code.google.com/p/gographviz/ast"
)

////// Debugging //////
import "github.com/davecgh/go-spew/spew"
var _ = spew.Config
//////////////////////

//Creates a Graph structure by analysing an Abstract Syntax Tree representing a parsed graph.
// func NewAnalysedGraph(graph *ast.Graph) *gographviz.Graph {
// 	g := gographviz.NewGraph()
// 	Analyse(graph, g)
// 	return g
// }

//Reads (analyses) an Abstract Syntax Tree representing a parsed graph into a newly created graph structure Interface.
func Read(graph *ast.Graph, g gographviz.Interface) {
  graph.Walk(&graphVisitor{g})
  delete(g.(*gographviz.Graph).Attrs, "rank")
  delete(g.(*gographviz.Graph).Attrs, "size")
  debug("Done.")
}

type nilVisitor struct {
}

func (this *nilVisitor) Visit(v ast.Elem) ast.Visitor {
  return this
}

type graphVisitor struct {
  g gographviz.Interface
}

func (this *graphVisitor) Visit(v ast.Elem) ast.Visitor {
  graph, ok := v.(*ast.Graph)
  if !ok {
    return this
  }
  this.g.SetStrict(graph.Strict)
  this.g.SetDir(graph.Type == ast.DIGRAPH)
  graphName := graph.Id.String()
  this.g.SetName(graphName)
  return newStmtVisitor(this.g, graphName)
}

func newStmtVisitor(g gographviz.Interface, graphName string) *stmtVisitor {
  return &stmtVisitor{g, graphName,
    make(gographviz.Attrs), make(gographviz.Attrs), make(gographviz.Attrs)}
}

type stmtVisitor struct {
  g                 gographviz.Interface
  graphName         string
  currentNodeAttrs  gographviz.Attrs
  currentEdgeAttrs  gographviz.Attrs
  currentGraphAttrs gographviz.Attrs
}

func (this *stmtVisitor) Visit(v ast.Elem) ast.Visitor {
  switch s := v.(type) {
  case ast.NodeStmt:
    //debug("Visiting " + s.String())
    return this.nodeStmt(s)
  case ast.EdgeStmt:
    //debug("Visiting " + s.String())
    return this.edgeStmt(s)
  case ast.NodeAttrs:
    return this.nodeAttrs(s)
  case ast.EdgeAttrs:
    return this.edgeAttrs(s)
  case ast.GraphAttrs:
    return this.graphAttrs(s)
  case *ast.SubGraph:
    //debug("Visiting " + s.String())
    //return this.subGraph(s)
  case *ast.Attr:
    return this.attr(s)
  case ast.AttrList:
    return &nilVisitor{}
  default:
    //fmt.Fprintf(os.Stderr, "unknown stmt %T\n", v)
  }
  return this
}

func ammend(attrs gographviz.Attrs, add gographviz.Attrs) gographviz.Attrs {
  for key, value := range add {
    if _, ok := attrs[key]; !ok {
      attrs[key] = value
    }
  }
  return attrs
}

func overwrite(attrs gographviz.Attrs, overwrite gographviz.Attrs) gographviz.Attrs {
  for key, value := range overwrite {
    attrs[key] = value
  }
  return attrs
}

func (this *stmtVisitor) nodeStmt(stmt ast.NodeStmt) ast.Visitor {
  attrs := gographviz.Attrs(stmt.Attrs.GetMap())
  attrs = ammend(attrs, this.currentNodeAttrs)
  this.g.AddNode(this.graphName, stmt.NodeId.String(), attrs)
  return &nilVisitor{}
}

func (this *stmtVisitor) edgeStmt(stmt ast.EdgeStmt) ast.Visitor {
  attrs := stmt.Attrs.GetMap()
  attrs = ammend(attrs, this.currentEdgeAttrs)
  src := stmt.Source.GetId()
  srcName := src.String()
  if stmt.Source.IsNode() {
    this.g.AddNode(this.graphName, srcName, this.currentNodeAttrs.Copy())
  }
  srcPort := stmt.Source.GetPort()
  for i := range stmt.EdgeRHS {
    //debug("  Processing " + stmt.EdgeRHS[i].String())
    directed := bool(stmt.EdgeRHS[i].Op)
    dst := stmt.EdgeRHS[i].Destination.GetId()
    dstName := dst.String()
    if stmt.EdgeRHS[i].Destination.IsNode() {
      this.g.AddNode(this.graphName, dstName, this.currentNodeAttrs.Copy())
      //debug("     AddNode: " + dstName)
    } else {
      // debug("     Peek: " + stmt.EdgeRHS[i].Destination.String())
      // spew.Dump("     -Destination: ", stmt.EdgeRHS[i].Destination)
      // spew.Dump("     -StmtList: ", stmt.EdgeRHS[i].Destination.(*ast.SubGraph).StmtList)
      d := stmt.EdgeRHS[i].Destination.(*ast.SubGraph).StmtList
      for j := range d {
        //spew.Dump(d[j]) // (*ast.NodeStmt)
        dstid := d[j].(*ast.NodeStmt).NodeId
        dstName = dstid.String()
        //spew.Dump(dstid, dstName)
        this.g.AddNode(this.graphName, dstName, this.currentNodeAttrs.Copy())
        dstPort := stmt.EdgeRHS[i].Destination.GetPort()
        this.g.AddEdge(srcName, srcPort.String(), dstName, dstPort.String(), directed, attrs)
      }
      continue
    }
    dstPort := stmt.EdgeRHS[i].Destination.GetPort()
    this.g.AddEdge(srcName, srcPort.String(), dstName, dstPort.String(), directed, attrs)
    src = dst
    srcPort = dstPort
    srcName = dstName
  }
  return this
}

func (this *stmtVisitor) nodeAttrs(stmt ast.NodeAttrs) ast.Visitor {
  this.currentNodeAttrs = overwrite(this.currentNodeAttrs, ast.AttrList(stmt).GetMap())
  return &nilVisitor{}
}

func (this *stmtVisitor) edgeAttrs(stmt ast.EdgeAttrs) ast.Visitor {
  this.currentEdgeAttrs = overwrite(this.currentEdgeAttrs, ast.AttrList(stmt).GetMap())
  return &nilVisitor{}
}

func (this *stmtVisitor) graphAttrs(stmt ast.GraphAttrs) ast.Visitor {
  attrs := ast.AttrList(stmt).GetMap()
  for key, value := range attrs {
    this.g.AddAttr(this.graphName, key, value)
  }
  this.currentGraphAttrs = overwrite(this.currentGraphAttrs, attrs)
  return &nilVisitor{}
}

func (this *stmtVisitor) subGraph(stmt *ast.SubGraph) ast.Visitor {
  subGraphName := stmt.Id.String()
  this.g.AddSubGraph(this.graphName, subGraphName, this.currentGraphAttrs)
  return newStmtVisitor(this.g, subGraphName)
}

func (this *stmtVisitor) attr(stmt *ast.Attr) ast.Visitor {
  this.g.AddAttr(this.graphName, stmt.Field.String(), stmt.Value.String())
  return this
}
