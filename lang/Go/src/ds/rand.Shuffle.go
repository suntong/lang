// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//fmt.Println("Hello, 世界")
	//a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	a := []int{1, 2, 3}
	randShuffle(a)
	fmt.Println(a)
	randShuffle(a)
	fmt.Println(a)
	randShuffle(a)
	fmt.Println(a)
	randShuffle(a)
	fmt.Println(a)
	randShuffle(a)
	fmt.Println(a)
	randShuffle(a)
	fmt.Println(a)
	randShuffle(a)
	fmt.Println(a)
}

func randShuffle(a []int) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
}
