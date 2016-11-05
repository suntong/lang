////////////////////////////////////////////////////////////////////////////
// Porgram: gxp_XMLBeautify.go
// Purpose: goxpath XML Beautify
// Authors: Tong Sun (c) 2016, All rights reserved
// Credits: ChrisTrenkamp https://github.com/ChrisTrenkamp/goxpath/issues/5
////////////////////////////////////////////////////////////////////////////

package main

import (
	"os"

	"github.com/ChrisTrenkamp/goxpath"
	"github.com/ChrisTrenkamp/goxpath/tree/xmltree"
)

func main() {

	parseTree := xmltree.MustParseXML(os.Stdin)
	goxpath.Marshal(parseTree, os.Stdout)
}

/*

go run gxp_XMLBeautify.go

*/
