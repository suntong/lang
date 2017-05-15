// Example 104 from page 85 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

abstract class A { 
  public abstract String Cl { get; }
}

class B : A { 
  public static String Name { get { return "B"; } }
  public String Ty { get { return "B"; } }
  public override String Cl { get { return "B"; } }
}

class C : B { 
  public new static String Name { get { return "C"; } }
  public new String Ty { get { return "C:" + base.Ty; } }
  public override String Cl { get { return "C:" + base.Cl; } }
}

class TestProperty {
  public static void Main(String[] args) {
    B b1 = new B();
    C c2 = new C();
    B b2 = c2;
    Console.WriteLine("B.Name = {0}, C.Name = {1}", B.Name, C.Name);
    Console.WriteLine("b1.Ty = {0}, b2.Ty = {1}, c2.Ty = {2}", b1.Ty, b2.Ty, c2.Ty);
    Console.WriteLine("b1.Cl = {0}, b2.Cl = {1}, c2.Cl = {2}", b1.Cl, b2.Cl, c2.Cl);
  }
}
