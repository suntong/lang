// Example 21 from page 19 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

public class StringEks {
  public static void Main() {
    // These statements print 35A and A1025 because (+) is left associative:
    Console.WriteLine(10 + 25 + "A");  // Same as (10 + 25) + "A"
    Console.WriteLine("A" + 10 + 25);  // Same as ("A" + 10) + 25
  }
}
