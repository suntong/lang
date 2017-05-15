// Example 167 from page 135 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;
using System.Threading;

class Person  {
  String name, fst, snd;

  public Person(String name, String fst, String snd) { 
    this.name = name; this.fst = fst; this.snd = snd; 
    new Thread(new ThreadStart(Run)).Start();
  }

  public void Run() {
    lock (fst) {
      Console.WriteLine(name + " got " + fst);
      Thread.Sleep(0);		// yield to other threads
      lock (snd) 
        { Console.WriteLine(name + " got " + snd); }
      Console.WriteLine(name + " released " + snd);
    }
    Console.WriteLine(name + " released " + fst);
  }
}

class TestDeadlock {
  public static void Main(String[] args) {
    String left  = "left shoe", right = "right shoe";
    new Person("groucho", left, right);
    new Person("harpo", right, left);
  }
}
