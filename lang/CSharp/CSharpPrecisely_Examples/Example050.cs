// Example 50 from page 41 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

class B {                        // One instance field nf, one static field sf
  public int nf; 
  public static int sf; 
  public B(int i) { nf = i; sf = i+1; } 
}
class C : B {                    // Two instance fields nf, one static field sf
  new public int nf; 
  new public static int sf; 
  public C(int i) : base(i+20) { nf = i; sf = i+2; } 
}
class D : C {                    // Three instance fields nf
  new public int nf; 
  public D(int i) : base(i+40) { nf = i; sf = i+4; } 
}

class FieldAccessExample {
  public static void Main(String[] args) {
    C c1 = new C(100);                  // c1 has type C; object has class C
    B b1 = c1;                          // b1 has type B; object has class C
    Print(C.sf,  B.sf);                 // Prints 102 121 
    Print(c1.nf, b1.nf);                // Prints 100 120
    C c2 = new C(200);                  // c2 has type C; object has class C
    B b2 = c2;                          // b2 has type B; object has class C
    Print(c2.nf, b2.nf);                // Prints 200 220
    Print(c1.nf, b1.nf);                // Prints 100 120
    D d3 = new D(300);                  // d3 has type D; object has class D
    C c3 = d3;                          // c3 has type C; object has class D
    B b3 = d3;                          // b3 has type B; object has class D
    Print(D.sf,  C.sf,  B.sf);          // Prints 304 304 361
    Print(d3.nf, c3.nf, b3.nf);         // Prints 300 340 360
  }

  static void Print(int x, int y) { Console.WriteLine(x+" "+y); }
  static void Print(int x, int y, int z) { Console.WriteLine(x+" "+y+" "+z); }
}
