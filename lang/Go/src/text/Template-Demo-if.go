// https://golang.org/pkg/text/template/#example_Template
package main

import (
	"os"
	"text/template"
)

type Inventory struct {
	Material, Thematerial string
	Count                 uint
}

func main() {
	{
		sweaters := Inventory{"", "", 17}
		tmpl, err := template.New("test").
			Parse(`{{$.Thematerial := "something"}}{{if .Material}} {{$.Thematerial := .Material}} {{end}}` +
				"{{.Count}} items are made of {{$.Thematerial}}.\n")
		check(err)
		err = tmpl.Execute(os.Stdout, sweaters)
		check(err)
	}
	{
		sweaters := Inventory{"wool", "", 17}
		tmpl, err := template.New("test").
			Parse(`{{$material := "something"}}{{if .Material}} {{$material := .Material}} {{end}}` +
				"{{.Count}} items are made of {{$material}}.\n")
		check(err)
		err = tmpl.Execute(os.Stdout, sweaters)
		check(err)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
