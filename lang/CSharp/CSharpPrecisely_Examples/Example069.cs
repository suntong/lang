// Example 69 from page 57 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

// The default argumentless constructor Frac(), which is unavoidable
// in a struct type, creates a Frac value that has d==0, violating the
// desirable invariant d!=0.  Fortunately, computing with such Frac
// values will throw DivideByZeroException, and conversion to double
// will produce a NaN.

using System;

struct Frac : IComparable {
  public readonly long n, d;       // NB: Meaningful only if d!=0

  public Frac(long n, long d) {
    long f = Gcd(n, d);
    this.n = n/f; 
    this.d = d/f; 
  }

  private static long Gcd(long m, long n) {
    while (m != 0) 
      m = n % (n = m); 
    return n;
  }

  public static Frac operator+(Frac r1, Frac r2) {
    return new Frac(r1.n*r2.d+r2.n*r1.d, r1.d*r2.d);
  }

  public static Frac operator*(Frac r1, Frac r2) {
    return new Frac(r1.n*r2.n, r1.d*r2.d);
  }

  // Both (or none) of the operators == and != must be defined:

  public static bool operator==(Frac r1, Frac r2) {
    return r1.n==r2.n && r1.d==r2.d;
  }

  public static bool operator!=(Frac r1, Frac r2) {
    return r1.n!=r2.n || r1.d!=r2.d;
  }

  // The preincrement and postincrement operator:

  public static Frac operator++(Frac r) {
    return r + 1;
  }

  // To implement the IComparable interface:

  public int CompareTo(Object that) {
    return ((double)this).CompareTo((double)(Frac)that);
  }

  // When == and != are defined, compatible methods Equals and
  // GetHashCode must be declared also:

  public override bool Equals(Object that) {
    return that is Frac && this == (Frac)that;
  }

  public override int GetHashCode() {
    return n.GetHashCode() ^ d.GetHashCode();
  }
  
  // Implicit conversion from int to Frac:

  public static implicit operator Frac(int n) {
    return new Frac(n, 1);
  }

  // Implicit conversion from long to Frac:

  public static implicit operator Frac(long n) {
    return new Frac(n, 1);
  }

  // Explicit conversion from Frac to long:

  public static explicit operator long(Frac r) {
    return r.n/r.d;
  }

  // Explicit conversion from Frac to float:

  public static explicit operator float(Frac r) {
    return ((float)r.n)/r.d;
  }

  // One cannot have an implicit conversion from Frac to double and at
  // the same time an implicit conversion from Frac to String; this
  // makes it impossible to decide which overload of WriteLine to use.

  public override String ToString() {
    if (d != 1) 
      return n + "/" + d;
    else
      return n.ToString();
  }

  public bool IsZero {
    get { return n==0 && d!=0; }
  }
}

class TestFrac {
  public static void Main(String[] args) {
    Frac r1 = new Frac(6, 2), r2 = new Frac(5, 2);
    Console.WriteLine("r1={0} and r2={1}", r1, r2);
    Console.WriteLine((double)r2);      // Explicit conversion to double
    r2 = r2 * r2;                       // Overloaded multiplication
    Console.WriteLine("{0} {1} {2} {3} {4}", r2, ++r2, r2, r2++, r2);
    r2 = 0;                             // Implicit conversion from long
    for (int i=1; i<=10; i++) {
      r2 += new Frac(1, i);             // Overloaded += derived from overloaded +
      Console.WriteLine(r2 + " " + (r2 == new Frac(11, 6)));
    }
    Console.WriteLine("r2.IsZero is {0}", r2.IsZero);
    // Console.WriteLine(new Frac() + 1);
    // Console.WriteLine(new Frac() * new Frac(2, 3));
    Frac[] fs = { 5, new Frac(7, 8), 4, 2, new Frac(11, 3) };
    Array.Sort(fs);
    foreach (Frac f in fs) 
      Console.WriteLine(f);
    // Using the user-defined conversions:
    Frac f1 = (byte)5;                         // Implicit int-->Frac
    Frac f2 = 1234567890123L;                  // Implicit long-->Frac
    int i1 = (int)f1;                          // Explicit Frac-->long
    double d2 = (double)f2;                    // Explicit Frac-->float
    Console.WriteLine(f1 + "==" + i1);   
    Console.WriteLine("Note loss of precision:");
    Console.WriteLine(f2 + "==" + d2);   
  }
}
