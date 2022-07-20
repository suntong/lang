// Given 2 strings, find if they are anagrams.
// anfield and delfina, return true if anagram.

package main
import "fmt"

func main() {
    fmt.Println(areAnagrams("anfield", "delfina"))
    fmt.Println(areAnagrams("foo", "ofo"))
    fmt.Println(areAnagrams("foo", "ofr"))
    fmt.Println(areAnagrams("foo", "bar"))
}

// O(n) where n means string length
func areAnagrams(s1, s2 string) bool {
    hash := map[rune]int{}
    for _, v := range s1 {
        hash[v]++
    }
    for _, v := range s2 {
        hash[v]--
    }
     for _, v := range hash {
         if v != 0 {
             return false
         }
     }
     return true
}