// Example 134 from page 105 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

class ExceptionExamples {
  public static void Main(String[] args) {
    try {
      Console.WriteLine(args[0] + " is weekday number " + WeekDayNo4(args[0]));
    } catch (WeekdayException x) {
      Console.WriteLine("Weekday problem: " + x);
    } catch (Exception x) {
      Console.WriteLine("Other problem: " + x);
    }
  }

  // Behaves the same as wdayno3, but throws Exception instead of
  // returning bogus weekday number:
  static int WeekDayNo4(String wday) {
    for (int i=0; i < wdays.Length; i++)
      if (wday.Equals(wdays[i]))
        return i+1;
    throw new WeekdayException(wday);
  }

  static readonly String[] wdays =
  { "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday" };
}

class WeekdayException : ApplicationException {
  public WeekdayException(String wday) : base("Illegal weekday: " + wday) {
  }
}
