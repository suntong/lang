// Example 135 from page 105 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

// A finite state machine recognizing the regular expression (a|b)*abb
// from Aho, Sethi, Ullman: Compilers, Principles, Techniques, and
// Tools. Addison-Wes;ey 1986 page 136.

using System;

class Example135 {
  public static void Main(String[] args) {
    if (args.Length != 1) 
      Console.WriteLine("Usage: Example135 <string>\n");
    else 
      Console.WriteLine(Match(args[0]) ? "Success" : "Failure");
  }

  public static bool Match(String str) {
    int stop = str.Length, i = 0;
  state1: 
    if (i==stop) return false;
    switch (str[i++]) {
    case 'a': goto state2;
    case 'b': goto state1;
    default: return false;
    }
  state2: 
    if (i==stop) return false;
    switch (str[i++]) {
    case 'a': goto state2;
    case 'b': goto state3;
    default: return false;
    }
  state3: 
    if (i==stop) return false;
    switch (str[i++]) {
    case 'a': goto state2;
    case 'b': goto state4;
    default: return false;
    }
  state4: 
    if (i==stop) return true;
    switch (str[i++]) {
    case 'a': goto state2;
    case 'b': goto state1;
    default: return false;
    }
  } 
}
