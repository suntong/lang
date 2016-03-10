////////////////////////////////////////////////////////////////////////////
// Porgram: Piping_Pipe.go
// Purpose: Go Internal piping with io.Pipe
// Authors: Tong Sun (c) 2016, All rights reserved
// Credits: https://mlafeldt.github.io/blog/decoding-yaml-in-go/
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"log"

	"gopkg.in/yaml.v2"
)

var data = `
hostname: 127.0.0.1
username: vagrant
ssh_key: "/long/path/to/private_key"
port: '2222'
last_action: create
`

type instanceConfig struct {
	Hostname string
	Username string
	SSHKey   string `yaml:"ssh_key"`
	Port     string
}

func (c *instanceConfig) Parse(data []byte) error {
	return yaml.Unmarshal(data, c)
}

func main() {
	test1()
	test2([]byte(data2))
}

func test1() {

	/*
	   filename := os.Args[1]
	   source, err := ioutil.ReadFile(filename)
	   if err != nil {
	       panic(err)
	   }
	*/
	source := []byte(data)

	var config instanceConfig

	err := config.Parse(source)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Printf("--- config:\n%+v\n\n", config)
}

var data2 = `
hostname: 127.0.0.1
m:
  username: vagrant
  ssh_key: "/long/path/to/private_key"
  port: '2222'
  last_action: create
`

type mapConfig struct {
	Hostname string
	M        map[string]string
}

func test2(source []byte) {
	var config mapConfig

	err := yaml.Unmarshal(source, &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Printf("--- config:\n%+v\n\n", config)
}
