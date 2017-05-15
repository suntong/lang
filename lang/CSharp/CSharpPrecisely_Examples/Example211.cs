// Example 211 from page 175 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

// Quicksort using a generic delegate to compare elements

using System;

// The DComparer delegate type

public delegate int DComparer<T>(T v1, T v2);

class DelegateQuicksort {
  public static void Main(String[] args) {
    int[] ia = { 5, 7, 3, 9, 12, 45, 4, 8 };
    DComparer<int> intCmp = IntCompare;
    Qsort<int>(ia, intCmp, 0, ia.Length-1);
    foreach (int i in ia)
      Console.Write("{0}   ", i);
    Console.WriteLine();
    String[] sa = { "New York", "Rome", "Dublin", "Riyadh", "Tokyo" };
    DComparer<String> strCmp = String.Compare;
    Qsort<String>(sa, strCmp, 0, sa.Length-1);
    foreach (String s in sa)
      Console.Write("{0}   ", s);
    Console.WriteLine();
  }

  // Quicksort: sorts arr[a..b] using delegate cmp to compare elements
  
  private static void Qsort<T>(T[] arr, DComparer<T> cmp, int a, int b) {
    if (a < b) { 
      int i = a, j = b;
      T x = arr[(i+j) / 2];             
      do {                              
        while (cmp(arr[i], x) < 0) i++;         // Call delegate cmp
        while (cmp(x, arr[j]) < 0) j--;         // Call delegate cmp
        if (i <= j) {
          T tmp = arr[i]; arr[i] = arr[j]; arr[j] = tmp;    
          i++; j--;                     
        }                             
      } while (i <= j);                 
      Qsort<T>(arr, cmp, a, j);                 
      Qsort<T>(arr, cmp, i, b);                 
    }                                   
  }

  // Type-safe comparison method for int 
  
  static int IntCompare(int i1, int i2) {
    return i1 < i2 ? -1 : i1 > i2 ? +1 : 0;
  }
}
