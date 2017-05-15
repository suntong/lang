// Example 120 from page 97 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

class FindingCountry {
  public static void Main(String[] args) {
    Console.WriteLine("44 is " + FindCountry(44));
  }

  static String FindCountry(int prefix) {
    switch (prefix) {
    default:  return "Unknown";
    case 1:   return "North America";
    case 44:  return "Great Britain";
    case 45:  return "Denmark";
    case 299: return "Greenland";
    case 46:  return "Sweden";
    case 7:   return "Russia";
    case 972: return "Israel";
    }
  }
}
