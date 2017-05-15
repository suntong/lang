// Example 175 from page 141 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;
using System.Threading.Tasks;	// Task, Task<T>

class MyTest {
  public static void Main(String[] args) {
    Console.WriteLine("\n-- Faulting a Task with an exception -------------");
    {
      Task task = new Task(delegate { throw new Exception("died"); });
      Console.WriteLine(task.Status);	// Created
      task.Start();
      Console.WriteLine(task.Status);   // Faulted
      try {
	task.Wait();     // Would throw AggregateException containing Exception("died")
      } catch (Exception exn) {
	Console.WriteLine("Caught " + exn);
      }
      Console.WriteLine(task.Status);   // Faulted
    }
    Console.WriteLine("\n-- Faulting a Task<T> with an exception ----------");
    {
      Task<int> task = new Task<int>(delegate { throw new Exception("died"); });
      Console.WriteLine(task.Status);	// Created
      task.Start();
      Console.WriteLine(task.Status);   // Faulted
      try {
	int res = task.Result; // Throws AggregateException with inner Exception("died")
      } catch (Exception exn) {
	Console.WriteLine("Caught " + exn);
      }
      Console.WriteLine(task.Status);   // Faulted
    }
    Console.WriteLine("\n--------------------------------------------------");
  }
}
