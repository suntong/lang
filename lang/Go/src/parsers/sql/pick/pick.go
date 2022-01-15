package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

func main() {
	// read whole the file
	b, err := ioutil.ReadFile("sql-2016.ebnf")
	if err != nil {
		panic(err)
	}

	re := regexp.MustCompile("1..")

	// Find and loop over all matching strings.
	results := re.FindAllString("123 124 125 200 211", -1)
	for i := range results {
		fmt.Println(results[i])
	}
}
