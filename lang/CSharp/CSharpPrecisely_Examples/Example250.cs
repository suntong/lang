// Example 250 from page 213 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

partial interface I {
  void M2(C.S n);
}

sealed partial class C : I {
  public void M1(S n) {
    if (n.x > 0)
      M2(n.Decr());
  }

  public partial struct S {
    public S(int x) { this.x = x; }
  }

  public static void Main() {
    C c = new C();
    c.M1(new S(5));
  }
}
