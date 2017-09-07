////////////////////////////////////////////////////////////////////////////
// Porgram: unicode_demo.go
// Purpose: Go unicode demo
// Authors: Tong Sun (c) 2017, All rights reserved
// Credits: http://www.gofragments.net/client/blog/fundamentals/2015/10/30/scanAndCleanRunes/index.html
////////////////////////////////////////////////////////////////////////////

/*

Analyzing a string and processing it.

Example: removing all punctuations and numbers from a not too clean string.

The strings, unicode and bufio packages, with the 'scanner' type and its functions are a way to Go.

*/

package main

import (
	"bufio"
	"fmt"
	"strings"
	"sync/atomic"
	"unicode"
)

func main() {
	// An artificial string with many parasites to be removed.
	input := "(Now) is the 18 18wi##nter of¶ /our [discontent],\n" +
		"Made !!!glorious@ 25summer@ by %this sun& of Y&ork...\n" +
		"And life turns on in spite of ourselves.\n"
		// Cleaning input from any punctuations ('unicode.Punct')
	var listOfPunct []string
	for c := range RunesFromRange(unicode.Punct) {
		// the Latin runes category < 0x600
		if c < 0x600 {
			listOfPunct = append(listOfPunct, string(c))
			// fmt.Printf("%04x %s\n", c, string(c))
		}
	}
	fmt.Printf("-------\nlistOfPunct: %v\n-------\n", listOfPunct)
	for _, punct := range listOfPunct {
		if strings.Contains(input, punct) {
			input = strings.Replace(input, punct, "", -1)
		}
	}
	fmt.Printf("First cleansed 'input' is:\n %q \n-------\n", input)
	fmt.Printf("Then leading and trailing 'numbers' being removed it is:\n")
	scanner := bufio.NewScanner(strings.NewReader(input))

	// Set the 'Split' function of the scanning operation on either
	// 'Words', 'Bytes' or 'Lines'.
	scanner.Split(bufio.ScanWords)
	var Acounter uint32 = 0

	for scanner.Scan() {
		c := atomic.AddUint32(&Acounter, 1)

		// Numbers or not 'numbers'
		f := func(r rune) bool {
			return !unicode.IsLetter(r) &&
				unicode.IsNumber(r) &&
				!unicode.IsMark(r)
		}
		// Removing leading and trailing Unicode satisfying f(), here 'numbers'.
		// signature: func TrimFunc(s string, f func(rune) bool) string
		fmt.Printf("token[%d] is:\t%v\n", c, strings.TrimFunc(scanner.Text(), f))
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("reading input: %s\n", err)
	}
	fmt.Printf("-------\nnumber of entries = %d\n", Acounter)
}

// the unicode package is such a resource to deal with 'runes'
func RunesFromRange(tab *unicode.RangeTable) <-chan rune {
	res := make(chan rune)
	go func() {
		defer close(res)
		for _, r16 := range tab.R16 {
			for c := r16.Lo; c <= r16.Hi; c += r16.Stride {
				res <- rune(c)
			}
		}
		for _, r32 := range tab.R32 {
			for c := r32.Lo; c <= r32.Hi; c += r32.Stride {
				res <- rune(c)
			}
		}
	}()
	return res
}

/* Expected Output:
-------
listOfPunct: [! " # % & ' ( ) * , - . / : ; ? @ [ \ ] _ { } ¡ § « ¶ · » ¿ ; · ՚ ՛ ՜ ՝ ՞ ՟ ։ ֊ ־ ׀ ׃ ׆ ׳ ״ Ή «]
-------
First cleansed 'input' is:
 "Now is the 18 18winter of our discontent\nMade glorious 25summer by this sun of York\nAnd life turns on in spite of ourselves\n"
-------
Then leading and trailing 'numbers' being removed it is:
token[1] is:   Now
token[2] is:   is
token[3] is:   the
token[4] is:
token[5] is:   winter
token[6] is:   of
token[7] is:   our
token[8] is:   discontent
token[9] is:   Made
token[10] is:  glorious
token[11] is:  summer
token[12] is:  by
token[13] is:  this
token[14] is:  sun
token[15] is:  of
token[16] is:  York
token[17] is:  And
token[18] is:  life
token[19] is:  turns
token[20] is:  on
token[21] is:  in
token[22] is:  spite
token[23] is:  of
token[24] is:  ourselves
-------
number of entries = 24
*/
