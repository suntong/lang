// Example 130 from page 103 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

class LoopExample5 {
  public static void Main(String[] args) {
    Console.WriteLine("Thursday is " + WeekDayNo3("Thursday"));
  }

  static int WeekDayNo3(String wday) {
    for (int i=0; i < wdays.Length; i++)
      if (wday.Equals(wdays[i]))
        return i+1;
    return -1;                                  // Here used to mean `not found'
  }

  static readonly String[] wdays =
  { "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday" };
}
