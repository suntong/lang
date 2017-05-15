// Example 29 from page 25 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

// The output is in LaTeX format.
// Used to generate example ex-number-format-custom.

using System;
using System.Threading;
using System.Globalization;
using System.IO;

public class CustomNumberFormat {
  public static void Main() {
    double[] numbers = { 1230.1, 17, 0.15, 0, -26 };
    String[] formats = {"{0:000.0}", "{0:###.#}", "{0:##0.0}", "{0:#0E+0}", "{0:00##;'(neg)';'-'}"};
    int noFormats = formats.Length;

    Thread.CurrentThread.CurrentCulture = new CultureInfo("en-US");

    StreamWriter fs = File.CreateText(Directory.GetCurrentDirectory().ToString() + @"\number-formats-custom.txt");

    fs.WriteLine(@"\begin{center}");
    fs.Write(@"\begin{tabular}{r|");
    for (int i=0; i<noFormats; i++)
      fs.Write("l");
    fs.WriteLine("}");
    fs.WriteLine(@"\hline\hline");
    fs.WriteLine(@"& \multicolumn{" + noFormats + @"}{c}{Format Specifications}\\");
    fs.WriteLine(@"\cline{2-" + (noFormats+1) + @"}\cline{2-" + (noFormats+1) + "}");
    fs.Write(@"Number");

    foreach(String format in formats)
      fs.Write(@" & \verb|" + format + "|");
    fs.WriteLine(@"\\ \hline");

    foreach(double number in numbers) {
      fs.Write(number);
      foreach(String format in formats)
        fs.Write(@" & \verb|" + String.Format(format, number) + "|");
      fs.WriteLine(@"\\");
    }

    fs.WriteLine(@"\hline\hline");
    fs.WriteLine(@"\end{tabular}");
    fs.WriteLine(@"\end{center}");

    fs.Close();
  }
}
