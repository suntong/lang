// Example 229 from page 191 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;
using System.Collections.Generic;

class MyTest {
  public static void Main(String[] args) {
    Dictionary<String, int> dict = new Dictionary<String, int>();
    dict.Add("Sweden", 46); dict.Add("Germany", 49);
    dict["Japan"] = 81;                         // New entry, no exception thrown
    Print(dict.Keys);                           // Japan Sweden Germany
    Console.WriteLine(dict.Count);              // 3
    // Console.WriteLine(dict["Greece"]);       // ArgumentException
    // dict.Add("Germany", 49);                 // ArgumentException
    bool b1 = dict.Remove("Greece");            // False (but no exception)
    Console.WriteLine(b1);                      // 
    bool b2 = dict.Remove("Japan");             // True
    Console.WriteLine(b2);                      // 
    Print(dict.Keys);                           // Sweden Germany
    bool b3 = dict.ContainsKey("Germany");      // True
    Console.WriteLine(b3);                      // 
    dict["Sweden"] = 45;                        // No exception
    Console.WriteLine(dict["Sweden"]);          // 45
  }

  public static void Print<T>(ICollection<T> coll) {
    foreach (T x in coll) 
      Console.Write("{0} ", x);
    Console.WriteLine();
  }

  public static void Print<K,V>(IDictionary<K,V> dict) {
    foreach (KeyValuePair<K,V> entry in dict) 
      Console.WriteLine("{0} --> {1}", entry.Key, entry.Value);
    Console.WriteLine();
  }
}
