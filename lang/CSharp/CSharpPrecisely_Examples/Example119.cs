// Example 119 from page 97 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

class Example119 {
  public static void Main(String[] args) {
    if (args.Length != 1)
      Console.WriteLine("Usage: Example119 <age>\n");
    else
      Console.WriteLine(AgeGroup(int.Parse(args[0])));
  }

  static String AgeGroup(int age) {
    if (age <= 12)      return "child";
    else if (age <= 19) return "teenager";
    else if (age <= 45) return "young";
    else if (age <= 60) return "middle-age";
    else                return "old";
  }
}
