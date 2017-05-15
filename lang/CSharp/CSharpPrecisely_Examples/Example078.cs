// Example 78 from page 67 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

public class CompoundAssignment {
  static double Multiply(double[] xs) {
    double prod = 1.0;
    for (int i=0; i<xs.Length; i++)
      prod *= xs[i];                        // equivalent to: prod = prod * xs[i]
    return prod;
  }

  public static void Main() { 
    Console.WriteLine(Multiply(new double[] { 7.1, 6.3, 10.0 }));
  }
}
