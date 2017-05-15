// Example 37 from page 29 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

class MyTest {
  public static void Main(String[] args) {
    Object[] a1 = new String[] { "a", "bc" };   // Legal:   array conversion
    Object[] a2 = new Object[] { 1, 2 };        // Legal:   conversion of 1, 2
    // Object[] a3 = new int[] { 1, 2 };        // Illegal: no array conversion
    // double[] a4 = new int[] { 1, 2 };        // Illegal: no array conversion
    foreach (String s in a1) 
      Console.WriteLine(s);
    foreach (int i in a2) 
      Console.WriteLine(i);
  }
}
