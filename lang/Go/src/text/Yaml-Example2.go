////////////////////////////////////////////////////////////////////////////
// Porgram: Yaml-Example2
// Purpose: Go Yaml Unmarshal Example
// Authors: Tong Sun (c) 2016, All rights reserved
// Credits: https://mlafeldt.github.io/blog/decoding-yaml-in-go/
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"log"

	"gopkg.in/yaml.v2"
)

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

// Note: there can be more data than in struct (last_action)
var data = `
hostname: 127.0.0.1
username: vagrant
ssh_key: "/long/path/to/private_key"
port: '2222'
last_action: create
`

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

type instanceConfig struct {
	Hostname string
	Username string
	SSHKey   string `yaml:"ssh_key"`
	Port     string
}

////////////////////////////////////////////////////////////////////////////
// Function definitions

func (c *instanceConfig) Parse(data []byte) error {
	return yaml.Unmarshal(data, c)
}

//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
// Function main

func main() {
	test1()

	test2([]byte(data2))
	test2([]byte(data22))

	test3([]byte(data2))
}

//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
// test1

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

//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
// test2

var data2 = `
hostname: 127.0.0.1
m:
  username: vagrant
  ssh_key: "/long/path/to/private_key"
  port: '2222'
  last_action: create
`

// Note: there can be less data than in struct too (m)
var data22 = `
hostname: 127.0.0.1
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

//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
// test3

/*

The problem with type mapConfig is that,
the order within the M map cannot guarantee to be as the original order.

The solution -- type yaml.MapSlice in next example

MapSlice encodes and decodes as a YAML map.
The order of keys is preserved when encoding and decoding.


*/

type mapSlice struct {
	Hostname string
	M        yaml.MapSlice
}

func test3(source []byte) {
	var config mapSlice

	err := yaml.Unmarshal(source, &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Printf("--- config:\n%+v\n\n", config)
}

/*

Output:

--- config:
{Hostname:127.0.0.1 Username:vagrant SSHKey:/long/path/to/private_key Port:2222}

--- config:
{Hostname:127.0.0.1 M:map[username:vagrant ssh_key:/long/path/to/private_key port:2222 last_action:create]}

--- config:
{Hostname:127.0.0.1 M:map[]}

--- config:
{Hostname:127.0.0.1 M:[{Key:username Value:vagrant} {Key:ssh_key Value:/long/path/to/private_key} {Key:port Value:2222} {Key:last_action Value:create}]}

*/
