// Example 18 from page 19 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

public class Example018 {
  public static void Main(String[] args) {
    if (args.Length != 1) 
      Console.WriteLine("Usage: Example018 abc\n");
    else {
      String s1 = "abc", s2 = "ab" + "c", s3 = null;  // Compile-time constants
      String s4 = args[0];                            // Value given at runtime
      // Assume command line argument args[0] is "abc":
      Console.WriteLine(s1==s2);                                // True
      Console.WriteLine((Object)s1==(Object)s2);                // Probably True
      Console.WriteLine(s2==s4);                                // True
      Console.WriteLine((Object)s2==(Object)s4);                // False
      Console.WriteLine("{0} {1} {2}", s3==s1, s3!=s1, s3==s3); // False True True
      Console.WriteLine("{0} {1}", s1!="", s3!="");             // True True
    }
  }
}
