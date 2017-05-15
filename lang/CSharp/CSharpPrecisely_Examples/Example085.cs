// Example 85 from page 71 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

interface I1 { }
interface I2 : I1 { }
class B : I2 { }
class C : B { }

class MyTest {
  public static void Main(String[] args) {
    Object n1 = new Exception(), n2 = "foo", n3 = null, n4 = 4711;
    Object n5 = new B(), n6 = new C();
    Object n7 = new C[10];
    Print("n1 is a String:      " + (n1 is String));     // False
    Print("n2 is a String:      " + (n2 is String));     // True
    Print("null is a String:    " + (n3 is String));     // False
    Print("4711 is an int:      " + (n4 is int));        // True
    Print("4711 is a long:      " + (n4 is long));       // False
    Print("4711 is a ValueType: " + (n4 is ValueType));  // True
    Print("n5 is an I1:         " + (n5 is I1));         // True
    Print("n5 is an I2:         " + (n5 is I2));         // True
    Print("n5 is a B:           " + (n5 is B));          // True
    Print("n5 is a C:           " + (n5 is C));          // False
    Print("n6 is a B:           " + (n6 is B));          // True
    Print("n6 is a C:           " + (n6 is C));          // True
    Print("n7 is an Array:      " + (n7 is Array));      // True
    Print("n7 is a B[]:         " + (n7 is B[]));        // True
    Print("n7 is a C[]:         " + (n7 is C[]));        // True
  }

  public static void Print(String s) {
    Console.WriteLine(s);
  }
}
