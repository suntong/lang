// Example 109 from page 91 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

delegate int D1(int x, ref int y, out int z);
delegate int D2(int x, ref int y);
delegate int D3(int x);
delegate void D4(int x, ref int y);
class Test {
  static D1 d11 = delegate(int x, ref int y, out int z) { z = y++; return x + y; };
  static D2 d21 = delegate(int x, ref int y) { y+=2; return x + y; };
  static D2 d22 = delegate { return 5; };
  public static D2 M(int mx) {
    if (mx < 6) 
      return delegate(int x, ref int y) { y+=2; return x + y; };
    else
      return delegate { return mx; };
  }
  public static void Main(String[] args) {
    D2[] ds = { d21, d22, M(4), M(7), delegate { return 8; } };
    int y = 0;
    foreach (D2 d in ds) 
      Console.WriteLine(d(2, ref y));   // Prints 4 5 6 7 8
    Console.WriteLine(y);               // Prints 4
  }
}
