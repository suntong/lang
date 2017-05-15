// Example 220 from page 183 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;                      // Console
using System.Collections.Generic;  // IList, IDictionary, List, Dictionary, ...

class TestCollections {
  public static void Main(String[] args) {
    IList<bool> list1 = new List<bool>();
    list1.Add(true); list1.Add(false); list1.Add(true); list1.Add(false); 
    Print(list1);                  // Must print: true false true false
    bool b1 = list1[3];            // false
    Console.WriteLine(b1);
    IDictionary<String, int> dict1 = new Dictionary<String, int>();
    dict1.Add("Sweden", 46); dict1.Add("Germany", 49); dict1.Add("Japan", 81); 
    Print(dict1.Keys);             // May print:  Japan Sweden Germany
    Print(dict1.Values);           // May print:  81 46 49
    int i1 = dict1["Japan"];       // 81
    Console.WriteLine(i1);
    Print(dict1);                  // Print key/value pairs in some order
    IDictionary<String, int> dict2 = new SortedDictionary<String, int>();
    dict2.Add("Sweden", 46); dict2.Add("Germany", 49); dict2.Add("Japan", 81); 
    Print(dict2.Keys);             // Must print: Germany Japan Sweden
    Print(dict2.Values);           // Must print: 49 81 46
    Print(dict2);                  // Print key/value pairs in sorted key order
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
