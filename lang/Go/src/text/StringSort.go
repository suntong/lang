package main

import (
	"fmt"
	"sort"
	"strings"
)

func SortString(w string) string {
	s := strings.Split(w, "\n")
	sort.Strings(s)
	return strings.Join(s, "\n")
}

func SortStringChar(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func main() {
	{
		w1 := "bcad\nabc\nabd\nbbb"
		w2 := SortString(w1)

		fmt.Println(w1)
		fmt.Println()
		fmt.Println(w2)
		fmt.Println()
	}
	{
		w1 := "bcad"
		w2 := SortStringChar(w1)

		fmt.Println(w1)
		fmt.Println(w2)
	}
}

/*

bcad
abc
abd
bbb

abc
abd
bbb
bcad

bcad
abcd

*/
