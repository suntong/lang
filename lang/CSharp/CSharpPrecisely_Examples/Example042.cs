// Example 42 from page 33 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

class MyTest {
  public static void Main(String[] args) {
    String[] a = { "Armonk", "Chicago", "Seattle", "London", "Paris" };
    Array.Reverse(a, 0, 3);
    Array.Reverse(a, 3, 2);
    Array.Reverse(a);
    foreach (String s in a) 
      Console.Write(s + " ");
    Console.WriteLine();
  }
}
