// Example 139 from page 109 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;
using System.IO;		// StreamReader, TextReader

class TestUsing {
  public static void Main(String[] args) {
    double[] xs = ReadRecord("foo");
    for (int i=0; i<xs.Length; i++)
      Console.WriteLine(xs[i]);
  }

  static double[] ReadRecord(String filename) {
    using (TextReader reader = new StreamReader(filename)) {
      double[] res = new double[3];
      res[0] = double.Parse(reader.ReadLine());
      res[1] = double.Parse(reader.ReadLine());
      res[2] = double.Parse(reader.ReadLine());
      return res;
    }
  }
}
