// Example 123 from page 99 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;
using System.Text;

class ForeachArray {
  public static void Main(String[] args) {
    String[] arr =  { "foo", "bar", "", "baz", "" };
    Console.WriteLine(ConcatenateBracketed(arr));
  }

  // Using foreach to iterate over an array

  static String ConcatenateBracketed(String[] arr) {
    StringBuilder sb = new StringBuilder(); 
    foreach (String s in arr) 
      sb.Append(s).Append(s);
    return sb.ToString();
  }
}
