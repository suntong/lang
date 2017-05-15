// Example 14 from page 17 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;
using System.Collections.Generic;

class MyTest {

  class Phone {
    public readonly String name;
    public readonly int phone;
    public Phone(String name, int phone) { 
      this.name = name;
      this.phone = phone;
    }
  }

  public static void Main(String[] args) {
    var x = 0.0;                        // Inferred type double
    var b = false;                      // Inferred type bool
    var ps = new List<int>();           // Inferred type List<int>
    ps.Add(2); ps.Add(3); ps.Add(5);
    Console.WriteLine(ps.GetType());
    Console.WriteLine(b);

    var d1 = 34;                      // Inferred type int
    int i1 = d1 * 2;                  // 
    int i2 = (int)d1 * 2;             // Cast (int)d1 succeeds at compile-time
    // bool b1 = d1;                  // Rejected at compile-time
    // d1 = true;                     // Rejected at compile-time
    var p1 = new Phone("Kasper", 5170);
    String s1 = p1.name;              // Field access checked only at compile-time
    // int n1 = p1.age;               // Field access rejected at compile-time
    var p2 = new { name = "Kasper", phone = 5170 };
    String s2 = p2.name;              // Field access checked only at compile-time
    // int n2 = p2.age;               // Field access rejected at compile-time

    foreach (var p in ps) 
      x += p;
    Console.WriteLine(x);
    for (var etor = ps.GetEnumerator(); etor.MoveNext(); ) 
      x += etor.Current;
    Console.WriteLine(x);    
    using (var etor = ps.GetEnumerator()) 
      while (etor.MoveNext())
        x += etor.Current;
    Console.WriteLine(x);
  }
}
