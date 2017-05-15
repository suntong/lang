// Example 28 from page 25 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

// The output is in LaTeX format.
// Used to generate example ex-number-format-table

using System;
using System.Threading;
using System.Globalization;
using System.IO;

public class DateFormatting {
  public static void Main() {
    String[] formats = {"{0:D4}", "{0,7}", "{0:F0}", "{0:F2}", "{0,8:F3}", "{0:E4}", "{0,9:C}"};
    int[] numbers1 = { 0, 1, 145, -1 };
    double[] numbers2 =  { 2.5, -1.5, 330.8, 1234.516 };

    Thread.CurrentThread.CurrentCulture = new CultureInfo("en-US");

    StreamWriter fs = File.CreateText(Directory.GetCurrentDirectory().ToString() + @"\number-formats.txt");

    fs.WriteLine(@"\begin{center}");
    fs.WriteLine(@"\begin{tabular}{r|lllllll}");
    fs.WriteLine(@"\hline\hline");
    fs.WriteLine(@"& \multicolumn{7}{c}{Format Specifications}\\");
    fs.WriteLine(@"\cline{2-8}\cline{2-8}");
    fs.Write(@"Number");

    foreach(String format in formats)
      fs.Write(@" & \verb|" + format + "|");
    fs.WriteLine(@"\\ \hline");

    foreach(int number in numbers1) {
      fs.Write(number);
      foreach(String format in formats)
        fs.Write(@" & \verb|" + String.Format(format, number) + "|");
      fs.WriteLine(@"\\");
    }

    foreach(double number in numbers2) {
      fs.Write(number);
      foreach(String format in formats) {
        try { fs.Write(@" & \verb|" + String.Format(format, number) + "|"); }
        catch (FormatException) { fs.Write(" &"); }
      }
      fs.WriteLine(@"\\");
    }

    fs.WriteLine(@"\hline\hline");
    fs.WriteLine(@"\end{tabular}");
    fs.WriteLine(@"\end{center}");

    fs.Close();
  }
}
