// Example 116 from page 95 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

public class VariableDeclExample {
  public static void Main(String[] args) {
    int a;
    const int year = 365, week = 7, weekMax = year / week + 1;
    Console.WriteLine(weekMax);
    int x, y = year, z = 3, ratio = z/y;
    const double pi = 3.141592653589;
    bool found = false;
    var stillLooking = true;
    a = x = y; 
    if (!found || stillLooking) 
      Console.WriteLine(a + x + y + z + ratio + pi);
  }
}
