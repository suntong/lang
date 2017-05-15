// Example 124 from page 99 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;
using System.Text;
using System.Collections;

class ForeachExpanded {
  public static void Main(String[] args) {
    String[] arr =  { "foo", "bar", "", "baz", "" };
    Console.WriteLine(ConcatenateBracketed(arr));
  }

  // Using an explicit enumerator instead of foreach

  static String ConcatenateBracketed(String[] arr) {
    StringBuilder sb = new StringBuilder(); 
    IEnumerator enm = arr.GetEnumerator();
    try {
      while (enm.MoveNext()) {
        String s = (String)enm.Current;
	sb.Append(s).Append(s);
      }
    } finally {
      Console.WriteLine("(now in finally block)");
      IDisposable disp = enm as System.IDisposable;
      if (disp != null) 
        disp.Dispose();
    }
    return sb.ToString();
  }
}
