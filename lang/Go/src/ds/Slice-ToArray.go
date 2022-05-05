/*

   - How to convert slice to fixed size array?
   - How to check if two slices are equal?

 https://tip.golang.org/ref/spec#Conversions_from_slice_to_array_pointer

*/

package main

import "fmt"

func main() {
	fmt.Println("Hello!")
	nums := []int{2, 3, 4, 2, 3, 4, 5}
	a2 := (*[2]int)(nums)
	a22 := (*[2]int)(nums[3:])
	fmt.Printf("%v=%v: %v\n", *a2, *a22, *a2 == *a22) // [2 3]=[2 3]: true
}
