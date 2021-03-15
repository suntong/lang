package main

import (
	"bytes"
	"fmt"

	"github.com/ChrisTrenkamp/goxpath"
	"github.com/ChrisTrenkamp/goxpath/tree"
	"github.com/ChrisTrenkamp/goxpath/tree/xmltree"

	"github.com/suntong/testing"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

type namespace map[string]string

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

var ns = make(namespace)

////////////////////////////////////////////////////////////////////////////
// Function definitions

//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
// Function main

func main() {
	var t *testing.T = testing.NewT()
	TestFollowingSibling(t)
	t.Report()
	TestFollowingSibling2(t)
	t.Report()
	TestFollowingSiblingA1(t)
	t.Report()
	TestFollowingSiblingA2(t)
	t.Report()
	TestFollowingSiblingA3(t)
	t.Report()
}

//==========================================================================
// test functions

func TestFollowingSibling(t *testing.T) {
	p := `//p2/following-sibling::node()`
	x := `<?xml version="1.0" encoding="UTF-8"?><p1><p21/><p2><p3/><p4/></p2><p5><p6/></p5></p1>`
	exp := []string{`<p5><p6></p6></p5>`}
	execPath(p, x, exp, nil, t)
}

func TestFollowingSibling2(t *testing.T) {
	p := `/following-sibling::node()`
	x := `<?xml version="1.0" encoding="UTF-8"?><p1></p1><p2 />`
	exp := []string{}
	execPath(p, x, exp, nil, t)
}

// https://www.w3schools.com/xml/xpath_axes.asp
// following-sibling - Selects *all* siblings after the current node
func TestFollowingSiblingA1(t *testing.T) {
	p := `/AAA/BBB/following-sibling::node()`
	x := `<?xml version="1.0" encoding="UTF-8"?><AAA><BBB><CCC /><DDD /></BBB><XXX><DDD><EEE /><FFF><GGG /></FFF><XXX /></DDD></XXX><CCC><DDD /></CCC></AAA>`
	exp := []string{`<XXX><DDD><EEE></EEE><FFF><GGG></GGG></FFF><XXX></XXX></DDD></XXX>`,
		`<CCC><DDD></DDD></CCC>`}
	execPath(p, x, exp, nil, t)
}

func TestFollowingSiblingA2(t *testing.T) {
	p := `//EEE/following-sibling::node()`
	x := `<?xml version="1.0" encoding="UTF-8"?><AAA><BBB><CCC /><DDD /></BBB><XXX><DDD><EEE /><FFF /><FFF><GGG /></FFF><XXX /></DDD></XXX><CCC><DDD /></CCC></AAA>`
	exp := []string{`<FFF></FFF>`,
		`<FFF><GGG></GGG></FFF>`, `<XXX></XXX>`}
	execPath(p, x, exp, nil, t)
}

func TestFollowingSiblingA3(t *testing.T) {
	// https://stackoverflow.com/questions/10246371/
	p := `//project[title = 'Project X']/following-sibling::project[1]`
	x := `<?xml version="1.0" encoding="UTF-8"?>
<projects>
    <project>
        <number>1</number>
        <title>Project X</title>
    </project>
    <project>
        <number>2</number>
        <title>Project Y</title>
    </project>
    <project>
        <number>3</number>
        <title>Project Z</title>
    </project>
</projects>
`
	exp := []string{`<project>
        <number>2</number>
        <title>Project Y</title>
    </project>`}
	execPath(p, x, exp, nil, t)
}

//==========================================================================
// support functions

func execPath(xp, x string, exp []string, ns map[string]string, t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Error("Panicked: from XPath expr: '" + xp)
			t.Error(r)
			//t.Error(string(debug.Stack()))
		}
	}()
	res := goxpath.MustParse(xp).MustExec(xmltree.MustParseXML(bytes.NewBufferString(x)), func(o *goxpath.Opts) { o.NS = ns }).(tree.NodeSet)

	if len(res) != len(exp) {
		t.Error("Result length not valid in XPath expression '"+xp+"':", len(res), ", expecting", len(exp))
		for i := range res {
			t.Error(goxpath.MarshalStr(res[i].(tree.Node)))
		}
		return
	}

	for i := range res {
		fmt.Printf("%#v\n", res[i].(tree.Node))
		r, err := goxpath.MarshalStr(res[i].(tree.Node))
		if err != nil {
			t.Error(err.Error())
			return
		}
		valid := false
		for j := range exp {
			if r == exp[j] {
				valid = true
			}
		}
		if !valid {
			t.Error("Incorrect result in XPath expression '" + xp + "':" + r)
			t.Error("Expecting one of:")
			for j := range exp {
				t.Error(exp[j])
			}
			return
		}
	}
}

/*

$ go run gxp_Path.go
--- PASS:  (0.00s)
--- PASS:  (0.00s)

*/
