// Example 1 from page 3 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

public class Sum {
  static void Main(String[] args) {
    int sum = 0;
    for (int i=0; i<args.Length; i++)
      sum += int.Parse(args[i]);
    Console.WriteLine("The sum is " + sum);
  }
}
