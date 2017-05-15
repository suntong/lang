// Example 131 from page 103 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;
using SC = System.Collections;

class BreakForeach {
  public static void Main(String[] args) {
    String[] arr =  { "foo", "", "bar", "baz", "" };
    SearchNonBlank1(arr);
    SearchNonBlank2(arr);
  }

  // Using break to exit the loop as soon as an empty string is found

  static void SearchNonBlank1(String[] arr) {
    bool found = false;
    foreach (String s in arr)
      if (s == "") {
        found = true;
        break;
      }
    Console.WriteLine(found);
  }

  // A solution with while instead of foreach and break is more cumbersome.

  // Note that method GetEnumerator on an array type returns a
  // non-generic IEnumerator, and that the cast to String in
  // ((String)enm.Current == "") is necessary to obtain a string
  // comparison; otherwise it will be a reference comparison.

  static void SearchNonBlank2(String[] arr) {
    bool found = false;
    SC.IEnumerator enm = arr.GetEnumerator();
    while (!found && enm.MoveNext())
      found = (String)enm.Current == "";
    Console.WriteLine(found);
  }
}
