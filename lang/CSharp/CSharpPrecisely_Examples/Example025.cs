// Example 25 from page 23 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

// Prints: 
// |     3D326|    250662|3D326|250662|

using System;

class MyTest {
  public static void Main(String[] args) {
    int i = 250662;
    String s = String.Format("|{0,10:X}|{1,10}|{2:X}|{3}|", i, i, i, i);
    Console.WriteLine(s);
  }
}
