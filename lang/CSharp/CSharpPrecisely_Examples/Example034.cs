// Example 34 from page 29 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

public class ArraysExample {
  public static void Main() {
    // Roll a die, count frequencies
    Random rnd = new Random();                  // Random number generator
    int[] freq = new int[6];                    // All elements initialized to 0
    for (int i=0; i<1000; i++) {
      int die = rnd.Next(1, 7);
      freq[die-1] += 1;
    }
    for (int c=1; c<=6; c++)
      Console.WriteLine(c + " came up " + freq[c-1] + " times");
  }
}
