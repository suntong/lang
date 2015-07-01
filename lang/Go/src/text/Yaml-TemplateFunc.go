////////////////////////////////////////////////////////////////////////////
// Porgram: Yaml-Template
// Purpose: Go templating with FuncMap from Yaml example
// Authors: Tong Sun (c) 2015, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	"gopkg.in/yaml.v2"
)

const data = `
Colors:
  - red
  - blue
  - white
`

const templ1 = `{{range .Colors}}{{.}}, {{end}}`
const templ2 = `{{join .Colors}}`

func main() {

	m := make(map[interface{}]interface{})

	err := yaml.Unmarshal([]byte(data), &m)
	checkError(err)
	fmt.Printf("--- m:\n%v\n\n", m)

	d, err := yaml.Marshal(&m)
	checkError(err)
	fmt.Printf("--- m dump:\n%s\n\n", string(d))

	t := template.New("Test template")
	t, err = t.Parse(templ1)
	checkError(err)

	err = t.Execute(os.Stdout, m)
	checkError(err)

	colors := []string{"red", "white", "blue"}
	t = template.Must(template.New("").Funcs(funcs).Parse(templ2))
	err = t.Execute(os.Stdout, struct{ Colors []string }{colors})
	checkError(err)

	/*

	   The following is impossible -- you can only type assert for interface{}.
	   For []interface{}, you have to do it element by element.

	   http://play.golang.org/p/Nq8efe5eWh

	   Kiki Sugiaman

	   	fmt.Println("\nNext")
	   	fmt.Println(strings.Join(m["Colors"].([]string), ", "))
	   	err = t.Execute(os.Stdout, struct{ Colors []string }{m["Colors"].([]string)})
	   	checkError(err)
	*/

}

var funcs = template.FuncMap{
	"join": func(s []string) string { return strings.Join(s, ", ") },
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
