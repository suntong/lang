// Example 16 from page 17 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;
using System.Collections.Generic;

class MyTest {
  private static readonly Random rnd = new Random();

  class Phone {
    public readonly String name;
    public readonly int phone;
    public Phone(String name, int phone) { 
      this.name = name;
      this.phone = phone;
    }
  }
  
  public static void Main(String[] args) {
    // Basic rules of type dynamic
    {
      dynamic d1 = 34;
      int i1 = d1 * 2;                  // OK: cast (int)(d1*2) at run-time
      int i2 = (int)d1 * 2;             // OK: cast (int)d1 at run-time
      // bool b1 = d1;                  // Compiles OK; cast (bool)d1 throws at run-time
      d1 = true;                        // OK
      bool b2 = d1;                     // OK: cast (bool)d1 succeeds at run-time
      dynamic p1 = new Phone("Kasper", 5170);
      String s1 = p1.name;              // Field access checked at run-time
      // int n1 = p1.age;               // Compiles OK; field access throws at run-time
      dynamic p2 = new { name = "Kasper", phone = 5170 };
      String s2 = p2.name;              // Field access checked at run-time
      // int n2 = p2.age;               // Compiles OK; fields access throws at run-time
    }

    // Dynamic operator resolution; run-time type determines meaning of "+" in Plus2() 
    {
      Console.WriteLine(Plus2(int.MaxValue-1));              // -2147483648, due to int overflow
      Console.WriteLine(Plus2((long)(int.MaxValue-1)));      //  2147483648, no long overflow
      Console.WriteLine(Plus2(11.5));                        // 13.5
      Console.WriteLine(Plus2("Spar"));                      // Spar2
      // Console.WriteLine(Plus2(false));        // Compiles OK; throws RuntimeBinderException   
    }

    // Dynamic receiver; run-time type determines whether to call Length on array or String
    {
      dynamic v;
      if (args.Length==0)      v = new int[] { 2, 3, 5, 7 };
      else                     v = "abc";
      int res = v.Length;                      
      Console.WriteLine(res);
    }

    // Dynamic overload resolution; run-time type of v determines which Process called at (**)
    {
      dynamic v;
      if (args.Length==0)      v = 5;
      else if (args[0] == "1") v = "abc";
      else                     v = (Func<int,int>)(x => x*3);
      dynamic r = Process(v);                                   // (**)
      if (args.Length==0 || args[0] == "1")
        Console.WriteLine(r);
      else
        Console.WriteLine(r(11));
      dynamic s = "abc";
      Console.WriteLine(Process(s).StartsWith("abca"));
     }

      // Run-time type tests
    {
      Console.WriteLine(Square(5));
      Console.WriteLine(Square("abc"));
      Func<int,int> f = x => x*3;
      Console.WriteLine(Square(f)(11));
    }

    {
      // Types dynamic[], List<dynamic>, IEnumerable<dynamic> 
      dynamic[] arr = new dynamic[] { 19, "Electric", (Func<int,int>)(n => n+2), 3.2, false };
      int number = arr[0] * 5;
      String street = arr[1].ToUpper();
      int result = arr[2](number);
      Console.WriteLine(number + " " + street);
      double sum = 0;
      List<dynamic> list = new List<dynamic>(arr);
      IEnumerable<dynamic> xs = list;
      foreach (dynamic x in xs) 
        if (x is int || x is double)
          sum += x;
      Console.WriteLine(sum);                                   // 22.2
    }

    // Dynamic and anonymous object expressions
    {
      dynamic v = new { x = 34, y = false };
      Console.WriteLine(v.x);
    }
  }

  // Run-time type of v determines meaning of "+": int+, long+, double+, String+, ...
  static dynamic Plus2(dynamic v) { return v + 2; }

  // Ordinary overloaded methods
  static int Process(int v) { return v * v; }

  static double Process(double v) { return v * v; }

  static String Process(String v) { return v + v; }

  static Func<int,int> Process(Func<int,int> v) { return x => v(v(x)); }
  
  // Explicit tests on run-time type.  Using "dynamic" rather than
  // "Object" here means that the compiler accepts the (v * v)
  // expression and the application of v to arguments, and that the
  // value returned by Square can be further processed: added to,
  // applied to arguments, and so on.
  static dynamic Square(dynamic v) {
    if (v is int || v is double)
      return v * v;
    else if (v is String)
      return v + v;
    else if (v is Func<int,int>)
      return (Func<int,int>)(x => v(v(x)));
    else
      throw new Exception("Don't know how to square " + v);
  }
}

class C : List<dynamic> { }
