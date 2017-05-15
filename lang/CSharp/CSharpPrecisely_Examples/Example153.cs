// Example 153 from page 123 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

// Quicksort using a delegate to compare elements

using System;

class DelegateQuicksort {
  public static void Main(String[] args) {
    Object[] ia = { 5, 7, 3, 9, 12, 45, 4, 8 };
    Qsort(ia, IntCompare, 0, ia.Length-1);
    foreach (int i in ia)
      Console.Write("{0}   ", i);
    Console.WriteLine();
    String[] sa = { "New York", "Rome", "Dublin", "Riyadh", "Tokyo" };
    Qsort(sa, StringReverseCompare, 0, sa.Length-1);
    foreach (String s in sa)
      Console.Write("{0}   ", s);
    Console.WriteLine();
    String[] sa2 = { "New York", "Rome", "Dublin", "Riyadh", "Tokyo" };
    Qsort(sa2, (v1, v2) => String.Compare((String)v2, (String)v1), 0, sa2.Length-1);
    foreach (String s in sa2)
      Console.Write("{0}   ", s);
    Console.WriteLine();
  }

  // Quicksort: sorts arr[a..b] using delegate cmp to compare elements
  
  private static void Qsort(Object[] arr, DComparer cmp, int a, int b) {
    if (a < b) { 
      int i = a, j = b;
      Object x = arr[(i+j) / 2];             
      do {                              
        while (cmp(arr[i], x) < 0) i++;         // Call delegate cmp
        while (cmp(x, arr[j]) < 0) j--;         // Call delegate cmp
        if (i <= j) {
          Object tmp = arr[i]; arr[i] = arr[j]; arr[j] = tmp;    
          i++; j--;                     
        }                             
      } while (i <= j);                 
      Qsort(arr, cmp, a, j);                 
      Qsort(arr, cmp, i, b);                 
    }                                   
  }

  // The DComparer delegate type
  
  public delegate int DComparer(Object v1, Object v2);

  // Comparison methods for int and String
  
  static int IntCompare(Object v1, Object v2) {
    int i1 = (int)v1, i2 = (int)v2;
    return i1 < i2 ? -1 : i1 > i2 ? +1 : 0;
  }
  
  static int StringReverseCompare(Object v1, Object v2) {
    return String.Compare((String)v2, (String)v1);
  }
}
