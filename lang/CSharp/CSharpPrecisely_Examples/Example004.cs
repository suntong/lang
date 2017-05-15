// Example 4 from page 5 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

class School {
  const int @class = 2004;
  const bool @public = true;
  String @delegate = "J. Smith  ";

  public static int @double(int i) {
    return 2 * @i;
  }
  
  public static void Main(String[] args) {
    School school = new School();
    Console.WriteLine(school.@delegate.Trim() + " " + School.@class);
  }
}
