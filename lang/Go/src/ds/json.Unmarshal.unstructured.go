// From:

package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	test1()
	test2()
	test3()
}

func test1() {
	// from https://blog.golang.org/json-and-go
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	var f interface{}
	//err :=
	json.Unmarshal(b, &f)
	fmt.Printf("--- test1::f:\n%v\n\n", f)

	m := f.(map[string]interface{})
	fmt.Printf("--- test1::m:\n%v\n\n", m)

}

func test2() {
	// from https://stackoverflow.com/questions/30341588/
	b := []byte(`{
   "k1" : "v1", 
   "k3" : 10,
   "result":["v4",12.3,{"k11" : "v11", "k22" : "v22"}]
}`)
	var f interface{}
	//err :=
	json.Unmarshal(b, &f)
	// fmt.Printf("--- test2::e:\n%v\n\n", err)
	fmt.Printf("--- test2::f:\n%v\n\n", f)
}

func test3() {
	// from https://www.sohamkamani.com/blog/2017/10/18/parsing-json-in-golang/#unstructured-data
	birdJson := `{"birds":{"pigeon":"likes to perch on rocks","eagle":"bird of prey"},"animals":"none"}`
	var result map[string]interface{}
	json.Unmarshal([]byte(birdJson), &result)

	// The object stored in the "birds" key is also stored as
	// a map[string]interface{} type, and its type is asserted from
	// the interface{} type
	birds := result["birds"].(map[string]interface{})

	fmt.Printf("--- test3\n")
	for key, value := range birds {
		// Each value is an interface{} type, that is type asserted as a string
		fmt.Println(key, value.(string))
	}
}

/*

$ go run json.Unmarshal.unstructured.go
--- test1::f:
map[Name:Wednesday Age:6 Parents:[Gomez Morticia]]

--- test1::m:
map[Name:Wednesday Age:6 Parents:[Gomez Morticia]]

--- test2::f:
map[k3:10 result:[v4 12.3 map[k11:v11 k22:v22]] k1:v1]

--- test3
pigeon likes to perch on rocks
eagle bird of prey

*/
