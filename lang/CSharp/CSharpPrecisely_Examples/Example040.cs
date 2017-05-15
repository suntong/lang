// Example 40 from page 33 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

class MyTest {
  static void ArrayInfo(String name, Array a) {
    Console.Write("{0} has length={1} rank={2} [", name, a.Length, a.Rank);
    for (int i=0, stop=a.Rank; i<stop; i++)
      Console.Write(" {0}", a.GetLength(i));
    Console.WriteLine(" ]");
  }

  public static void Main(String[] args) {
    double[,] r2 = new double[3,2];
    for (int i=0; i<3; i++) 
      for (int j=0; j<2; j++)
        r2[i,j] = i + 0.1 * j;
    double[][] t2 = { new double[] {0.0}, 
                      new double[] {1.0, 1.1}, 
                      new double[] {2.0, 2.1, 2.2}};
    ArrayInfo("r2", r2);                            // length=6 rank=2 [ 3 2 ]
    ArrayInfo("t2", t2);                            // length=3 rank=1 [ 3 ]
    r2.SetValue(10.0, 1, 0);                        // Same as r2[1,0] = 10.0;
    r2.SetValue(21.0, 2, 1);                        // Same as r2[2,1] = 21.0;
    ((double[])t2.GetValue(1)).SetValue(10.0, 0);   // Same as t2[1][0] = 10.0;
    ((double[])t2.GetValue(2)).SetValue(21.0, 1);   // Same as t2[2][1] = 21.0;
    foreach (double d in r2)                        // 0 0.1 10 1.1 2.0 21
      Console.Write(d + " ");           
    Console.WriteLine();
    foreach (double[] row in t2)                    // 0 10 1.1 2 21 2.2
      foreach (double d in row) 
        Console.Write(d + " ");
    Console.WriteLine();
  }
}
