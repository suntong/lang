// Sample code that crashes Go.
package main

import "fmt"

type BaseRecord struct {
	RawLine string
	Index   string
	Valid   bool
}

func (b BaseRecord) String() string {
	// X: return fmt.Sprintf("%+v", b)
	return fmt.Sprintf("ID: %s, Status: %v", b.Index, b.Valid)
}

func main() {
	fmt.Println("Hello, 世界")

	r := BaseRecord{
		Index:   "2",
		RawLine: "Test",
		Valid:   false,
	}
	//t := fmt.Sprintf("%+v", r)
	t := fmt.Sprintf("%v", r)
	_ = t
	//fmt.Sprintf("%s", t)
}
