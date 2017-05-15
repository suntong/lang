// Example 118 from page 97 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

class If1 {
  public static void Main(String[] args) {
    Console.WriteLine(Absolute(-12));
    Console.WriteLine(Absolute(12));
    Console.WriteLine(Absolute(0));
  }

  static double Absolute(double x) {
    if (x >= 0)
      return x;
    else
      return -x;
  }
}
