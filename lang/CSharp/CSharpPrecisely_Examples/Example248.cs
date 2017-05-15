// Example 248 from page 211 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

// Namespaces and using

// Compile with
//    MS .Net:  csc /target:library Example201.cs
//    Mono:     mcs /target:library Example201.cs

using System;

namespace N1 {
  public class C11 { N3.C31 c31; }      // N1 depends on N3
  namespace N2 { 
    public class C121 { } 
  }
}
class C1 { }                            // Default accessibility: internal
namespace N1 {
  public struct S13 { }
}
namespace N1.N2 {
  internal class C122 { }
}

namespace N3 { 
  class C31 { N1.C11 c11; }             // N3 depends on N1
}

class MyTest {
  public static void Main(String[] args) {
    N1.C11 c11;
    N1.N2.C121 c121;
    C1 c1;
    N1.S13 c13;
    N1.N2.C122 c122;
  }
}
