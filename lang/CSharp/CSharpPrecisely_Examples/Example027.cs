// Example 27 from page 23 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

public class StringFormattingAlignment {
  public static void Main() {
    // Fill a 3x5 -matrix with random integers from [0,999]
    Random rnd = new Random();          // Random number generator
    int[,] m = new int[3,5];            // 3x5 matrix
    for (int i=0; i<m.GetLength(0); i++)
      for (int j=0; j<m.GetLength(1); j++)
        m[i,j] = rnd.Next(1000);        // Random integer in range 0..999
    
    // Print matrix
    for (int i=0; i<m.GetLength(0); i++)
      Console.WriteLine("{0,4} {1,4} {2,4} {3,4} {4,4}", m[i,0], m[i,1], m[i,2], m[i,3], m[i,4]);
  }
}
