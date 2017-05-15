// Example 111 from page 93 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;
using System.Collections.Generic;
using System.Linq;

class MyTest {
  public static void Main(String[] args) {
    String[] animals = { "cat", "elephant", "dog", "fox", "squirrel" };
    foreach (String animal in animals.Where(s => s.Length==3))
      Console.WriteLine(animal);
    Func<int,int> f1 = x => x * x; 
    Func<double,double> f2 = x => x * x;
    Console.WriteLine(f1(5));
    Console.WriteLine(f2(0.5));
    Func<int,int,int> f3 = (x, y) => x+y;             // Uncurried, call as f3(11, 22)
    Func<int,Func<int,int>> f4 = x => y => x+y;       // Curried, call as f4(11)(22)
    Console.WriteLine(f3(11, 22));
    Console.WriteLine(f4(11)(22));
    var f5 = f4(11);                                  // The function that adds 11
    Console.WriteLine(f5(22));
    // var bad1 = (int x) => x * x;                   // Illegal, cannot infer type
    // Object bad2 = (int x) => x * x;                // Illegal, not delegate type
    // dynamic bad3 = (int x) => x * x;               // Illegal, not delegate type
    // var bads = new [] { (int x) => x * x };        // Illegal, no best type
    Func<int,double> fib1 = null;
    fib1 = n => n < 2 ? 1 : fib1(n-1) + fib1(n-2);
    for (int i=0; i<39; i++)
      Console.Write(fib1(i) + " ");
    Console.WriteLine();
    Func<int,double> fib2 = Recursive<int,double>(fib => n => n < 2 ? 1 : fib(n-1) + fib(n-2));
    for (int i=0; i<39; i++)
      Console.Write(fib2(i) + " ");
    Console.WriteLine();
    Func<int,double> fib3 = RecursiveMemoize<int,double>(fib => n => n < 2 ? 1 : fib(n-1) + fib(n-2));
    for (int i=0; i<39; i++)
      Console.Write(fib3(i) + " ");
    Console.WriteLine();    
  }

  public static Func<A,R> Recursive<A,R>(Func<Func<A,R>,Func<A,R>> protoF) {
    Func<A,R> f = null;
    return f = protoF(x => f(x));
  }

  public static Func<A,R> RecursiveMemoize<A,R>(Func<Func<A,R>,Func<A,R>> protoF) where A : IEquatable<A> {
    var memoTable = new Dictionary<A,R>();
    Func<A,R> f = null;
    return f = protoF(x => memoTable.ContainsKey(x) ? memoTable[x] : memoTable[x] = f(x));
  }

  public static Func<A,C> Compose<A,B,C>(Func<B,C> f, Func<A,B> g) {
    return x => f(g(x));
  }

  public static Func<A,Func<B,C>> Curry<A,B,C>(Func<A,B,C> f) {
    return x => y => f(x,y);
  }

  public static Func<A,B,C> UnCurry<A,B,C>(Func<A,Func<B,C>> f) {
    return (x, y) => f(x)(y);
  }
}
