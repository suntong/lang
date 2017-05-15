// Example 227 from page 189 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;
using System.Collections.Generic;

class MyTest {
  public static void Main(String[] args) {
    List<int> lst = new List<int>();
    lst.Add(7); lst.Add(9); lst.Add(13); lst.Add(7);
    Print(lst);                         // 7 9 13 7
    int i1 = lst[2];                    // 13
    Console.WriteLine(i1);
    int i2 = lst.IndexOf(7);            // 0
    Console.WriteLine(i2);
    int i3 = lst.IndexOf(12);           // -1
    Console.WriteLine(i3);
    lst.Remove(8); Print(lst);          // 7 9 13 7
    lst.Remove(7); Print(lst);          // 9 13 7
    lst.Insert(3, 88); Print(lst);      // 9 13 7 88
    int count = 100000;
    Console.WriteLine("Adding elements at end of list (fast) ...");  
    for (int i=0; i<count; i++) {
      lst.Add(i);
      if (i % 5000 == 0)
        Console.Write("{0} ", i);
    }
    Console.WriteLine();
    lst.Clear();
    Console.WriteLine("Adding elements at head of list (slow) ...");
    for (int i=0; i<count; i++) {
      lst.Insert(0, i);
      if (i % 5000 == 0)
        Console.Write("{0} ", i);
    }
    Console.WriteLine();
  }

  public static void Print<T>(ICollection<T> coll) {
    foreach (T x in coll) 
      Console.Write("{0} ", x);
    Console.WriteLine();
  }
}
