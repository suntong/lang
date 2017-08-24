// https://play.golang.org/p/aH2fWgn-a0

package main

import (
	"bytes"
	"os"
	"text/template"
)

type Person struct {
	Name string //exported field since it begins with a capital letter
}

func main() {
	t := template.New("hello template")  //create a new template with some name
	t, _ = t.Parse("hello {{.Name}}!\n") //parse some content and generate a template, which is an internal representation

	p := Person{Name: "Mary"} //define an instance with required field

	t.Execute(os.Stdout, p) //merge template ‘t’ with content of ‘p’

	type dict struct {
		// struct fields must be public
		Title   string
		Release int
	}
	params := dict{Title: "Go", Release: 60}
	t, _ = template.New("template_name").Parse(
		"<h1>{{.Title}}</h1>r{{.Release}}\n")
	buf := new(bytes.Buffer)
	t.Execute(buf, params)
	print(buf.String())

	t = template.New("template test")
	t = template.Must(t.Parse("This is just static text. \n{{\"This is pipeline data - because it is evaluated within the double braces.\"}} {{`So is this, but within reverse quotes {{}}.`}}\n"))
	t.Execute(os.Stdout, nil)

}

/*

hello Mary!
<h1>Go</h1>r60
This is just static text.
This is pipeline data - because it is evaluated within the double braces. So is this, but within reverse quotes {{}}.

*/
