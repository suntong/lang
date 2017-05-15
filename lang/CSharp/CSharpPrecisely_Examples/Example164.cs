// Example 164 from page 133 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

// 1. Run this program to see that mutual exclusion works: the program 
//    alternately prints a dash (-) and a slash (/) forever.
// 2. Then comment out lock(mutex) as in 
//      /* lock (mutex) */ { 
//    and compile and run the program again.  Now the strict alternation
//    between dash (-) and slash (/) in the output will break.

using System;
using System.Threading;

class Printer {
  static readonly Object mutex = new Object();
  public static void Run() {
    for (;;) {
      lock (mutex) {
        Console.Write("-");
        Util.Pause(100,300);
        Console.Write("/");
      }
      Util.Pause(200);
} } }

class TestPrinter {
  public static void Main(String[] args) { 
    Console.WriteLine("Observe concurrent threads.  Use ctrl-C to stop.\n");
    new Thread(new ThreadStart(Printer.Run)).Start();
    new Thread(new ThreadStart(Printer.Run)).Start();
  }
}

// Pseudo-random numbers and sleeping threads

class Util {
  private static readonly Random rnd = new Random();
  
  public static void Pause(int length) { 
    Thread.Sleep(length); 
  } 

  public static void Pause(int a, int b) { 
    Pause(Random(a, b));
  }
  
  public static int Random(int a, int b) {
    return rnd.Next(a, b);
  }
}
