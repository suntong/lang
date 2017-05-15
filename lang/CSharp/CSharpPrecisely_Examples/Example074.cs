// Example 74 from page 63 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

public class Year {
  static bool LeapYear(int y) { 
    return y % 4 == 0 && y % 100 != 0 || y % 400 == 0; 
  }
}
