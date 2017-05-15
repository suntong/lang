// Example 141 from page 111 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;
using System.Collections.Generic;

public class Example141 {
  public static void Main(String[] args) {
    if (args.Length != 1) 
      Console.WriteLine("Usage: Example141 <queencount>\n");
    else {
      int n = int.Parse(args[0]);
      foreach (int[] sol in Queens(n-1, n)) {
	foreach (int r in sol) 
	  Console.Write("{0} ", r);
	Console.WriteLine();
      }
    }
  }

  // A result from the ienumerable produced by Queens(w, n) is an int
  // array whose columns 0..w contain a partial solution to the
  // n-queens problem: w+1 queens have been safely placed in the w+1
  // first columns.  It follows that a result of Queens(n-1, n) is a
  // solution to the n-queens problem.

  public static IEnumerable<int[]> Queens(int w, int n) {
    if (w < 0)
      yield return new int[n];
    else 
      foreach (int[] sol in Queens(w-1, n)) 
        for (int r=1; r<=n; r++) {
          for (int c=0; c<w; c++) 
            if (sol[c] == r || sol[c]+(w-c) == r || sol[c]-(w-c) == r)
              goto fail;
          sol[w] = r;
          yield return sol;
          fail: { }
        }
  }
}
