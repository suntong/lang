// Example 113 from page 93 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

class MyTest {
  public static void Main(String[] args) {
    var ds = new [] { 2, 3, 5.0, 7 };                               // Type double[]
    Console.WriteLine(ds.GetType());

    var r1 = new [,] { { 0.0, 0.1 }, { 1.0, 1.1 }, { 2.0, 2.1 } };  // Type double[,]
    Console.WriteLine(r1.GetType());

    var arr = new [] { new { n = 22, r = "XXII" }, new { n = 5, r = "V" } };
    int sum = arr[0].n + arr[1].n;
    Console.WriteLine(sum);
  }
}
