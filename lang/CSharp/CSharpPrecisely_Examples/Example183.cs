// Example 183 from page 147 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;
using System.IO;

class BasicIOExample {
  public static void Main() {
    TextReader r = Console.In;
    int count = 0;
    String s = r.ReadLine();
    while (s != null && !s.Equals("")) {
      count++;
      s = r.ReadLine();
    }
    Console.WriteLine("You entered " + count + " nonempty lines");
  }
}
