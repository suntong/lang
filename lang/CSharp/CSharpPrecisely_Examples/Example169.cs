// Example 169 from page 137 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;
using System.Threading.Tasks;           // Parallel
using System.Diagnostics;               // Stopwatch

class MyTest {
  public static void Main(String[] args) {
    { 
      Stopwatch stopwatch = new Stopwatch();
      stopwatch.Reset();
      stopwatch.Start();
      Console.WriteLine("Computing SlowFib(40) * 3 + SlowFib(43) = {0}", SequentialSlowFib());
      stopwatch.Stop();
      Console.WriteLine("Sequential: {0} ms", stopwatch.ElapsedMilliseconds);
    } 
    {
      Stopwatch stopwatch = new Stopwatch();
      stopwatch.Reset();
      stopwatch.Start();
      Console.WriteLine("Computing SlowFib(40) * 3 + SlowFib(43) = {0}", ParallelSlowFib());
      stopwatch.Stop();
      Console.WriteLine("Parallel:   {0} ms", stopwatch.ElapsedMilliseconds);
    }
  }

  public static double SequentialSlowFib() {
    double fib40 = SlowFib(40); 
    double fib43 = SlowFib(43);
    double result = fib40 * 3 + fib43;
    return result;
  }


  public static double ParallelSlowFib() {
    double fib40 = 0.0, fib43 = 0.0;    // Definite assignment rules require initialization
    Parallel.Invoke(delegate { fib40 = SlowFib(40); }, 
                    delegate { fib43 = SlowFib(43); });
    double result = fib40 * 3 + fib43;
    return result;
  }

  public static double SlowFib(int n) {
    if (n < 2) 
      return 1;
    else 
      return SlowFib(n-1) + SlowFib(n-2);
  }  
}
