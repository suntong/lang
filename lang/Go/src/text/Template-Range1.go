////////////////////////////////////////////////////////////////////////////
// Porgram: Template-Range
// Purpose: Go template "range" demo
// Authors: Tong Sun (c) 2017, All rights reserved
// Credits: Dave C, https://stackoverflow.com/a/27475244/2125837
////////////////////////////////////////////////////////////////////////////

package main

import (
	"os"
	"text/template"
)

type Context struct {
	People []Person // all fields referenced by the template must be exported
}

type Person struct {
	Name   string
	Senior bool
}

// Static regexp and templates are often initialized/compiled/parsed at
// program initialization and then executed as required. One advantage
// of this is that any panics from template.Must (or regexp.MustCompile)
// happen imediately on startup instead of requiring the appropriate
// function actually get called. I.e. fail early.
//
// It's fine to move this to some kind of initialisation function
// especially if it's a big template and only rarely gets used. But it
// really should not be in a function that gets called repeatedly or you
// parse the same static text over and over (of course if it was somehow
// a dynamic template that would be different).
//
// Another way to defer parsing until required, but avoid re-parsing is
// with sync.Once.
var tRange = template.Must(template.New("Range Example").
	Parse(`
{{range $i, $x := $.People}}Name={{$x.Name}} Senior={{$x.Senior}}
{{end}}
Or:
{{range .People}}Name={{.Name}} {{if .Senior}}is{{else}}is not{{end}} a senior
{{end}}`))

func main() {
	// Note that unlike your initialization this one leaves
	// out the "Person" type on each item. It is not needed
	// here and is not recommended (gofmt -s will remove it).
	//
	// You optionally could also remove the field names People, Name,
	// and Senior. That's often done when the initialization happens
	// in the same package as the type definition. However, it's
	// considered good practice to use field names when the type
	// comes from some other package (go vet will warn otherwise).
	ctx2 := Context{
		People: []Person{
			{Name: "Mary", Senior: false},
			{Name: "Joseph", Senior: true},
		},
	}
	tRange.Execute(os.Stdout, ctx2)
}

/*

Name=Mary Senior=false
Name=Joseph Senior=true

Or:
Name=Mary is not a senior
Name=Joseph is a senior

*/
