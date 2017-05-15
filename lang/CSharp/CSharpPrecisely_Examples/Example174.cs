// Example 174 from page 141 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;                   // OperationCanceledException
using System.Threading.Tasks;   // Task, Task<T>
using System.Threading;         // CancellationToken, CancellationTokenSource

class MyTest {
  public static void Main(String[] args) {
    Console.WriteLine("\n-- Cancellation with acknowledgement -------------");
    {
      CancellationTokenSource cts = new CancellationTokenSource();
      CancellationToken token = cts.Token;
      Task task = TaskEx.Run(() => ComputeTaskWithAcknowledgement(token), token);
      Thread.Sleep(0);  // Allow task to be scheduled
      Console.WriteLine(task.Status);   // Running
      cts.Cancel();
      Thread.Sleep(0);
      Console.WriteLine(task.Status);   // Canceled
      try {
        task.Wait();    // Throws AggregateException containing TaskCanceledException
      } catch (Exception exn) {
        Console.WriteLine("Caught " + exn);
      }
      Console.WriteLine(task.Status);   // Canceled
    }
    Console.WriteLine("\n-- Cancellation without acknowledgement ----------");
    {
      CancellationTokenSource cts = new CancellationTokenSource();
      CancellationToken token = cts.Token;
      Task task = TaskEx.Run(() => ComputeTaskWithoutAcknowledgement(token), token);
      Thread.Sleep(0);
      Console.WriteLine(task.Status);   // Running
      cts.Cancel();
      Console.WriteLine(task.Status);   // Running
      task.Wait();
      Console.WriteLine(task.Status);   // RanToCompletion
    }
    Console.WriteLine("\n-- Cancellation before Start ---------------------");
    {
      // Cancel before running
      CancellationTokenSource cts = new CancellationTokenSource();
      CancellationToken token = cts.Token;
      Task task = new Task(delegate { }, token);
      Console.WriteLine(task.Status);   // Created
      cts.Cancel();
      Console.WriteLine(task.Status);   // Canceled
      try {
        task.Start();   // Throws InvalidOperationException
      } catch (Exception exn) {
        Console.WriteLine("Caught " + exn);
      }
      Console.WriteLine(task.Status);   // Canceled
    }
    Console.WriteLine("\n-- Completing before cancellation ----------------");
    {
      CancellationTokenSource cts = new CancellationTokenSource();
      CancellationToken token = cts.Token;
      Task task = new Task(delegate { }, token);
      Console.WriteLine(task.Status);   // Created
      task.Start();
      Thread.Sleep(0);  // Allow task to be scheduled
      Console.WriteLine(task.Status);   // RanToCompletion
      cts.Cancel();
      Console.WriteLine(task.Status);   // RanToCompletion
    }
  }

  public static void ComputeTaskWithAcknowledgement(CancellationToken token) {
    for (int i=0; i<100000000; i++) 
      token.ThrowIfCancellationRequested();
  }  

  public static void ComputeTaskWithoutAcknowledgement(CancellationToken token) {
    for (int i=0; i<1000000000; i++) 
      { /* do nothing */ }
  }  
}
