// Example 195 from page 161 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;
using System.IO;        // FileInfo, StreamReader

class MyTest {
  public static void Main(String[] args) {
    FileInfo fi1 = new FileInfo("example3\\Prog.cs.old" ); // Windows, Relative
    Console.WriteLine(fi1.Extension);                      // Extension is ".old"
    FileInfo fi2 = new FileInfo("c:tmp\\foo");             // Windows, Volume+relative
    Console.WriteLine(fi2.Extension);                      // Extension is ""
    FileInfo fi3 = new FileInfo("c:\\tmp\\foo");           // Windows, Volume+absolute
    FileInfo fi4 = new FileInfo("example3/Prog.cs");       // Unix, Relative
    Console.WriteLine(fi4.Name);                           // Prog.cs
    Console.WriteLine(fi4.FullName);                       // C:\tmp\example3\Prog.cs
    FileInfo fi5 = new FileInfo("/etc/passwd");            // Unix, Absolute
    Console.WriteLine("--- Printing contents of {0} ---", fi4.Name);
    StreamReader sr = fi4.OpenText();
    String line;
    while ((line = sr.ReadLine()) != null) 
      Console.WriteLine(line);
    sr.Close();
  }
}
