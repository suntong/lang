// Example 79 from page 67 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

class Expressions3 {
  public static void Main(String[] args) {
    Console.WriteLine(Absolute(-12));
    Console.WriteLine(Absolute(12));
    Console.WriteLine(Absolute(0));
  }

  // Returns the absolute value of x (always non-negative)
  static double Absolute(double x)
  { return (x >= 0 ? x : -x); }
}
