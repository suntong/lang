// Example 38 from page 31 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

public class ArraysExample {
  public static void Main() {
    // Rectangular array creation
    double[,] r1 = { { 0.0, 0.1 }, { 1.0, 1.1 }, { 2.0, 2.1 } };
    double[,] r2 = new double[3,2];
    for (int i=0; i<3; i++) 
      for (int j=0; j<2; j++)
        r2[i,j] = i + 0.1 * j;
    
    // Jagged array creation
    double[] row0 = { 0.0 }, row1 = { 1.0, 1.1 }, row2 = { 2.0, 2.1, 2.2 };
    double[][] t1 = { row0, row1, row2 };
    double[][] t2 = { new double[] {0.0}, 
                      new double[] {1.0, 1.1}, 
                      new double[] {2.0, 2.1, 2.2}};
    double[][] t3 = new double[3][];          // Create first dimension array
    for (int i=0; i<3; i++) {
      t3[i] = new double[i+1];                // Create second dimension arrays
      for (int j=0; j<=i; j++)
        t3[i][j] = i + 0.1 * j;
    }    
    // double[][] t4 = new double[3][3];      // Illegal array creation 
    Print(r1); Print(r2); 
    Print(t1); Print(t2); Print(t3); 
  }

  private static void Print(double[,] arr) {
    for (int i=0; i<arr.GetLength(0); i++) { 
      for (int j=0; j<arr.GetLength(1); j++) 
        Console.Write("{0:F1} ", arr[i,j]);
      Console.WriteLine();
    }
    Console.WriteLine();
  }      

  private static void Print(double[][] arr) {
    for (int i=0; i<arr.Length; i++) { 
      for (int j=0; j<arr[i].Length; j++) 
        Console.Write("{0:F1} ", arr[i][j]);
      Console.WriteLine();
    }
    Console.WriteLine();
  }      
}
