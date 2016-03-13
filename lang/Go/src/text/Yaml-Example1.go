////////////////////////////////////////////////////////////////////////////
// Porgram: Yaml-Example
// Purpose: Go Yaml example
// Authors: Tong Sun (c) 2016, All rights reserved
// Credict: https://github.com/go-yaml/yaml
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"log"

	"gopkg.in/yaml.v2"
)

var data = `
  d: [3, 4]
`

type T struct {
	D []int // `yaml:",flow"`
}

func main() {
	t := T{}

	err := yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t:\n%v\n\n", t)

	d, err := yaml.Marshal(&t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t dump:\n%s\n\n", string(d))

	m := make(map[interface{}]interface{})

	err = yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- m:\n%v\n\n", m)

	d, err = yaml.Marshal(&m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- m dump:\n%s\n\n", string(d))
}

/*
This example will generate the following output:

--- t:
{[3 4]}

--- t dump:
d:
- 3
- 4


--- m:
map[d:[3 4]]

--- m dump:
d:
- 3
- 4

With the `yaml:",flow"` flags, the "t dump" output will be:

--- t dump:
d: [3, 4]


*/
