// Example 162 from page 129 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

// To exercise all paths through the try-catch-finally statement in
// method M, run this program with each of these arguments: 
// 101 102 201 202 301 302 411 412 421 422 431 432
// like this:
//    Example162 101
//    Example162 102
//    etc

using System;

class Example162 {
  public static void Main(String[] args) { 
    if (args.Length != 1) 
      Console.WriteLine("Usage: Example162 <integer>\n");
    else
      Console.WriteLine(M(int.Parse(args[0]))); 
  }

  static String M(int a) {
    try {
      Console.Write("try ... ");
      if (a/100 == 2) return "returned from try";
      if (a/100 == 3) throw new Exception("thrown by try");
      if (a/100 == 4) throw new ApplicationException("thrown by try");
    } catch (ApplicationException) {
      Console.Write("catch ... ");
      if (a/10%10 == 2) return "returned from catch";
      if (a/10%10 == 3) throw new Exception("thrown by catch");
    } finally {
      Console.WriteLine("finally");
      // return "foo"; // Would be illegal
      if (a%10 == 2) throw new Exception("thrown by finally");
    }
    return "terminated normally with " + a;
  }
}
