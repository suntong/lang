// Example 56 from page 45 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

class InitializerExample {
  static double[] ps = new double[6];
  static readonly Random rnd = new Random();

  static InitializerExample() {         // Static constructor
    double sum = 0;
    for (int i=0; i<ps.Length; i++)     // Fill with increasing random numbers
      ps[i] = sum += rnd.NextDouble();  // Random number 0 <= x < 1
    for (int i=0; i<ps.Length; i++)     // Scale so last ps element is 1.0
      ps[i] /= sum;
  }

  static int roll() {
    double p = rnd.NextDouble();
    int i = 0;
    while (p > ps[i])
      i++;
    return i+1;
  }

  public static void Main(String[] args) {
    for (int i=0; i<36; i++)
      Console.Write(roll() + " ");
    Console.WriteLine();
  }
}
