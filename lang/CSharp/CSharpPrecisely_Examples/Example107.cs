// Example 107 from page 89 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

class B { }
class C : B { 
  private String s;
  public C(String s) { this.s = s; }
  public static explicit operator C(String s) { return new C(s + s); }
}

class MyTest {
  public static void Main(String[] args) {
    int i = 4711;
    long ll = (byte)i + (long)i;        // Simple type conversions
    String s = "ole";
    B b1 = new C("foo"), b2 = new B();
    C c1 = (C)b1;                       // Succeeds, b1 has class C
    C c2 = (C)b2;                       // Fails, b2 has class B
    C c3 = (C)s;                        // User-defined conversion String-->C
    Object o = (Object)s;               // Always succeeds
    C c4 = (C)(String)o;                // Succeeds, Object-->String-->C
    C c5 = (C)o;                        // Fails, no Object-->C conversion
    // Array arr = (Array)s;            // Rejected at compile-time
  }
}
