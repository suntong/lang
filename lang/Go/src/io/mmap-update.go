// https://brunocalza.me/discovering-and-exploring-mmap-using-go/

package main

import (
	"fmt"
	"os"

	"github.com/edsrzf/mmap-go"
)

func main() {
	f, _ := os.OpenFile("./file", os.O_RDWR, 0644)
	defer f.Close()

	mmap, _ := mmap.Map(f, mmap.RDWR, 0)
	defer mmap.Unmap()
	fmt.Println(string(mmap))

	mmap[0] = 'X'
	mmap[5] = 'U'
	mmap[12] = 'F'
	mmap.Flush()
}

/*

go get -v -u github.com/edsrzf/mmap-go

echo 'MMAP update file' > file

go run mmap-update.go

$ cat file
XMAP Update File

rm file

*/
