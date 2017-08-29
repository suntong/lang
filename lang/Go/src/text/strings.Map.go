// http://golangcookbook.com/chapters/strings/processing/test_without_punctuation_using_map.go

package main

import (
	"fmt"
	"strings"
)

func main() {
	removePunctuation := func(r rune) rune {
		if strings.ContainsRune("._-", r) {
			return ' '
		} else if strings.ContainsRune(",:;", r) {
			return -1
		} else {
			return r
		}
	}

	s := "This, that, and the others like THIS_AND-THAT."
	s = strings.Map(removePunctuation, s)
	words := strings.Fields(s)
	for i, word := range words {
		fmt.Println(i, " => ", word)
	}
}

/*

0  =>  This
1  =>  that
2  =>  and
3  =>  the
4  =>  others
5  =>  like
6  =>  THIS
7  =>  AND
8  =>  THAT

*/
