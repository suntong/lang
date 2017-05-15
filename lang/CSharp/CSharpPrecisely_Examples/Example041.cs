// Example 41 from page 33 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

class MyTest {
  static String[] a = { "Armonk", "Chicago", "London", "Paris", "Seattle" };

  static void Search(String c) {
    int i = Array.BinarySearch(a, c);
    if (i >= 0)
      Console.WriteLine("{0} found in position {1}", c, i);
    else
      Console.WriteLine("{0} not found; belongs in position {1}", c, ~i);
  }

  public static void Main(String[] args) {
    Search("London");                   // found in position 2
    Search("Aachen");                   // belongs in position 0
    Search("Copenhagen");               // belongs in position 2
    Search("Washington");               // belongs in position 5
  }
}
