// Example 110 from page 91 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;
using System.Threading;

class MyTest {
  public static void Main(String[] args) {
    int v = 0;
    (new Thread(new ThreadStart(delegate { 
      Console.Write(v++);
      Thread.Sleep(0);
    }))).Start();
    (new Thread(new ThreadStart(delegate { 
      Console.Write(v--);
      Thread.Sleep(0);
    }))).Start();
    Console.WriteLine();
    Console.WriteLine("\nv = {0}", v);
  }
}

public delegate void D();
