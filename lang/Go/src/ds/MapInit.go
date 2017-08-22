package main

import (
	"fmt"
)

var b2s = map[bool]string{
	true:  "foo",
	false: "bar",
}

func test1() {
	meds := map[string]int{
		"Ativan":   15,
		"Xanax":    20,
		"Klonopin": 30,
	}
	fmt.Println(meds)
}

func main() {
	fmt.Println(b2s)
	test1()
}
