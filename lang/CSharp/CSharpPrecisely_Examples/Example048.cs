// Example 48 from page 39 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

class B { 
  public int f;
  public B(int f) { this.f = f; }
  public void M1()         { Console.Write("B.M1 "); }   // Non-virtual instance method
  public virtual void M2() { Console.Write("B.M2 "); }   // Virtual instance method
  public int FVal { get { return f; } }                  // Property
  public int this[int i] { get { return f+i; } }         // Indexer
}

class C : B { 
  public new int f;
  public C(int f) : base(f/2) { 
    this.f = f; 
    Console.WriteLine("{0} {1} {2}", base.f, base.FVal, base[5]); // 11 11 16
    Console.WriteLine("{0} {1} {2}", f, FVal, this[5]);           // 22 22 27
  }
  public new void M1()      { base.M1(); Console.Write("C.M1 "); }
  public override void M2() { base.M2(); Console.Write("C.M2 "); }
  public new int FVal { get { return f; } } 
  public new int this[int i] { get { return f+i; } }
}

class MyTest {
  public static void Main(String[] args) {
    C c = new C(22);
    c.M1(); c.M2();                                      // B.M1 C.M1 B.M2 C.M2
    Console.WriteLine(); 
  }
}
