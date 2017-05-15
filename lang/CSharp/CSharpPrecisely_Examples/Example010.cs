// Example 10 from page 13 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

interface I1 { }
interface I2 : I1 { }
interface J { }
class B : I2 { }
class C : B, J { }
delegate void D(String s);

class MyTest {
  public static void Main(String[] args) {
    Object b1 = new B();            // Implicit B-->Object
    I2 b2 = new B();                // Implicit B-->I2
    B c1 = new C();                 // Implicit C-->B
    I1 b3 = b2;                     // Implicit I2-->B
    I1[] i2a1 = new I2[5];          // Implicit I2[]-->I1[]
    Array inta1 = new int[5];       // Implicit int[]-->Array
    Delegate d1 = new D(Print);     // Implicit D-->Delegate
    C n = null;                     // Implicit null type-->C
    			           
    B b4 = (B)b1;                   // Explicit Object-->B
    C c2 = (C)c1;                   // Explicit B-->C
    J b5 = (J)c1;                   // Explicit C-->J
    B b6 = (B)b2;                   // Explicit I2-->B
    I1 i2 = (I1)b2;                 // Explicit I2-->I1
    I2[] i2a2 = (I2[])i2a1;         // Explicit I1[]-->I2[]
    int[] inta2 = (int[])inta1;     // Explicit Array-->int[]
    D d2 = (D)d1;                   // Explicit Delegate-->D
  }

  public static void Print(String s) {
    Console.WriteLine(s);
    return;
  }
}
