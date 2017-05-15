// Example 127 from page 101 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

class LoopExample3 {
  public static void Main(String[] args) {
    Console.WriteLine("Infinite loop!  Stop it by pressing ctrl-C\n\n");
    int i=0;
    while (i<10);
      i++;
  }
}
