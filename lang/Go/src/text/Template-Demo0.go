// https://golang.org/pkg/text/template/#example_Template
package main

import (
	"os"
	"text/template"
)

type Inventory struct {
	Material string
	Count    uint
}

func main() {
	{
		sweaters := Inventory{"wool", 17}
		tmpl, err := template.New("test").
			Parse("{{.Count}} items are made of {{.Material}}\n")
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
