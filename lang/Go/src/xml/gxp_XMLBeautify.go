////////////////////////////////////////////////////////////////////////////
// Porgram: gxp_XMLBeautify.go
// Purpose: goxpath XML Beautify
// Authors: Tong Sun (c) 2016, All rights reserved
// Credits: ChrisTrenkamp https://github.com/ChrisTrenkamp/goxpath/issues/7
////////////////////////////////////////////////////////////////////////////

package main

import (
	"os"

	"github.com/ChrisTrenkamp/goxpath"
	"github.com/ChrisTrenkamp/goxpath/tree/xmltree"
)

func main() {

	parseTree := xmltree.MustParseXML(os.Stdin, func(s *xmltree.ParseOptions) {
		s.Strict = false
	})
	goxpath.Marshal(parseTree, os.Stdout)
}

/*

echo "<root><this><is>a</is><test /></this></root>" | go run gxp_XMLBeautify.go
<root><this><is>a</is><test></test></this></root>

*/
