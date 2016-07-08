////////////////////////////////////////////////////////////////////////////
// Porgram: dotu0.go
// Purpose: Demo the .dot (graphviz) file handling in GO
// Authors: Tong Sun (c) 2016, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"

	"github.com/awalterschulze/gographviz"
)

var dotStr = `
digraph callgraph {
  "(*deduper/minhash.MinHasher).bandMatrix" -> "(*deduper/minhash.MinHasher).bandColumn"
  "(*deduper/server.Server).ListenAndServe" -> "deduper/server/middleware.NewLeadWrite"
  "(*deduper/server.Server).ListenAndServe" -> "(*deduper/server.Server).connectionString"
  "(*deduper/server/middleware.LeaderWrite).ServeHTTP" -> "(*deduper/server/middleware.LeaderWrite).matches"
}

`

func main() {
	//g, err := dot.ParseFile("dotu0.dot")
	g, err := gographviz.Read([]byte(dotStr))
	if err != nil {
		panic(err)
	}
	s := g.String()
	fmt.Println(s)
}

/*

go run dotu0.go > /tmp/dcg-1.dot
dot -Tpng /tmp/dcg-1.dot > /tmp/dcg-dn.png

*/
