// Example 181 from page 145 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

class MathGaussian {
  public static void Main(String[] args) {
    PrintGaussians(100);
  }

  // From http://www.taygeta.com/random/gaussian.html 2001-09-21:
  // The most basic form of the transformation looks like: 
  //               y1 = sqrt( - 2 ln(x1) ) cos( 2 pi x2 )
  //               y2 = sqrt( - 2 ln(x1) ) sin( 2 pi x2 )

  // We start with two independent random numbers, x1 and x2, which
  // come from a uniform distribution (in the range from 0 to 1). Then
  // apply the above transformations to get two new independent random
  // numbers which have a Gaussian distribution with zero mean and a
  // standard deviation of one.

  static void PrintGaussians(int n) {
    Random rnd = new Random();
    for (int i=0; i<n; i+=2) {
      double x1 = rnd.NextDouble(), x2 = rnd.NextDouble();
      Print(Math.Sqrt(-2 * Math.Log(x1)) * Math.Cos(2 * Math.PI * x2));
      Print(Math.Sqrt(-2 * Math.Log(x1)) * Math.Sin(2 * Math.PI * x2));
    }
  }

  static void Print(double d) {
    Console.WriteLine(d);
  }
}
