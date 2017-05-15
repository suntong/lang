// Example 187 from page 153 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;
using System.IO;                // StreamWriter, Textwriter

public class TextWriterExample {
  public static void Main() {
    TextWriter tw = new StreamWriter("dice.txt");
    Random rnd = new Random();
    for (int i=1; i<=1000; i++) {
      int die = (int)(1 + 6 * rnd.NextDouble());
      tw.Write(die); tw.Write(' ');
      if (i % 20 == 0) tw.WriteLine();
    }
    tw.Close();                 // Without this, the output file may be empty
  }
}
