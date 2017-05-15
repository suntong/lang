// Example 19 from page 19 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

public class Example019 {
  static int eCount(String s) {
    int ecount = 0;
    for (int i=0; i<s.Length; i++)
      if (s[i] == 'e') 
        ecount++;
    return ecount;
  }

  public static void Main(String[] args) {
    if (args.Length != 1) 
      Console.WriteLine("Usage: Example019 <string>\n");
    else {
      Console.WriteLine("Number of e's is " + eCount(args[0]));
    }
  }
}
