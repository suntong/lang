// Example 77 from page 67 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

class Assignments {
  public static void Main() {
    double d;
    d = 12;               // legal:   implicit conversion from int to double
    byte b;
    b = 252 + 1;          // legal:   252 + 1 is a compile-time constant
    // b = 252 + 5;       // illegal: 252 + 5 is too large
    // b = b + 2;         // illegal: b + 2 has type int
    b = (byte)(b + 2);    // legal:   right-hand side has type byte
    b += 2;               // legal:   equivalent to b = (byte)(b + 2)
    // b += 257;          // illegal: b = 257 would be illegal
    Console.WriteLine("b = {0}", b);
  }
}
