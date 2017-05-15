// Example 150 from page 119 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

// This example should be illegal (by C# Language Specification
// 13.4.1) if the I1 interface were removed from class C's list of
// interfaces, but both MS csc 1.1, 2.0 alpha and 2.0 March 2004 CTP,
// and Mono mcs 0.25, 0.28 and 0.91 accept it.

using System;

interface I1 {
  void M0();
}
interface I2 : I1 {
  new void M0();
  int M1();
}
interface I3 : I1 {
  void M1();
  int P { get; }
  int this[int i] { get; }
}
class C : I1, I2, I3 {
  public void M0() { Console.Write("C.M0 "); }
  void I1.M0() { Console.Write("C:I1.M0 "); }
  void I2.M0() { Console.Write("C:I2.M0 "); }
  int  I2.M1() { Console.Write("C:I2.M1 "); return 1; }
  void I3.M1() { Console.Write("C:I3.M1 "); }
  int I3.P { get { return 11; } }
  int I3.this[int i] { get { return i+((I3)this).P; } }
  // void I3.M0() { }                     // Illegal: M0 not explicitly in I3
}
class D : C { }

class MyTest {
  public static void Main(String[] args) {
    C c = new C();
    // C.M0 C:I1.M0 C:I2.M0 C:I2.M1 C:I3.M1
    c.M0(); ((I1)c).M0(); ((I2)c).M0(); ((I2)c).M1(); ((I3)c).M1();
    Console.WriteLine();
    D d = new D();
    // C.M0 C:I1.M0 C:I2.M0 C:I2.M1 C:I3.M1
    d.M0(); ((I1)d).M0(); ((I2)d).M0(); ((I2)d).M1(); ((I3)d).M1();
    Console.WriteLine();
  }
}
