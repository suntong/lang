// Example 11 from page 13 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

interface I { void Print(); }
struct S : I {
  public int i;
  public S(int i) { this.i = i; }
  public void Print() { Console.WriteLine(i); }
}

class MyTest {
  public static void Main(String[] args) {
    int i = 7;                          
    Object o = i;                       // Implicit boxing int-->Object
    int j = 5 + (int)o;                 // Explicit unboxing Object-->int
    Console.WriteLine("o {0}", o);                       // 12
    Console.WriteLine("o is int:    {0}", o is int);     // True
    Console.WriteLine("o is double: {0}", o is double);  // False
    S s1 = new S(11);
    I s2 = s1;                          // Implicit boxing S-->I
    s1.i = 22;                          
    s1.Print();                         // 22
    s2.Print();                         // 11
    S s3 = (S)s2;                       // Explicit unboxing I-->S
    s3.Print();                         // 11
  }
}

