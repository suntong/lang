// Example 33 from page 27 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;
using System.Text;              // StringBuilder

class Example033 {
  public static void Main(String[] args) {
    if (args.Length != 1)
      Console.WriteLine("Usage: Example033 <length>\n");
    else {
      Console.WriteLine("Timing character replacement in a string:");
      Random rnd = new Random();
      int length = int.Parse(args[0]);
      char[] cbuf = new char[length];
      for (int i=0; i<length; i++)
        cbuf[i] = (char)(65 + rnd.Next(26));
      String s = new String(cbuf);
      for (int i=0; i<10; i++) {
        StringBuilder sb = new StringBuilder(s);
        Timer t = new Timer();
        ReplaceCharString(sb, 'A', "HA");
        Console.Write(t.Check() + " ");
      }
      Console.WriteLine();
    }
  }

  // In-place replacement in a StringBuffer; very inefficient and strange

  static void ReplaceCharString(StringBuilder sb, char c1, String s2) {
    int i = 0;                                  // Inefficient
    while (i < sb.Length) {                     // Inefficient
      if (sb[i] == c1) {                        // Inefficient
        sb.Remove(i, 1);                        // Inefficient
        sb.Insert(i, s2);                       // Inefficient
        i += s2.Length;                         // Inefficient
      } else                                    // Inefficient
        i += 1;                                 // Inefficient
    }                                           // Inefficient
  }

  private class Timer {
    private DateTime start;

    public Timer() {
      start = DateTime.Now;
    }

    public double Check() {
      TimeSpan dur = DateTime.Now - start;
      return dur.TotalSeconds;
    }
  }
}
