// Example 31 from page 27 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;
using System.Text;		// StringBuilder

class StringBuilderConcatenate {
  public static void Main(String[] args) {
    StringBuilder res = new StringBuilder();
    for (int i=0; i<args.Length; i++)
      res.Append(args[i]);
    Console.WriteLine(res.ToString());
  }
}
