// Example 60 from page 47 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

class MyTest {
  public static void Main(String[] args) {
    Console.WriteLine(Convert("765"));
    Console.WriteLine(Convert("765", 10));
    Console.WriteLine(Convert("765", 8));
  }

  public static uint Convert(String s, uint radix = 10) {
    uint result = 0;
    foreach (char ch in s) {
      int d = Convert(ch);
      if (d < 0 || d >= radix)
        throw new ArgumentException("Illegal digit");
      result = result * radix + (uint)d;
    }
    return result;
  }

  private static int Convert(char ch) {
    if ('0' <= ch && ch <= '9')
      return ch - '0';
    else if ('A' <= ch && ch <= 'Z')
      return ch - 'A' + 10;
    else if ('a' <= ch && ch <= 'z')
      return ch - 'a' + 10;
    else 
      return -1;
  }
}
