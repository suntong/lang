// Example 32 from page 27 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;
using System.Text;		// StringBuilder

class Example032 {
  public static void Main(String[] args) {
    if (args.Length != 1)
      Console.WriteLine("Usage: Example032 <length>\n");
    else {
      Console.WriteLine("Timing character replacement in a string:");
      Random rnd = new Random();
      int length = int.Parse(args[0]);
      char[] cbuf = new char[length];
      for (int i=0; i<length; i++)
        cbuf[i] = (char)(65 + rnd.Next(26));
      String s = new String(cbuf);
      for (int i=0; i<10; i++) {
        Timer t = new Timer();
        String res = ReplaceCharString(s, 'A', "HA");
        Console.Write(t.Check() + " ");
      }
      Console.WriteLine();
    }
  }

  static String ReplaceCharString(String s, char c1, String s2) {
    StringBuilder res = new StringBuilder();
    for (int i=0; i<s.Length; i++)
      if (s[i] == c1)
        res.Append(s2);
      else
        res.Append(s[i]);
    return res.ToString();
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
