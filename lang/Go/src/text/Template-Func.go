////////////////////////////////////////////////////////////////////////////
// Porgram: Template-Func.go
// Purpose: Go html template FuncMap usage demo
// Authors: Tong Sun (c) 2015, All rights reserved
// Based on http://play.golang.org/p/WpnEXoF-6V by ksug
////////////////////////////////////////////////////////////////////////////

package main

import (
	"html/template"
	"os"
)

var funcMap = template.FuncMap{
	"whatEver": sayHello,
	"add":      add,
}

var regularMap = map[string]interface{}{
	"foo": bar(),
}

func sayHello() string {
	return "Hello World"
}

func add(a int, b int) int {
	return a + b
}

func bar() string {
	return "Baz"
}

func main() {
	tmpl, _ := template.New("base").Funcs(funcMap).Parse("{{whatEver}} - {{.foo}} {{add 2 3}}\n")
	tmpl.Execute(os.Stdout, regularMap)

	t, err := template.New("Template-Func.tmpl").
		Funcs(funcMap).ParseFiles("Template-Func.tmpl")
	if nil != err {
		panic(err)
	}
	t.Execute(os.Stdout, regularMap)

}
