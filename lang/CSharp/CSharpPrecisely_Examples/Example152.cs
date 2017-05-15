// Example 152 from page 121 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

class MyTest {
  public static void Main(String[] args) {
    Console.WriteLine(Color.Red + " " + (uint)Color.Red);
  }
}
public enum Color : uint {
  Red = 0xFF0000, Green = 0x00FF00, Blue = 0x0000FF
}
