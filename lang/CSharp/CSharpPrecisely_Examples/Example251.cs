// Example 251 from page 213 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

partial interface I {
  void M1(C.S n);
}

public partial class C {
  public partial struct S {
    public int x;
    public S Decr() { x--; return this; }
  }

  public void M2(S n) {
    Console.WriteLine("n.x={0} ", n.x);
    M1(n);
  }
}
