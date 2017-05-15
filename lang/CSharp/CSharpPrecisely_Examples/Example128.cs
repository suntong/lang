// Example 128 from page 101 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

class LoopExample4 {
  public static void Main(String[] args) {
    Console.WriteLine("Counting sum of eyes until 5 or 6 comes up (10000 dice).");
    int[] wait = new int[1000];
    for (int i=0; i<10000; i++)
      wait[WaitSum()]++;
    Console.WriteLine("sum: frequency");
    for (int w=5; w<20; w++)
      Console.WriteLine(w + ": " + wait[w]);
  }
  
  private static readonly Random rnd = new Random();
  
  // Roll a die and compute sum until five or six comes up
  static int WaitSum() {
    int sum = 0, eyes;
    do {
      eyes = 1 + rnd.Next(6);
      sum += eyes;
    } while (eyes < 5);
    return sum;
  }
}
