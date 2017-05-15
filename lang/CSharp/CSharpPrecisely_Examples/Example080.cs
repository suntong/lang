// Example 80 from page 67 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;
using System.Collections.Generic;

class LongComparer : IComparer<long> {
  public int Compare(long v1, long v2) {
    return v1<v2 ? -1 : v1>v2 ? +1 : 0;
  }

  public static void Main(String[] args) {
    // Prints -1 -1 0 0 1 1 1
    LongComparer cmp = new LongComparer();
    Console.WriteLine(cmp.Compare(5L, 7L));
    Console.WriteLine(cmp.Compare(long.MinValue, long.MaxValue));
    Console.WriteLine(cmp.Compare(7L, 7L));
    Console.WriteLine(cmp.Compare(long.MinValue, long.MinValue));
    Console.WriteLine(cmp.Compare(7L, 5L));
    Console.WriteLine(cmp.Compare(0L, long.MinValue));
    Console.WriteLine(cmp.Compare(long.MaxValue, long.MinValue));
  }
}


