// Example 13 from page 15 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

class MyTest {
  public static void Main(String[] args) {
    int x, y, z;
    if (args.Length == 0) 
      x = y = 10;
    else 
      x = args.Length;
    Console.WriteLine(x);           // x definitely assigned, y and z not (#1)
    y = x;
    for (int i=0; i<y; i++)         // x and y definitely assigned, z not (#2)
      z = i;
    // Console.WriteLine(z);        // z still not definitely assigned!   (#3)
  }
}
