// Example 188 from page 153 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;
using System.IO;		// FileStream, StreamWriter, Textwriter

public class TextWriterExample {
  public static void Main() {
    TextWriter tw = new StreamWriter(new FileStream("temperature.html", FileMode.Create));
    tw.WriteLine("<table border><tr><th>Fahrenheit<th>Celsius</tr>");
    for (double f=100; f<=400; f+=10) {
      double c = 5 * (f - 32) / 9;
      tw.WriteLine("<tr align=right><td>{0:#0}<td>{1:0.0}", f, c);
    }
    tw.WriteLine("</table>");
    tw.Close();                 // Without this, the output file may be empty
  }
}
