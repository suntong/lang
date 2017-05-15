// Example 207 from page 173 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

// Generic quicksort in functional style (the most efficient one)

using System;
using SC = System.Collections;                  // IComparer
using System.Collections.Generic;               // IComparer<T>

class GenericFunQuicksort {
  public static void Main(String[] args) {
    int[] ia = { 5, 7, 3, 9, 12, 45, 4, 8 };
    Qsort<int>(ia, new IntComparer(), 0, ia.Length-1);
    foreach (int i in ia)
      Console.Write("{0}   ", i);
    Console.WriteLine();
    String[] sa = { "New York", "Rome", "Dublin", "Riyadh", "Tokyo" };
    Qsort<String>(sa, new StringReverseComparer(), 0, sa.Length-1);
    foreach (String s in sa)
      Console.Write("{0}   ", s);
    Console.WriteLine();
  }

  // Generic functional-style quicksort: sorts arr[a..b]
  
  private static void Qsort<T>(T[] arr, IComparer<T> cmp, int a, int b) {
    if (a < b) { 
      int i = a, j = b;
      T x = arr[(i+j) / 2];             
      do {                              
        while (cmp.Compare(arr[i], x) < 0) i++;     
        while (cmp.Compare(x, arr[j]) < 0) j--;     
        if (i <= j) {
          T tmp = arr[i]; arr[i] = arr[j]; arr[j] = tmp;    
          i++; j--;                     
        }                             
      } while (i <= j);                 
      Qsort<T>(arr, cmp, a, j);                 
      Qsort<T>(arr, cmp, i, b);                 
    }                                   
  }
}

// Comparers for int and String

public class IntComparer : SC.IComparer, IComparer<int> {
  public int Compare(Object o1, Object o2) {
    return Compare((int)o1, (int)o2);
  }
  public int Compare(int v1, int v2) {
    return v1 < v2 ? -1 : v1 > v2 ? +1 : 0;
  }
  public bool Equals(int v1, int v2) {
    return v1 == v2; 
  }
  public int GetHashCode(int v) {
    return v; 
  }
}

public class StringReverseComparer : IComparer<string> {
  public int Compare(String v1, String v2) {
    return String.Compare(v2, v1);
  }
  public bool Equals(String v1, String v2) {
    return String.Equals(v2, v1);
  }
  public int GetHashCode(String v) {
    return v.GetHashCode(); 
  }
}
