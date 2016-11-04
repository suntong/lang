////////////////////////////////////////////////////////////////////////////
// Porgram: gxp_Edit.go
// Purpose: goxpath xml element attribute change demo
// Authors: Tong Sun (c) 2016, All rights reserved
// Credits: ChrisTrenkamp https://github.com/ChrisTrenkamp/goxpath/issues/5
////////////////////////////////////////////////////////////////////////////

/*

Demo to use goxpath to,

- locate all //Request nodes; and for each node found,
- change its attribute according to other attributes.

For example, the following changes were to change

//Request's ./[ReportingName] use the basename of ./[Url](and also prefixing it with ../TransactionTimer[Name] as well).

*/

package main

import (
	"encoding/xml"
	"net/url"
	"os"
	"path"

	"github.com/ChrisTrenkamp/goxpath"
	"github.com/ChrisTrenkamp/goxpath/tree"
	"github.com/ChrisTrenkamp/goxpath/tree/xmltree"
	"github.com/ChrisTrenkamp/goxpath/tree/xmltree/xmlnode"
)

func basenameFunc(c tree.Ctx, args ...tree.Result) (tree.Result, error) {
	parsedURL, err := url.Parse(args[0].String())
	return tree.String(path.Base(parsedURL.Path)), err
}

func register(o *goxpath.Opts) {
	o.NS["ms"] = "http://microsoft.com/schemas/VisualStudio/TeamTest/2010"
	o.Funcs[xml.Name{Local: "basename"}] = tree.Wrap{Fn: basenameFunc, NArgs: 1}
}

func main() {
	// no first command line arguments
	if len(os.Args) <= 1 {
		println("Usage:\r\n input.xml")
	}

	urlBasename := goxpath.MustParse(`basename(../@Url)`)
	transTimer := goxpath.MustParse(`ancestor::ms:TransactionTimer/@Name`)

	f, err := os.Open(os.Args[1])

	if err != nil {
		panic(err)
	}

	parseTree := xmltree.MustParseXML(f)
	reportingNames, err := goxpath.MustParse(`//ms:Request/@ReportingName`).ExecNode(parseTree, register)

	if err != nil {
		panic(err)
	}

	for _, i := range reportingNames {
		attr := i.(xmlnode.XMLNode)
		val := attr.Token.(*xml.Attr)

		val.Value = transTimer.MustExec(attr, register).String() + " - " +
			urlBasename.MustExec(attr, register).String()
	}

	goxpath.Marshal(parseTree, os.Stdout)
}

// go run gxp_Edit.go gxp_Edit_i.xml > gxp_Edit_o.xml
