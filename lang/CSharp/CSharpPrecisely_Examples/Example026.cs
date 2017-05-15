// Example 26 from page 23 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

public class ArraysExample {
  public static void Main() {
    // Roll a die, count frequencies
    Random rnd = new Random();          // Random number generator
    int[] freq = new int[6];            // All initialized to 0
    for (int i=0; i<1000; i++) {
      int die = rnd.Next(1, 7);         // Random integer in range 1..6
      freq[die-1] += 1;
    }
    for (int c=1; c<=6; c++)
      Console.WriteLine("{0} came up {1} times", c, freq[c-1]);
  }
}
