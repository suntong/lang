// Example 171 from page 139 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;
using System.Threading.Tasks;		// Task, Task<T>

class MyTest {
  public static void Main(String[] args) {
    const int n = 43;
    Console.Write("Computing SlowFib({0}) = ", n); 
    double result = SlowFib(n);
    Console.WriteLine(result);
    Console.Write("Computing SlowFibTask({0}) = ", n); 
    Task<double> task = SlowFibTask(n);         // Returns a Running task
    Console.Write("[task is running] ");
    Console.WriteLine(task.Result);             // Blocks until task completes
  }

  public static double SlowFib(int n) {
    if (n < 2) 
      return 1;
    else 
      return SlowFib(n-1) + SlowFib(n-2);
  }  

  public static Task<double> SlowFibTask(int n) {
    return TaskEx.Run(() => SlowFib(n));
  }  
}
