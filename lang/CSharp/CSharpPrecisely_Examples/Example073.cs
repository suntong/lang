// Example 73 from page 63 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

public class Checked {
  public static void Main() {
    char a = char.MaxValue;
    // Run-time overflow of conversion from int to char
    char b = (char)(a + 66);           // b = 'A'
    char c = checked((char)(a + 66));  // Throws OverflowException

    int max = int.MaxValue;
    // Run-time overflow of max+1:
    int j = max+1;                     // j = -2147483648
    int k = checked(max+1);            // Throws OverflowException
    int l = checked(Add(max,1));       // l = -2147483648

    // Compile-time constant overflow
    int m = int.MaxValue+1;            // Compile-time error!
    int n = unchecked(int.MaxValue+1); // n = -2147483648
    int p = checked(int.MaxValue+1);   // Compile-time error!

    Console.WriteLine(b); // 'A'
    Console.WriteLine(c); // 'A'
    Console.WriteLine(j); // -2147483648
    Console.WriteLine(l); // -2147483648
    Console.WriteLine(n); // -2147483648
  }

  static int Add(int i, int j) {
    return i+j;
  }
}
