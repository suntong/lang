// Example 30 from page 25 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;
using System.Threading;		// Thread
using System.Globalization;	// CultureInfo

public class Tryformatting {

  public static void Main(String[] args) {
    CultureInfo ci;
    ci = new CultureInfo("en-US");	// USA
    // ci = new CultureInfo("fr-FR");	// France
    // ci = new CultureInfo("de-DE");	// Germany
    // ci = new CultureInfo("da-DK");	// Denmark
    Thread.CurrentThread.CurrentCulture = ci;
    Maketable();
  }

  static void Maketable() {
    DateTime now = DateTime.Now;
    String[] fmts = { "{0:F}", "{0:f}", "{0:G}", "{0:g}", "{0:s}",
		      "{0:u}", "{0:U}", "{0:R}", "{0:D}", "{0:Y}", 
		      "{0:M}", "{0:T}", "{0:t}" };
          
    Console.WriteLine("{0,-12} {1}", "Format code", "Formatted date");
    for (int j=0; j<fmts.Length; j++) {
      Console.Write("{0,-12} ", fmts[j]);
      Console.WriteLine(String.Format(fmts[j], now));
    }
  }
}
