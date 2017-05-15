// Example 173 from page 139 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;
using System.Threading.Tasks;

class TestGui {
  public static void Main(String[] args) {
    {
      int n = 37;
      Console.WriteLine(SlowFibTimeout1Task(n++).Result);
      Console.WriteLine(SlowFibTimeout1Task(n++).Result);
      Console.WriteLine(SlowFibTimeout1Task(n++).Result);
      Console.WriteLine(SlowFibTimeout1Task(n++).Result);
      Console.WriteLine(SlowFibTimeout1Task(n++).Result);
    }
    {
      int n = 37;
      Console.WriteLine(SlowFibTimeout2Task(n++).Result);
      Console.WriteLine(SlowFibTimeout2Task(n++).Result);
      Console.WriteLine(SlowFibTimeout2Task(n++).Result);
      Console.WriteLine(SlowFibTimeout2Task(n++).Result);
      Console.WriteLine(SlowFibTimeout2Task(n++).Result);
    }
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

  // These two versions of SlowFibTimeout are equivalent:

  public static Task<double> SlowFibTimeout1Task(int n) {
    Task<double> slowFibTask = SlowFibTask(n);
    return TaskEx.WhenAny(slowFibTask, TaskEx.Delay(1000))
                 .ContinueWith<double>((Task<Task> task) => 
                                       task.Result == slowFibTask ? slowFibTask.Result : -1);
  }

  public static async Task<double> SlowFibTimeout2Task(int n) {
    Task<double> slowFibTask = SlowFibTask(n);
    Task completed = await TaskEx.WhenAny(slowFibTask, TaskEx.Delay(1000));
    return completed == slowFibTask ? slowFibTask.Result : -1;
  }
}
