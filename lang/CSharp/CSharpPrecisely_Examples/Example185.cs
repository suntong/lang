// Example 185 from page 151 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

// Read text file one line at a time, parse one number from each line,
// and compute the sum of these numbers.  NB: double.Parse respects
// the current culture, so one must use culture en-US or similar to
// parse a number whose decimal point is a period (.).

using System;
using System.IO;		// StreamReader, TextReader
using System.Threading;		// Thread
using System.Globalization;	// CultureInfo

class MyTest {
  public static void Main(String[] args) {
    Thread.CurrentThread.CurrentCulture = new CultureInfo("en-US");
    // = new CultureInfo("fr-FR");	// France
    // = new CultureInfo("de-DE");	// Germany
    // = new CultureInfo("da-DK");	// Denmark
    double sum = 0.0;
    TextReader rd = new StreamReader("foo");
    String line;
    while (null != (line = rd.ReadLine())) 
      sum += double.Parse(line);
    rd.Close();
    Console.WriteLine("The sum is {0}", sum);
  }
}
