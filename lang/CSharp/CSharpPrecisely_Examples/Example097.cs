// Example 97 from page 81 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

// Experiments with method modifiers (except for access modifiers)

using System;

abstract class A {
  public static void M1() { Console.Write("A.M1 "); }   
  public void M2() { Console.Write("A.M2 "); }  
  public virtual void M3() { Console.Write("A.M3 "); }  
  public abstract void M4();
}

class B : A { 
  public override void M4() { Console.Write("B.M4 "); } 
}

class C : B {
  public new static void M1() { Console.Write("C.M1 "); }       
  public new void M2() { Console.Write("C.M2 "); }      
  public override void M3() { Console.Write("C.M3 "); } 
}

abstract class D : C { 
  public new abstract void M2();
  public new virtual void M3() { Console.Write("D.M3 "); } 
  public abstract override void M4();
}

class E : D { 
  public sealed override void M2() { Console.Write("E.M2 "); }
  public override void M3() { Console.Write("E.M3 "); }
  public override void M4() { Console.Write("E.M4 "); } 
}

class MyTest {
  public static void Main(String[] args) {
    E ee = new E(); D de = ee; C ce = ee; B be = ee; A ae = ee;
    A ab = new B(); A ac = new C(); 
    A.M1(); B.M1(); C.M1(); D.M1(); E.M1();        // A.M1 A.M1 C.M1 C.M1 C.M1
    Console.WriteLine();
    ae.M2(); be.M2(); ce.M2(); de.M2(); ee.M2();   // A.M2 A.M2 C.M2 E.M2 E.M2
    Console.WriteLine();
    ae.M3(); be.M3(); ce.M3(); de.M3(); ee.M3();   // C.M3 C.M3 C.M3 E.M3 E.M3
    Console.WriteLine();
    ab.M2(); ac.M2(); ae.M2();                     // A.M2 A.M2 A.M2
    Console.WriteLine();
    ab.M3(); ac.M3(); ae.M3();                     // A.M3 C.M3 C.M3
    Console.WriteLine();
    ab.M4(); ac.M4(); ae.M4();                     // B.M4 B.M4 E.M4
    Console.WriteLine();
  }
}
