// Example 137 from page 107 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;
using System.IO;		// StreamReader, TextReader

class TryFinally {
  public static void Main(String[] args) {
    double[] xs = ReadRecord("foo");
    for (int i=0; i<xs.Length; i++)
      Console.WriteLine(xs[i]);
  }

  static double[] ReadRecord(String filename) {
    TextReader reader = new StreamReader(filename);
    double[] res = new double[3];
    try {
      res[0] = double.Parse(reader.ReadLine());
      res[1] = double.Parse(reader.ReadLine());
      res[2] = double.Parse(reader.ReadLine());
    } finally {
      reader.Close();
    }
    return res;
  }
}
