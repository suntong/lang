// Example 132 from page 103 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

class ContinueForeach {
  public static void Main(String[] args) {
    String[] arr =  { "foo", "", "bar", "baz", "" };
    PrintNonBlank3(arr);
  }

  // Using continue to skip empty strings when printing

  static void PrintNonBlank3(String[] arr) {
    Console.WriteLine("---------");
    foreach (String s in arr) {
      if (s == "")
        continue;
      Console.WriteLine(s);
    }
    Console.WriteLine("---------");
  }
}
