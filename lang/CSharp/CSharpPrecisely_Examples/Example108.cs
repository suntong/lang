// Example 108 from page 89 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

interface I { }
class B { } 
class C : B, I { }
delegate int D(int i);
struct S : I { }
class G<T> {
  public static void WriteType() {
    Console.WriteLine(typeof(T));
  }
}

class MyTest {
  public static void Main(String[] args) {
    Console.WriteLine(typeof(String));              // System.String
    Console.WriteLine(typeof(int));                 // System.Int32  (int)
    Console.WriteLine(typeof(double));              // System.Double (double)
    Console.WriteLine(typeof(int[]));               // System.Int32[]
    Console.WriteLine(typeof(int[][]));             // System.Int32[][]
    Console.WriteLine(typeof(int[,]));              // System.Int32[,]
    Console.WriteLine(typeof(void));                // System.Void
    Console.WriteLine(typeof(B));                   // B
    Console.WriteLine(typeof(C));                   // C
    Console.WriteLine(typeof(I));                   // I
    Console.WriteLine(typeof(D));                   // D
    Console.WriteLine(typeof(S));                   // S
    Console.WriteLine(typeof(G<int>));              // G[System.Int32]
    Console.WriteLine(typeof(G<String>));           // G[System.String]
    G<int>.WriteType();                             // System.Int32
    Console.WriteLine(typeof(int)==typeof(Int32));  // True
  }
}
