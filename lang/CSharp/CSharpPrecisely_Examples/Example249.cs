// Example 249 from page 211 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

// Namespaces and using

// You must compile the example declaring N1 and N1.N2 to a .dll first.

// Then compile with
//    MS .Net: csc /reference:Example201.dll Example202.cs
//    Mono:    mcs /reference:Example201.dll Example202.cs

using System;
using N1;
using N1.N2;                    // using N2; does not suffice here

class MyTest {
  public static void Main(String[] args) {
    C11 c11;
    C121 c121;
    // C1 c1;                   // Inaccessible: internal to above example
    S13 c13;
    // C122 c122;               // Inaccessible: internal to above example
  }
}
