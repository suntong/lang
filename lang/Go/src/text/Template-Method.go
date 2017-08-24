/*

Call a method from a Go template
https://stackoverflow.com/questions/10200178/call-a-method-from-a-go-template

Q: How to use method from a template?

A: Just omit the parentheses and it should be fine.

According to the documentation, you can call any method which returns one
value (of any type) or two values if the second one is of type error. In the
later case, Execute will return that error if it is non-nil and stop the
execution of the template.

*/

package main

import (
	"html/template"
	"log"
	"os"
)

type Person string

func (p Person) Label() string {
	return "This is " + string(p)
}

// You can even pass parameters to function

func (p Person) Label2(param1, param2 string) string {
	return "This is " + string(p) + " - " + param1 + "," + param2
}

func main() {
	tmpl, err := template.New("").Parse("{{.Label}}\n" +
		`{{ .Label2 "value1" "value2"}}`)
	if err != nil {
		log.Fatalf("Parse: %v", err)
	}
	tmpl.Execute(os.Stdout, Person("Bob"))
}

/*

This is Bob
This is Bob - value1,value2

*/
