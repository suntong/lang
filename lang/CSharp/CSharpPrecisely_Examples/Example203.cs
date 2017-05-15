// Example 203 from page 169 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;
using System.Collections.Generic;

// A constraint may involve type parameters 
// A type may have multiple constraints 

struct ComparablePair<T,U> : IComparable<ComparablePair<T,U>>
  where T : IComparable<T> 
  where U : IComparable<U> {
  public readonly T Fst;
  public readonly U Snd;
  
  public ComparablePair(T fst, U snd) {
    Fst = fst; Snd = snd;
  }
  
  // Lexicographic ordering
  public int CompareTo(ComparablePair<T,U> that) {
    int firstCmp = this.Fst.CompareTo(that.Fst);
    return firstCmp != 0 ? firstCmp : this.Snd.CompareTo(that.Snd);
  }

  public override String ToString() {
    return "(" + Fst + ", " + Snd + ")";
  }
}

// Sorting soccer world champions by country and year

class MyTest {
  public static void Main(String[] args) {
    List<ComparablePair<String,int>> lst 
      = new List<ComparablePair<String,int>>();
    lst.Add(new ComparablePair<String,int>("Brazil", 2002));
    lst.Add(new ComparablePair<String,int>("Italy", 1982));
    lst.Add(new ComparablePair<String,int>("Argentina", 1978 ));
    lst.Add(new ComparablePair<String,int>("Argentina", 1986 ));
    lst.Add(new ComparablePair<String,int>("Germany", 1990));
    lst.Add(new ComparablePair<String,int>("Brazil", 1994));
    lst.Add(new ComparablePair<String,int>("France", 1998));
    lst.Sort();
    foreach (ComparablePair<String,int> pair in lst) 
      Console.WriteLine(pair);
  }
}
