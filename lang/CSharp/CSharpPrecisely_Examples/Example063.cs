// Example 63 from page 49 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

class TestInit {
  public static void Main(String[] args) {
    new C();			// Should print: 1 2 3 4 5 6 7
  }
}

class B { 
  public readonly int fb1 = Print(3);
  public B(int k) { Print(5); }
  public readonly int fb2 = Print(4);
  public static int Print(int i) { Console.Write(i + " "); return i; }
}

class C : B {
  public readonly int fc1 = Print(1);
  public C() : this(0) { Print(7); }
  public C(int k) : base(k) { Print(6); }
  public readonly int fc2 = Print(2);
}
