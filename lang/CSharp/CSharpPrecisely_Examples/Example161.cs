// Example 161 from page 129 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

class WeekdayException : ApplicationException {
  public WeekdayException(String wday) : base("Illegal weekday: " + wday) { 
  }
}
