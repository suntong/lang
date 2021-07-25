// https://play.golang.org/p/kQUgUgQwkOg
// https://play.golang.org/p/IMV_IAt-VQX
// https://stackoverflow.com/questions/53057237/

package main

import (
	"fmt"
	"time"
)

func main() {
	ii := 0

	go func(d int, ii *int) {
		for true {
			*ii++
			fmt.Println("Hello", *ii)
			time.Sleep(time.Duration(d) * time.Second)
		}
	}(1, &ii)

	// wait for 10 seconds before app finished
	time.Sleep(10 * time.Second)
}
