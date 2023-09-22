package main

import (
	"fmt"
	"time"

	"github.com/teris-io/shortid"
)

var shortIDGenerator *shortid.Shortid

func main() {
	shortIDGenerator = shortid.MustNew(1, shortid.DefaultABC, uint64(time.Now().UnixNano()))

	fmt.Println("Hello World", shortIDGenerator.MustGenerate(),
		shortIDGenerator.MustGenerate(), shortIDGenerator.MustGenerate())

	for i := 1; i <= 10; i++ {
		fmt.Println(shortIDGenerator.MustGenerate())
	}
}
