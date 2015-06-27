////////////////////////////////////////////////////////////////////////////
// Porgram: Yaml-Template
// Purpose: Go templating from Yaml example
// Authors: Tong Sun (c) 2015, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"log"
	"os"
	"text/template"

	"gopkg.in/yaml.v2"
)

/*
# error: yaml: did not find expected key

# Hashes
- name: Xavier
  age: 24
# in flow style...
- {name: Xavier2, age: 22}
*/

const data = `
A: Easy!
B:
  C: 2
  D: [\"3A\", \"4B\"]

# Arrays
Colors1:
  - red
  - blue
# in flow style...
Colors2: [red, blue]
Colors3: ['"red"', '"blue"']

# Hashes
H: 
  - Name: Xavier
    Age: 24
  # in flow style...
  - {Name: Xavier2, Age: 22}

`

const templ = `Templating from Yaml with Go is {{.A}}.
{{range .H}}
  The name is {{.Name}}.
  The age is {{.Age}}.
{{end}}
The colors are: {{range .Colors1}}{{.}}, {{end}}.
The quoted colors are: {{range .Colors3}}{{.}}, {{end}}.
D values are: {{range .B.D}}{{.}}, {{end}}.
`

func main() {

	m := make(map[interface{}]interface{})

	err := yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- m:\n%v\n\n", m)

	d, err := yaml.Marshal(&m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- m dump:\n%s\n\n", string(d))

	t := template.New("Test template")
	t, err = t.Parse(templ)
	checkError(err)

	err = t.Execute(os.Stdout, m)
	checkError(err)

}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}

/*
$ go run Yaml-Template.go
--- m:
map[A:Easy! B:map[C:2 D:[\"3A\" \"4B\"]] Colors1:[red blue] Colors2:[red blue] Colors3:["red" "blue"] H:[map[Name:Xavier Age:24] map[Name:Xavier2 Age:22]]]

--- m dump:
A: Easy!
B:
  C: 2
  D:
  - \"3A\"
  - \"4B\"
Colors1:
- red
- blue
Colors2:
- red
- blue
Colors3:
- '"red"'
- '"blue"'
H:
- Age: 24
  Name: Xavier
- Age: 22
  Name: Xavier2


Templating from Yaml with Go is Easy!.

  The name is Xavier.
  The age is 24.

  The name is Xavier2.
  The age is 22.

The colors are: red, blue, .
The quoted colors are: "red", "blue", .
D values are: \"3A\", \"4B\", .

*/
