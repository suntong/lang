// Example 75 from page 65 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

public class Bitwise {
  public static void Main() {
    int a = 0x3;                        // Bit pattern   0011
    int b = 0x5;                        // Bit pattern   0101
    WriteLine4(a);                      // Prints        0011
    WriteLine4(b);                      // Prints        0101
    WriteLine4(~a);                     // Prints        1100
    WriteLine4(~b);                     // Prints        1010
    WriteLine4(a & b);                  // Prints        0001
    WriteLine4(a ^ b);                  // Prints        0110
    WriteLine4(a | b);                  // Prints        0111
    Console.WriteLine(1 << 48);         // Prints           65536
    Console.WriteLine(1L << 48);        // Prints 281474976710656
    Console.WriteLine(1024 >> 40);      // Prints               4
    Console.WriteLine(1024L >> 40);     // Prints               0
    Console.WriteLine(1 << -2);         // Prints      1073741824
  }
  static void WriteLine4(int n) {
    for (int i=3; i>=0; i--)
      Console.Write(n >> i & 1);
    Console.WriteLine();
  }
}
