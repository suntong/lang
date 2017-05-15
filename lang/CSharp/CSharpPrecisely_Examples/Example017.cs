// Example 17 from page 19 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

public class StringEks {
  public static void Main() {
    Console.WriteLine("\u0041BC");             // ABC
    Console.WriteLine(@"\u0041BC");            // \u0041BC
    Console.WriteLine("Say \"Hello\"!");       // Say "Hello"!
    Console.WriteLine(@"Say ""Hello""!");      // Say "Hello"!
    String s1 = @"Line 1
    and Line 2";                               // Newline allowed only in verbatim string
    String s2 = "Line 1\n    and Line 2";
    Console.WriteLine(s1 == s2);               // True   
  }
}
