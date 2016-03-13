package main

import (
	"fmt"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Foo string
	Bar []string
}

type Configs struct {
	Cfgs []Config `foobars`
}

var data = `
foobars:
 - foo: 1
   bar:
    - one
    - two
    - three

 - foo: 2
   bar:
    - one1
    - two2
    - three3
`
var data0 = `
 - foo: 1
   bar:
    - one
    - two
    - three

 - foo: 2
   bar:
    - one1
    - two2
    - three3
`

func main() {

	///////////////////////////////////////////////////////////////
	// Unmarshal

	var config Configs

	/*
	   filename := os.Args[1]
	   source, err := ioutil.ReadFile(filename)
	   if err != nil {
	       panic(err)
	   }
	*/

	source := []byte(data)

	err := yaml.Unmarshal(source, &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	//fmt.Printf("Value: %#v\n", config.Bar[0])

	fmt.Printf("--- config:\n%v\n\n", config)
	// --- config:
	// {[{1 [one two three]} {2 [one1 two2 three3]}]}

	///////////////////////////////////////////////////////////////
	// Marshal

	type T struct {
		F int "a,omitempty"
		B int
	}

	r, _ := yaml.Marshal(&T{B: 2}) // Returns "b: 2\n"
	fmt.Printf("--- T:\n%s\n\n", string(r))

	r, _ = yaml.Marshal(&T{F: 1}) // Returns "a: 1\nb: 0\n"
	fmt.Printf("--- T:\n%s\n\n", string(r))

	/*
			err = yaml.Unmarshal([]byte(data0), &config)

		  yaml: unmarshal errors:
		  line 2: cannot unmarshal !!seq into main.Configs

			if err != nil {
				log.Fatalf("error: %v", err)
			}
			fmt.Printf("--- config:\n%v\n\n", config)
	*/

}
