// Example 166 from page 135 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;
using System.Threading;

// In general, a while-loop and not an if-statement is needed around
// the Wait operation, in case there are several producers or
// consumers, and somebody may execute Pulse or PulseAll on the buffer
// object.

class Buffer {
  private int contents;
  private bool empty = true;
  public int Get() {
    lock (this) {
      while (empty)
        Monitor.Wait(this); 
      empty = true;
      Monitor.PulseAll(this);
      return contents;
  } }
  public void Put(int v) {
    lock (this) {
      while (!empty)
        Monitor.Wait(this);
      empty = false; 
      contents = v; 
      Monitor.PulseAll(this);
  } }
}

class TestBuffer {
  static readonly Buffer buf = new Buffer();

  public static void Main(String[] args) {
    new Thread(new ThreadStart(producer)).Start(); 
    new Thread(new ThreadStart(consumer)).Start(); 
  }

  private static void producer() {
    for (int i=1; true; i++) {
      buf.Put(i);
      Util.Pause(10, 100);
    } 
  }

  private static void consumer() {
    for (;;) 
      Console.WriteLine("Consumed " + buf.Get());
  }
}

// Pseudo-random numbers and sleeping threads

class Util {
  private static readonly Random rnd = new Random();
  
  public static void Pause(int length) { 
    Thread.Sleep(length); 
  } 

  public static void Pause(int a, int b) { 
    Pause(Random(a, b));
  }
  
  public static int Random(int a, int b) {
    return rnd.Next(a, b);
  }
}
