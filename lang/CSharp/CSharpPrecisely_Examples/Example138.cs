// Example 138 from page 109 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

public class CheckedUncheckedStatements {
  public static void Main() {
    String big = "9999999999";               // 9999999999 > int.MaxValue

    checked {
      Console.WriteLine(int.MaxValue + 1);   // Compile-time error
      Console.WriteLine(int.MinValue - 1);   // Compile-time error
      Console.WriteLine((uint)(0-1));        // Compile-time error
      int i = int.Parse("9999999999");       // Throws OverflowException
    }
    unchecked {
      Console.WriteLine(int.MaxValue + 1);   // -2147483648 (wrap-around)
      Console.WriteLine(int.MinValue - 1);   // 2147483647  (wrap-around)
      Console.WriteLine((uint)(0-1));        // 4294967295  (wrap-around)
      int i = int.Parse("9999999999");       // Throws OverflowException
    }
  }
}


