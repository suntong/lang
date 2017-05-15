// Example 65 from page 51 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

class B { 
  protected static int bx = 10;
  private static int bz = 10;
}

class C : B {
  private static int cx = 11;
  public class D {
    private static int dx = bx + cx;  // Can access protected bx and private cx
    // private static int dz = bz;    // Cannot access private bz in base class
  }
  static void m() { 
    // int z = D.dx;                  // Cannot access private dx in nested class
  }
}

class MyTest {
  public static void Main(String[] args) {
  }
}
