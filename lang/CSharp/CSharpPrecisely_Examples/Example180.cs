// Example 180 from page 145 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

class MathFactorial {
  public static void Main(String[] args) {
    for (int i=0; i<=100; i++) 
      Console.WriteLine(i + "! = " + Fact(i));
  }

  static double Fact(int n) {
    double res = 0.0;
    for (int i=1; i<=n; i++) 
      res += Math.Log(i);
    return Math.Exp(res);
  }
}
