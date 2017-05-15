// Example 90 from page 75 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

class B { 
  public static void M() { Console.WriteLine("B.M"); }
  private static void M(int i) { Console.WriteLine("B.M(int)"); }
}

class E {
  static void M() { Console.WriteLine("E.M"); }
  static void M(int i) { Console.WriteLine("E.M(int)"); }

  class C : B {
    public static void Main(String[] args) { 
      M(); 
    }
  }
}

/* 
   1. The call to M() in C will call B.M() because it is inherited by C.

   2. If B.M() is made private it is no longer inherited by C, and the
      call to M() in C will call E.M().

   3. If furthermore B.M(int) is made public, then the call to M() in
      C is no longer legal.  
*/
