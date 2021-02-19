// https://play.golang.org/p/QZPqgd-g7mf
// go test -v xpath_example_test.go

// Copyright 2017 Santhosh Kumar Tekuri. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/santhosh-tekuri/dom"
	"github.com/santhosh-tekuri/xpath"
)

func Example() {
	str := `
	<developer>
		<name>Santhosh Kumar Tekuri</name>
		<email>santhosh.tekuri@gmail.com</email>
	</developer>
	`
	doc, err := dom.Unmarshal(xml.NewDecoder(strings.NewReader(str)))
	if err != nil {
		fmt.Println(err)
		return
	}

	expr, err := new(xpath.Compiler).Compile("/developer/name")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("xpath %v returns value of type %v\n", expr, expr.Returns())

	result, err := expr.EvalString(doc, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Result: %s", result)
	// Output:
	// xpath /developer/name returns value of type node-set
	// Result: Santhosh Kumar Tekuri
}

func ExampleVariableMap() {
	uri := "www.jroller.com/santhosh/"

	compiler := &xpath.Compiler{
		Namespaces: map[string]string{
			"ns": uri,
		},
	}
	expr, err := compiler.Compile("$v1 + $v2 * $ns:v3 - $ns:v4")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("xpath %v returns value of type %v\n", expr, expr.Returns())

	result, err := expr.EvalNumber(nil, xpath.VariableMap{
		"v1":                            float64(2),
		"v2":                            float64(3),
		"{www.jroller.com/santhosh/}v3": float64(4),
		xpath.ClarkName(uri, "v4"):      float64(1),
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Result: %.2f", result)
	// Output:
	// xpath $v1 + $v2 * $ns:v3 - $ns:v4 returns value of type number
	// Result: 13.00
}

func ExampleFunctionMap() {
	join := func(args []interface{}) interface{} {
		sep := args[0].(string)
		var a []string
		for _, v := range args[1:] {
			a = append(a, v.(string))
		}
		return strings.Join(a, sep)
	}

	uri := "www.jroller.com/santhosh/"

	compiler := &xpath.Compiler{
		Namespaces: map[string]string{
			"x": uri,
		},
		Functions: xpath.FunctionMap{
			"{www.jroller.com/santhosh/}join": &xpath.Function{
				Returns: xpath.String,
				Args: xpath.Args{
					xpath.Mandatory(xpath.String),
					xpath.Variadic(xpath.String),
				},
				Compile: xpath.CompileFunc(join),
			},
		},
	}
	expr, err := compiler.Compile("x:join(':', 'one', 'two', 'three')")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("xpath %v returns value of type %v\n", expr, expr.Returns())

	result, err := expr.EvalString(nil, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Result: %s", result)
	// Output:
	// xpath x:join(':', 'one', 'two', 'three') returns value of type string
	// Result: one:two:three
}

// https://github.com/santhosh-tekuri/xpath/issues/1
func ExampleIssueNo1() {
	str :=
		`
<simservs xmlns:cp="urn:ietf:params:xml:ns:common-policy" xmlns:ss="http://uri.etsi.org/ngn/params/xml/simservs/xcap">
  <ss:communication-diversion ss:active="true">
    <cp:ruleset>
      <cp:rule id="CFU">
        <cp:conditions></cp:conditions>
        <cp:actions>
          <ss:forward-to>
            <ss:target>tel:+4792869749</ss:target>
          </ss:forward-to>
        </cp:actions>
      </cp:rule>
    </cp:ruleset>
  </ss:communication-diversion>
</simservs>
`
	doc, err := dom.Unmarshal(xml.NewDecoder(strings.NewReader(str)))
	if err != nil {
		panic(err)
	}

	ns := map[string]string{
		"cp": "urn:ietf:params:xml:ns:common-policy",
		"ss": "http://uri.etsi.org/ngn/params/xml/simservs/xcap",
	}

	q := "/simservs/ss:communication-diversion/cp:ruleset/cp:rule[@id=\"CFU\"]"
	expr, err := (&(xpath.Compiler{Namespaces: ns})).Compile(q)
	if err != nil {
		panic(err)
	}
	fmt.Println(expr.String())
	// reported BUG as:  nodeSet == nil
	nodeSet, err := expr.EvalNodeSet(doc, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(nodeSet)

	result, err := expr.EvalString(doc, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(":", strings.TrimSpace(result))
	// Output:
	// /simservs/ss:communication-diversion/cp:ruleset/cp:rule[@id="CFU"]
	// [cp:rule]
	// : tel:+4792869749
}
