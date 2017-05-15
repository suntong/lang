// Example 125 from page 101 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

class WdaynoWhile {
  public static void Main(String[] args) {
    Console.WriteLine("Thursday is " + WeekDayNo1("Thursday"));
  }

  static int WeekDayNo1(String wday) {
    int i=0;
    while (i < wdays.Length && wday != wdays[i])
      i++;
    // Now i >= wdays.Length or wday == wdays[i]
    if (i < wdays.Length) return i+1;
    else                  return -1;           // Here used to mean `not found'
  }

  static readonly String[] wdays =
  { "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday" };
}
