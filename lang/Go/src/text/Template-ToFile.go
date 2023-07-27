package main

import (
	"os"
	"text/template"
)

func main() {

	// create a new file (overwrite)
	file, _ := os.Create("/tmp/greeting.txt")
	defer file.Close()

	// variables
	vars := make(map[string]interface{})
	vars["Greeting"] = "Hello"
	vars["Name"] = "John"

	// parse & apply the template to the vars map and write the result to file.
	tmpl, _ := template.New("1").Parse(template1)
	tmpl.Execute(file, vars)
	// another one
	tmpl = template.Must(template.New("2").Parse(template2))
	tmpl.Execute(file, vars)
}

const (
	template1 = "{{.Greeting}} {{.Name}}!\n"
	template2 = "{{.Greeting}} my friend {{.Name}}!\n"
)
