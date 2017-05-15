// Example 216 from page 177 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

// Fun<R,A> is a function with result type R and argument type A

public delegate R Fun<R,A>(A x);

// A struct of type Option<T> represents a value of value type T that
// may be absent.

public struct Option<T> {
  public readonly bool HasValue;
  private readonly T value;
  // Rejected by compiler, but why?  
  // public static readonly Option<T> None = new Option<T>();

  public Option(T value) { this.HasValue = true; this.value = value; }

  public T Value {
    get { 
      if (HasValue) return value;
      else throw new InvalidOperationException("No value");
    }
  }

  public static implicit operator Option<T>(T value) {
    return new Option<T>(value);
  }

  public static explicit operator T(Option<T> option) {
    return option.Value;
  }
  
  public Option<U> Apply<U>(Fun<U,T> fun) {
    if (HasValue)
      return new Option<U>(fun(value));
    else
      return new Option<U>();
  }

  public override String ToString() {
    return HasValue ? value.ToString() : "[No value]";
  }
}

class MyTest {
  public static Option<double> Sqrt(double x) {
    return x >= 0.0 ? new Option<double>(Math.Sqrt(x)) : new Option<double>();
  }

  public static void Main(String[] args) {
    Option<double> res1 = Sqrt(5.0);
    Option<double> res2 = Sqrt(-5.0);
    Console.WriteLine("res1={0} and res2={1}", res1, res2);
    double res3 = (double)res1;         // Explicit Option<double> --> double
    res2 = 17.0;                        // Implicit double --> Option<double>
    Console.WriteLine("res3={0} and res2={1}", res3, res2);
    res1 = res1.Apply(new Fun<double,double>(Math.Log));
    Console.WriteLine("res1={0}", res1);
  }
}
