// Example 72 from page 63 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

public class ArithmeticOperators {
public static void Main() {
  int max = 2147483647;                         // = int.MaxValue
  int min = -2147483648;                        // = int.MinValue
  WriteLine(max+1);                             // Prints -2147483648
  WriteLine(min-1);                             // Prints  2147483647
  WriteLine(-min);                              // Prints -2147483648
  Write(   10/3); WriteLine(   10/(-3));        // Prints  3 -3
  Write((-10)/3); WriteLine((-10)/(-3));        // Writes -3  3
  Write(   10%3); WriteLine(   10%(-3));        // Prints  1  1
  Write((-10)%3); WriteLine((-10)%(-3));        // Prints -1 -1
}
static void Write(int i)   { Console.Write(i + " "); }
static void WriteLine(int i) { Console.WriteLine(i); }
}
