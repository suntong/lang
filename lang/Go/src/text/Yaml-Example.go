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

func main() {

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
}
