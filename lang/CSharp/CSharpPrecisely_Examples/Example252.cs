// Example 252 from page 215 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

// Compile with 
//   
//    csc /d:DEBUG Example252.cs

using System;
using System.Diagnostics;

class Example252 {
  public static void Main(String[] args) {
    if (args.Length != 1) 
      Console.WriteLine("Usage: Example252 <integer>\n");
    else {
      int x = int.Parse(args[0]);
      Console.WriteLine("Integer square root of " + x + " is " + Sqrt(x));
    }
  }

  // Modified for C# from C code on Paul Hsieh's square root page

  static int Sqrt(int x) {  // Algorithm by Borgerding, Hsieh, Ulery
    if (x < 0) 
      throw new ArgumentOutOfRangeException("sqrt: negative argument");
    int temp, y = 0, b = 0x8000, bshft = 15, v = x;;
    do {
      if (v >= (temp = (y<<1)+b << bshft--)) {
        y += b; v -= temp;
      }
    } while ((b >>= 1) > 0);
    Debug.Assert((long)y * y <= x && (long)(y+1)*(y+1) > x);
    return y;
  }
}
