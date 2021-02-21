////////////////////////////////////////////////////////////////////////////
// Porgram: gxp_XMLBeautify.go
// Purpose: goxpath XML Beautify
// Authors: Tong Sun (c) 2016, All rights reserved
// Credits: ChrisTrenkamp https://github.com/ChrisTrenkamp/goxpath/issues/7
////////////////////////////////////////////////////////////////////////////

package main

import (
	"flag"
	"os"

	"github.com/ChrisTrenkamp/goxpath"
	"github.com/ChrisTrenkamp/goxpath/tree/xmltree"
)

var strict = flag.Bool("s", false, "strict with xmltree parse")

func main() {

	flag.Parse()

	parseTree := xmltree.MustParseXML(os.Stdin, func(s *xmltree.ParseOptions) {
		s.Strict = *strict
	})
	goxpath.Marshal(parseTree, os.Stdout)
}

/*

echo "<root><this><is>a</is><test /></this></root>" | go run gxp_XML_RW.go
<root><this><is>a</is><test></test></this></root>

*/
