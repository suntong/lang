// Example 35 from page 29 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

class ArrayCheckdate {
  public static void Main(String[] args) {
    Console.WriteLine("August 31 is legal: " + CheckDate(8, 31));
    Console.WriteLine("April 31 is legal:  " + CheckDate(4, 31));
  }

  static readonly int[] days = { 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31 };
 
  static bool CheckDate(int mth, int day) 
  { return (mth >= 1) && (mth <= 12) && (day >= 1) && (day <= days[mth-1]); }
}
