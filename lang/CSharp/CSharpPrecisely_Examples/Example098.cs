// Example 98 from page 81 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

class MyTest {
  public static void Main(String[] args) {
    D2 d2 = new D2();
    d2.M2();
  }
}

class D1 { 
  public D1() { M2(); }
  public virtual void M1() { Console.WriteLine("D1.M1 "); }
  public virtual void M2() { Console.Write("D1.M2 "); M1(); }
}
  
class D2 : D1 { 
  int f;
  public D2() { f = 7; }
  public override void M1() { Console.WriteLine("D2.M1:" + f); }
}
