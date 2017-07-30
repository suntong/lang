// https://golang.org/pkg/builtin/

package main

import "fmt"

var (
	v1 = int8(127)
	v2 = uint8(255)
	v3 = int8(-1)
)

func main() {
	fmt.Println(v1, v2, v3)
	// 127 255 -1
	{
		var i int8 = int8(v2)
		var j uint8 = uint8(i)
		fmt.Println(i, j, j == v2)
		// -1 255 true
	}

	{
		var i int64 = -1
		var j uint64

		j = uint64(i)
		var k int = int(j)
		fmt.Println(i, j, k)
		// -1 18446744073709551615 -1
		j = uint64(k)
		fmt.Println(j)
		// 18446744073709551615
	}

}
