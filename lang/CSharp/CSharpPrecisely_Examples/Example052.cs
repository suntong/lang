// Example 52 from page 43 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

public class Overloading {
  double M(int i) { return i; }
  bool M(bool b) { return !b; }
  static double M(int x, double y) { return x + y + 1; }
  static double M(double x, double y) { return x + y + 3; }

  public static void Main(String[] args) {
    Console.WriteLine(M(10, 20));              // Prints 31
    Console.WriteLine(M(10, 20.0));            // Prints 31
    Console.WriteLine(M(10.0, 20));            // Prints 33
    Console.WriteLine(M(10.0, 20.0));          // Prints 33
  }
}
