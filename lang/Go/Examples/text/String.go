////////////////////////////////////////////////////////////////////////////
// Porgram: String
// Purpose: Go string manipulating demo
// Authors: Tong Sun (c) 2013, All rights reserved
////////////////////////////////////////////////////////////////////////////

// Style: gofmt -tabs=false -tabwidth=2

package main

import (
  "fmt"
  "strings"
)

func main() {

  // left(s, n) ~ s[:n]
  // right(s, n) ~ s[len(s)-n:]
  // substr(s, m, n) ~ s[m:n]

  // http://tip.golang.org/pkg/strings/#pkg-examples

  //////
  fmt.Println("\nContains returns true if substr is within s.")

  fmt.Println(strings.Contains("seafood", "foo"))
  fmt.Println(strings.Contains("seafood", "bar"))
  fmt.Println(strings.Contains("seafood", ""))
  fmt.Println(strings.Contains("", ""))
  // true
  // false
  // true
  // true

  //////
  fmt.Println("\nContainsAny returns true if any Unicode code points in chars are within s.")

  fmt.Println(strings.ContainsAny("team", "i"))
  fmt.Println(strings.ContainsAny("failure", "u & i"))
  fmt.Println(strings.ContainsAny("foo", ""))
  fmt.Println(strings.ContainsAny("", ""))
  // false
  // true
  // false
  // false

  //////
  fmt.Println("\nCount counts the number of non-overlapping instances of sep in s.")

  fmt.Println(strings.Count("cheese", "e"))
  fmt.Println(strings.Count("five", "")) // before & after each rune
  // 3
  // 5

  //////
  fmt.Println("\nEqualFold reports whether s and t, interpreted as UTF-8 strings, are equal under Unicode case-folding.")

  fmt.Println(strings.EqualFold("Go", "go"))
  // true

  //////
  fmt.Println("\nFields splits the string s around each instance of one or more consecutive white space characters, as defined by unicode.IsSpace, returning an array of substrings of s or an empty list if s contains only white space.")

  fmt.Printf("Fields are: %q\n", strings.Fields("  foo bar  baz   "))
  // Fields are: ["foo" "bar" "baz"]

  //////
  fmt.Println("\nIndex returns the index of the first instance of sep in s, or -1 if sep is not present in s.")

  fmt.Println(strings.Index("chicken", "ken"))
  fmt.Println(strings.Index("chicken", "dmr"))
  // 4
  // -1

  //////
  fmt.Println("\nIndexRune returns the index of the first instance of the Unicode code point r, or -1 if rune is not present in s.")

  fmt.Println(strings.IndexRune("chicken", 'k'))
  fmt.Println(strings.IndexRune("chicken", 'd'))
  // 4
  // -1

  //////
  fmt.Println("\nJoin concatenates the elements of a to create a single string. The separator string sep is placed between elements in the resulting string.")

  s := []string{"foo", "bar", "baz"}
  fmt.Println(strings.Join(s, ", "))
  // foo, bar, baz

  //////
  fmt.Println("\nLastIndex returns the index of the last instance of sep in s, or -1 if sep is not present in s.")

  fmt.Println(strings.Index("go gopher", "go"))
  fmt.Println(strings.LastIndex("go gopher", "go"))
  fmt.Println(strings.LastIndex("go gopher", "rodent"))
  // 0
  // 3
  // -1

  //////
  fmt.Println("\nMap returns a copy of the string s with all its characters modified according to the mapping function. If mapping returns a negative value, the character is dropped from the string with no replacement.")

  rot13 := func(r rune) rune {
    switch {
    case r >= 'A' && r <= 'Z':
      return 'A' + (r-'A'+13)%26
    case r >= 'a' && r <= 'z':
      return 'a' + (r-'a'+13)%26
    }
    return r
  }
  fmt.Println(strings.Map(rot13, "'Twas brillig and the slithy gopher..."))
  // 'Gjnf oevyyvt naq gur fyvgul tbcure...

  //////
  fmt.Println("\nNewReplacer returns a new Replacer from a list of old, new string pairs. Replacements are performed in order, without overlapping matches.")

  r := strings.NewReplacer("<", "&lt;", ">", "&gt;")
  fmt.Println(r.Replace("This is <b>HTML</b>!"))
  // This is &lt;b&gt;HTML&lt;/b&gt;!

  //////
  fmt.Println("\nRepeat returns a new string consisting of count copies of the string s.")

  fmt.Println("ba" + strings.Repeat("na", 2))
  // banana

  //////
  fmt.Println("\nReplace returns a copy of the string s with the first n non-overlapping instances of old replaced by new. If n < 0, there is no limit on the number of replacements.")

  fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
  fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))
  // oinky oinky oink
  // moo moo moo

  //////
  fmt.Println("\nSplit slices s into all substrings separated by sep and returns a slice of the substrings between those separators. If sep is empty, Split splits after each UTF-8 sequence. It is equivalent to SplitN with a count of -1.")

  fmt.Printf("%q\n", strings.Split("a,b,c", ","))
  fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))
  fmt.Printf("%q\n", strings.Split(" xyz ", ""))
  fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))
  // ["a" "b" "c"]
  // ["" "man " "plan " "canal panama"]
  // [" " "x" "y" "z" " "]
  // [""]

  //////
  fmt.Println("\nSplitAfter slices s into all substrings after each instance of sep and returns a slice of those substrings. If sep is empty, SplitAfter splits after each UTF-8 sequence. It is equivalent to SplitAfterN with a count of -1.")

  fmt.Printf("%q\n", strings.SplitAfter("a,b,c", ","))
  // ["a," "b," "c"]

  //////
  fmt.Println("\nSplitAfterN slices s into substrings after each instance of sep and returns a slice of those substrings. If sep is empty, SplitAfterN splits after each UTF-8 sequence. The count determines the number of substrings to return:")
  // n > 0: at most n substrings; the last substring will be the unsplit remainder.
  // n == 0: the result is nil (zero substrings)
  // n < 0: all substrings

  fmt.Printf("%q\n", strings.SplitAfterN("a,b,c", ",", 2))
  // ["a," "b,c"]

  //////
  fmt.Println("\nSplitN slices s into substrings separated by sep and returns a slice of the substrings between those separators. If sep is empty, SplitN splits after each UTF-8 sequence. The count determines the number of substrings to return:")

  // n > 0: at most n substrings; the last substring will be the unsplit remainder.
  // n == 0: the result is nil (zero substrings)
  // n < 0: all substrings

  fmt.Printf("%q\n", strings.SplitN("a,b,c", ",", 2))
  z := strings.SplitN("a,b,c", ",", 0)
  fmt.Printf("%q (nil = %v)\n", z, z == nil)
  // ["a" "b,c"]
  // [] (nil = true)

  //////
  fmt.Println("\nTitle returns a copy of the string s with all Unicode letters that begin words mapped to their title case.")

  // BUG: The rule Title uses for word boundaries does not handle Unicode punctuation properly.

  fmt.Println(strings.Title("her royal highness"))
  // Her Royal Highness

  //////
  fmt.Println("\nToLower returns a copy of the string s with all Unicode letters mapped to their lower case.")

  fmt.Println(strings.ToLower("Gopher"))
  // gopher

  //////
  fmt.Println("\nToTitle returns a copy of the string s with all Unicode letters mapped to their title case.")

  fmt.Println(strings.ToTitle("loud noises"))
  // LOUD NOISES

  //////
  fmt.Println("\nToUpper returns a copy of the string s with all Unicode letters mapped to their upper case.")

  fmt.Println(strings.ToUpper("Gopher"))
  // GOPHER

  //////
  fmt.Println("\nTrim returns a slice of the string s with all leading and trailing Unicode code points contained in cutset removed.")

  fmt.Printf("[%q]\n", strings.Trim(" !!! Achtung !!! ", "! "))
  // ["Achtung"]

  // TrimPrefix returns s without the provided leading prefix string. If s doesn't start with prefix, s is returned unchanged.

  {
    // var s = "Goodbye,, world!"
    // s = strings.TrimPrefix(s, "Goodbye,")
    // s = strings.TrimPrefix(s, "Howdy,")
    // fmt.Print("Hello" + s)
    // Hello, world!
  }

  //////
  fmt.Println("\nTrimSpace returns a slice of the string s, with all leading and trailing white space removed, as defined by Unicode.")

  fmt.Println(strings.TrimSpace(" \t\n a lone gopher \n\t\r\n"))
  // a lone gopher

  // TrimSuffix returns s without the provided trailing suffix string. If s doesn't end with suffix, s is returned unchanged.

  {
    // var s = "Hello, goodbye, etc!"
    // s = strings.TrimSuffix(s, "goodbye, etc!")
    // s = strings.TrimSuffix(s, "planet")
    // fmt.Print(s, "world!")
    // Hello, world!
  }

}
