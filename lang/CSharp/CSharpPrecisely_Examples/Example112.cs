// Example 112 from page 93 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

class MyTest {
  public static void Main(String[] args) {
    double z = 3.14;
    var p1 = new { x = 13, y = "foo" };         // Fields: x, y
    var p2 = new { x = 42, p1.y, z };           // Fields: x, y, z
    var p3 = new { };                           // Fields: none
    int sum = p1.x + p2.x + (int)p2.z + p2.y.Length;
    Console.WriteLine(sum);
    Console.WriteLine(p1);    
    Console.WriteLine(p2);    
    Console.WriteLine(p3); 
  }
}
