// Example 182 from page 145 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

class MathSpecial {
  public static void Main(String[] args) {
    MathTest();
  }

  static void MathTest() {
    Print("Illegal arguments, NaN results:");
    Print(Math.Sqrt(-1));               // NaN
    Print(Math.Log(-1));                // NaN
    Print(Math.Pow(-1, 2.5));           // NaN
    Print(Math.Acos(1.1));              // NaN
    Print("Infinite results:");
    Print(Math.Log(0));                 // -Infinity
    Print(Math.Pow(0, -1));             // Infinity
    Print(Math.Exp(1000.0));            // Infinity (overflow)
    Print("Infinite arguments:");
    double infinity = Double.PositiveInfinity;
    Print(Math.Sqrt(infinity));         // Infinity
    Print(Math.Log(infinity));          // Infinity
    Print(Math.Exp(-infinity));         // 0
    Print(Math.Pow(infinity, 0.5));     // Infinity
    Print(Math.Pow(0.5, infinity));     // 0
    Print(Math.Pow(0.5, -infinity));    // Infinity
    Print(Math.Pow(2, infinity));       // Infinity
    Print(Math.Pow(2, -infinity));      // 0
    Print("Special cases:");
    Print(Math.Pow(0, 0));              // 1.0
    Print(Math.Pow(infinity, 0));       // 1.0
    Print(Math.Pow(-infinity, 0));      // 1.0
    Print(Math.Pow(-infinity, 0.5));    // Infinity
    Print(Math.Pow(1, infinity));       // NaN
    Print(Math.Pow(1, -infinity));      // NaN
    // For all (x, y) except (0.0, 0.0):
    // sign(Cos(Atan2(y, x))) == sign(x) && sign(Sin(Atan2(y, x))) == sign(y)
    for (double x=-100; x<=100; x+=0.125) {
      for (double y=-100; y<=100; y+=0.125) {
	double r = Math.Atan2(y, x);
	if (!(sign(Math.Cos(r))==sign(x) && sign(Math.Sin(r))==sign(y)))
	  Print("x = " + x + "; y = " + y);
      }
    }
  }
  
  // The built-in Math.Sign method cannot be used because Sin and
  // Cos are inexact

  static int sign(double x) {
    double tolerance = 1E-14;
    if (x < -tolerance) 
      return -1;
    else if (x > +tolerance)
      return +1;
    else 
      return 0;
  }

  static void Print(String d) {
    Console.WriteLine(d);
  }

  static void Print(double d) {
    Console.WriteLine(d);
  }

  static void Print(long d) {
    Console.WriteLine(d);
  }
}
