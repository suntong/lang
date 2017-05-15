// Example 9 from page 11 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

class MyTest {
  public static void Main(String[] args) {
    double d = 2.9;
    Console.WriteLine((int)d);                    // ET double-->int; prints 2
    Console.WriteLine((int)(-d));                 // ET double-->int; prints -2
    uint seconds = (uint)(24 * 60 * 60);          // EB int-->uint
    double avgSecPerYear = 365.25 * seconds;      // I  uint-->double
    float f = seconds;                            // IL uint-->float
    long nationalDebt1 = 14349539503882;
    double perSecond = 45138.89;
    decimal perDay =                              // ED double-->decimal
      seconds * (decimal)perSecond;               // I  uint-->decimal
    double nd2 = nationalDebt1 + (double)perDay;  // ER decimal-->double
    long nd3 = (long)nd2;                         // ET double-->long
    float nd4 = (float)nd2;                       // ER double-->float
    Console.WriteLine(nationalDebt1);
    Console.WriteLine(nd2);
    Console.WriteLine(nd3);
    Console.WriteLine(nd4);
  }
}
