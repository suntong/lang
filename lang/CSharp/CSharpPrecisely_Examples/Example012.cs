// Example 12 from page 15 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

class Scope {
  void M1(int x) {    // declaration of parameter x (#1); shadows x (#5)
    x = 7;            // x #1 in scope; legal, but no effect outside M1
  }                   //
  void M3() {         //
    int x;            // declaration of variable x (#2); shadows x (#5)
    x = 7;            // x #2 in scope
  }                   
  void M4() {         //
    x = 7;            // x #5 in scope
    // int x;         // would be ILLEGAL, giving a new meaning to x
  }                   
  void M5() {         // 
    {                 //
      int x;          // declaration of variable x (#3); shadows x (#5)
      x = 7;          // x #3 in scope
    }                 // 
    {                 //
      int x;          // declaration of variable x (#4); shadows x (#5)
      x = 7;          // x #4 in scope
    }                 // 
  }                   
  public int x;       // declaration of field x (#5)

  public static void Main(String[] args) { 
    Scope s = new Scope();
    s.x = 88;
    s.M1(8);
    s.M3();
    s.M4();
  }
}
