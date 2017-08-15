// Go read stdin line by line
// https://tutorialedge.net/post/golang/reading-console-input-golang/
// https://siongui.github.io/2016/04/06/go-readlines-from-file-or-string/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

/*

$ seq 12 | xargs -n 3 | sed 's/$/   end/' | go run FileRead_line-by-line.go
1 2 3   end
4 5 6   end
7 8 9   end
10 11 12   end

*/
