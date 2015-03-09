////////////////////////////////////////////////////////////////////////////
// Porgram: WordsReverse
// Purpose: Words reverse problem demo
// Authors: Tong Sun (c) 2014, All rights reserved
//
// Ref: http://play.golang.org/p/7oUDPvq8P-
// By: RickyS <fredistic@gmail.com>
//
////////////////////////////////////////////////////////////////////////////

// Style: gofmt -tabs=false -tabwidth=2

/*

The words reverse problem:

Given a string say "I am a human being" the output will reverse words but
nothing else, i.e., it should reverse all alphanumeric within each word but
otherwise preserve everything. Eg: output should be "I ma a namuh gnieb". 

NB, 

- No build-in string reverse function or any form of loop statement except the
single one that iterates through the given string in a single pass should be
used. I.e., the program should finishes in O(N) time instead of O(2N) time.

- Pay attention to what the "reverse words but nothing else" requirement
implies. Design more test data to proof that the requirement is not violated.

*/


/*

Implementation explanation:

The reason to "copy each string into an array of runes" is because of the
following note posted by Carlos Castillo <cookieo9@gmail.com> on Fri, Apr
25, 2014 to golang-nuts@googlegroups.com.

If you want to manipulate individual characters or bytes in a string, either for re-ordering or concatenation, you should probably convert them to slices first. Otherwise most operations you perform that generates a new string (notable exception: sub-string) will allocate memory and copy the relevant data so you will be doing extra work, but more importantly putting more pressure on the memory system / garbage collector.

You can convert any string to a []byte to manipulate individual bytes, or you can convert a string to a []rune to manipulate text in a unicode safe manner.

Here's runes: http://play.golang.org/p/zvXq4OY04K

Strings.Split *does not* allocate new memory for the substrings, nor does it copy data, but Strings.Join does. Therefore there are technically more efficient solutions, but I like the symmetry and simplicity of the "backwords" function. The big benefits of the approach is that:

- It respects unicode characters (manipulating []bytes will likely trash utf-8 strings that have non-ascii characters)
- The reversal code doesn't allocate, it does all its work on the original runes
- You have a nice framework for reducing the number of allocations even further, for example, you could convert the entire string to []rune of the start, and have only two allocation / copy operations in total: one from string to []rune and one back, you just have to re-write the tokenizing function, here's an example in backwords2, and it reuses the reverse function un-altered: http://play.golang.org/p/D8A7Fi0MPv


*/

package main

import (
  "fmt"
  "unicode"
)

// Input test data.  Unicode to ensure not just byte-swapping.
var stringlist = [...]string{
  "I am a human being",
  "I can read 汉字 ...",
  "He said: \"I don't like \t Win8\"",
  "antidisestablishmentarianism",
  "x",
  "",
}

// for each string in the test array, reverse words in it.
func main() {
  for _, str := range stringlist {
    fmt.Printf("\n'%s' =>\n", str)
    rs := []rune(str) // Copy each string into an array of runes.
    fmt.Printf("'%s'.\n", reverseWords(rs))
  }
}

/*

Program output:

'I am a human being' =>
'I ma a namuh gnieb'.

'I can read 汉字 ...' =>
'I nac daer 字汉 ...'.

'He said: "I don't like 	 Win8"' =>
'eH dias: "I nod't ekil 	 8niW"'.

'antidisestablishmentarianism' =>
'msinairatnemhsilbatsesiditna'.

'x' =>
'x'.

'' =>
''.

*/


// reverseWords reverses words in the given string
func reverseWords(rs []rune) (ret string) {
  start := 0
  for start < len(rs) {
    _ret, next := reverseIt(rs, start, false)
    ret += _ret
    start = next
  }
  return
}

// reverseIt returns the reversed word if at the start of the word,
//  else returns the none-word character
func reverseIt(rs []rune, start int, inWord bool) (ret string, next int) {
  // returns the none-word character if at it
  if !inWord && !IsAlnum(rs[start]) {
    //fmt.Printf("] none-word char '%c' at %d\n", rs[start], start)
    ret = string(rs[start])
    next = start + 1
    return
  }

  // == reversing the word
  // -- stop condition
  _next := start + 1
  if _next >= len(rs) || !IsAlnum(rs[_next]) {
    //fmt.Printf("] word end with char '%c' at %d\n", rs[start], start)
    ret = string(rs[start])
    next = _next
    return
  }

  // -- reversed word building
  //fmt.Printf("] word char '%c' at %d\n", rs[start], start)
  _ret, next := reverseIt(rs, start+1, true)
  ret = _ret + string(rs[start])
  //fmt.Printf("]  ==> reversed word '%s' at %d\n", ret, next)
  return
}

// IsAlnum returns true if the rune is a letter or decimal digit
func IsAlnum(r rune) bool {
  return unicode.IsLetter(r) || unicode.IsDigit(r)
}
