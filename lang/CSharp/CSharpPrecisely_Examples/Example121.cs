// Example 121 from page 99 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

class LoopExample1 {
  public static void Main(String[] args) {
    for (int i=1; i<=4; i++) {                  // Output:
      for (int j=1; j<=i; j++)                  // *      
        Console.Write("*");                     // **     
      Console.WriteLine();                      // ***    
    }                                           // ****   
  }
}
