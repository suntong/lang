// Example 208 from page 173 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

// Generic quicksort in object-oriented style

using System;
using System.Collections.Generic;               // IComparable<T>

class GenericObjQuicksort {
  public static void Main(String[] args) {
    MyString[] sa = { new MyString("New York"), new MyString("Rome"), 
                      new MyString("Dublin"), new MyString("Riyadh"), 
                      new MyString("Tokyo") };
    Qsort<MyString>(sa, 0, sa.Length-1);
    foreach (MyString s in sa)
      Console.Write("{0}   ", s.Value);
    Console.WriteLine();
  }

  // Generic object-oriented style quicksort: sorts arr[a..b]
  
  private static void Qsort<T>(T[] arr, int a, int b) 
    where T : IComparable<T> {
    if (a < b) { 
      int i = a, j = b;
      T x = arr[(i+j) / 2];             
      do {                              
        while (arr[i].CompareTo(x) < 0) i++;     
        while (x.CompareTo(arr[j]) < 0) j--;     
        if (i <= j) {
          T tmp = arr[i]; arr[i] = arr[j]; arr[j] = tmp;    
          i++; j--;                     
        }                             
      } while (i <= j);                 
      Qsort<T>(arr, a, j);                 
      Qsort<T>(arr, i, b);                 
    }                                   
  }
}

class MyString : IComparable<MyString> {
  private readonly String s;
  public MyString(String s) {
    this.s = s;
  }
  public int CompareTo(MyString that) {
    return String.Compare(that.Value, s);       // Reverse ordering
  }
  public String Value {
    get { return s; }
  }
}
