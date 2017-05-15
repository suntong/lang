// Example 20 from page 19 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

class StringConcatenate {
  public static void Main(String[] args) {
    String res = "";                                        // Inefficient
    for (int i=0; i<args.Length; i++)                       // Inefficient
      res += args[i];                                       // Inefficient
    Console.WriteLine(res);
  }
}
