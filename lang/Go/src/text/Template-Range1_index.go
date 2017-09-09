////////////////////////////////////////////////////////////////////////////
// Porgram: Template-Range1_index
// Purpose: Go template "range" demo
// Authors: Tong Sun (c) 2017, All rights reserved
// Credits: Jan Mercl
//          https://groups.google.com/d/msg/golang-nuts/Cpp82FoOMyU/Udk9oh6AAgAJ
////////////////////////////////////////////////////////////////////////////

/*

Q:

If I have Go templates Range declared like this,

    {{range $i, $x := $.People}}

How can I test its index value?

I thought it should be `{{if $i ne 0}}` but it is not working.
I've also tried `{{if $i ne "0"}}`, but that failed on me as well.

A:

ne is a function, not an infix operator

https://golang.org/pkg/text/template/#hdr-Functions
https://play.golang.org/p/sJSzoNyrJz

Jan Mercl

*/

package main

import (
	"os"
	templ "text/template"
)

type Context struct {
	People []Person
}
type Person struct {
	Name   string //exported field since it begins with a capital letter
	Senior bool
}

func main() {
	// Range example
	tRange := templ.New("Range Example")
	ctx2 := Context{People: []Person{Person{Name: "Mary", Senior: false}, Person{Name: "Joseph", Senior: true}}}
	tRange = templ.Must(
		tRange.Parse(`
{{range $i, $x := $.People}} {{if ne $i 0}} Name={{$x.Name}} Senior={{$x.Senior}}{{end}}
{{end}}
`))
	tRange.Execute(os.Stdout, ctx2)
}
