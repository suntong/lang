// Example 168 from page 137 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;
using System.Diagnostics;       // For Stopwatch
using System.Threading.Tasks;   // For Parallel

class MyTest {
  public static void Main(String[] args) {
    if (args.Length != 3) 
      Console.WriteLine("Usage: MatrixMultiply <aRows> <aCols> <bCols>\n");
    else {
      int 
        aRows = int.Parse(args[0]), 
        aCols = int.Parse(args[1]), 
        bCols = int.Parse(args[2]);
      double[,] 
        A = RandomMatrix(aRows, aCols), 
        B = RandomMatrix(aCols, bCols),
        R = new double[aRows, bCols];
      {
        Console.WriteLine("Sequential matrix multiplication");
        Stopwatch timer = new Stopwatch();
        timer.Reset();
        timer.Start();
        int count = 10000;
        for (int i=0; i<count; i++) 
          Multiply(R, A, B);
        timer.Stop();
        double time = timer.ElapsedMilliseconds * 1E0;
        Console.WriteLine("{0} ms/multiplication", time);
      }
      {
        Console.WriteLine("Parallel matrix multiplication");
        Stopwatch timer = new Stopwatch();
        timer.Reset();
        timer.Start();
        int count = 10000;
        for (int i=0; i<count; i++) 
          MultiplyParallel(R, A, B);
        timer.Stop();
        double time = timer.ElapsedMilliseconds * 1E0;
        Console.WriteLine("{0} ms/multiplication", time);
      }
    }
  }

  public static void Multiply(double[,] R, double[,] A, double[,] B) {
    int 
      aRows = A.GetLength(0), 
      aCols = A.GetLength(1),
      bRows = B.GetLength(0), 
      bCols = B.GetLength(1),
      rRows = R.GetLength(0), 
      rCols = R.GetLength(1);
    if (aCols==bRows && rRows==aRows && rCols==bCols) {
      for (int r=0; r<rRows; r++) 
        for (int c=0; c<rCols; c++) {
          double sum = 0.0;
          for (int k=0; k<aCols; k++)
            sum += A[r,k]*B[k,c];
          R[r,c] = sum;
        }
    }
  }

  public static void MultiplyParallel(double[,] R, double[,] A, double[,] B) {
    int 
      aRows = A.GetLength(0), 
      aCols = A.GetLength(1),
      bRows = B.GetLength(0), 
      bCols = B.GetLength(1),
      rRows = R.GetLength(0), 
      rCols = R.GetLength(1);
    if (aCols==bRows && rRows==aRows && rCols==bCols) {
      Parallel.For(0, rRows, r => 
        {
          for (int c=0; c<rCols; c++) {
            double sum = 0.0;
            for (int k=0; k<aCols; k++)
              sum += A[r,k]*B[k,c];
            R[r,c] = sum;
          }
        });
    }
  }

  private static readonly Random rnd = new Random(117);

  public static double[,] RandomMatrix(int rows, int cols) { 
    double[,] res = new double[rows, cols];
    for (int r=0; r<rows; r++) 
      for (int c=0; c<cols; c++)
        res[r,c] = rnd.NextDouble();
    return res;
  }

  public static void PrintMatrix(double[,] M) {
    Console.WriteLine();
    for (int i=0; i<M.GetLength(0); i++) {
      for (int j=0; j<M.GetLength(1); j++) 
        Console.Write("{0} ", M[i,j]);
      Console.WriteLine();
    }
  }      
}
