// Example 163 from page 131 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;
using System.Threading;

class ThreadDemo {
  private static int i;
  
  public static void Main() {
    Thread u = new Thread(new ThreadStart(Run));
    u.Start();
    Console.WriteLine("Repeatedly press Enter to get the current value of i:");
    for (;;) {
      Console.ReadLine();                        // Wait for keyboard input
      Console.WriteLine(i); 
    }
  }

  private static void Run() { 
    for (;;) {                                    // Forever 
      i++;                                        //   increment i
      Thread.Sleep(0);                            //   yield to other thread
    }
  }
}

