// Example 154 from page 123 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

class TestDelegate {
  public static void Main(String[] args) {
    TestDelegate o = new TestDelegate();
    D dlg1 = o.M1, dlg2 = M2, dlg3 = dlg1 + dlg2;
    dlg3 += dlg3;
    int y = 0;
    Console.WriteLine(dlg3(ref y));     // Prints: M1/1 M2/2 M1/3 M2/4 4
    dlg3 -= o.M1;
    Console.WriteLine(dlg3(ref y));     // Prints: M1/5 M2/6 M2/7 7
  }
  
  public delegate int D(ref int x);

  int M1(ref int x) { x++; Console.Write("M1/{0} ", x); return x; }
  static int M2(ref int x) { x++; Console.Write("M2/{0} ", x); return x; }
}
