// Example 122 from page 99 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

// Reversing an array of strings

using System;

class MyTest {
  public static void Main(String[] args) {
    Reverse(args);
    Console.WriteLine("Reversed input:");
    Console.WriteLine("--------------------");
    foreach (String s in args) 
      Console.WriteLine(s);
    Console.WriteLine("--------------------");
  }

  public static void Reverse(Object[] arr) {
    for (int s=0, t=arr.Length-1; s<t; s++, t--) {
      Object tmp = arr[s]; arr[s] = arr[t]; arr[t] = tmp;
    }
  }
}
