// Example 95 from page 77 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

class MyTest {
  public static int Max(int a, double b) {
    Console.Write("Max(int, double): ");
    return a > b ? a : (int) b;
  }

  public static int Max(int a, int b, int c) {
    Console.Write("Max(int, int, int): ");
    a = a > b ? a : b;
    return a > c ? a : c;
  }

  public static int Max(int x0, params int[] xr) {
    Console.Write("Max(int, int[]): ");
    foreach (int i in xr) 
      if (i > x0) 
        x0 = i;
    return x0;
  }

  public static void Main(String[] args) {
    Console.WriteLine(Max(2, 1));             // Calls Max(int, int[])
    Console.WriteLine(Max(4));                // Calls Max(int, int[])
    Console.WriteLine(Max(5, 8, 7));          // Calls Max(int, int, int)
    Console.WriteLine(Max(8, 16, 10, 11));    // Calls Max(int, int[])
    int[] xr = { 13, 32, 15 };
    Console.WriteLine(Max(12, xr));           // Calls Max(int, int[])
    // Console.WriteLine(Max(16, ref xr[0])); // Illegal: no ref params
  }
}
