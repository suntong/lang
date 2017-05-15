// Example 224 from page 185 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;
using SC = System.Collections;                 // IComparer
using System.Collections.Generic;              // IComparer<T>, IEqualityComparer<T>

public struct IntPair {
  public readonly int Fst, Snd;
  public IntPair(int fst, int snd) {
    this.Fst = fst; this.Snd = snd;
  }
  public override String ToString() {
    return String.Format("({0},{1})", Fst, Snd);
  }
}

public class IntPairComparer : SC.IComparer, IComparer<IntPair>, IEqualityComparer<IntPair>  {
  public int Compare(Object o1, Object o2) {    // For SC.IComparer
    return Compare((IntPair)o1, (IntPair)o2);
  }
  public int Compare(IntPair v1, IntPair v2) {  // For IComparer<T>
    return v1.Fst<v2.Fst ? -1 : v1.Fst>v2.Fst ? +1
         : v1.Snd<v2.Snd ? -1 : v1.Snd>v2.Snd ? +1 : 0;
  }
  public bool Equals(IntPair v1, IntPair v2) {  // For IEqualityComparer<T>
    return v1.Fst==v2.Fst && v1.Snd==v2.Snd;
  }
  public int GetHashCode(IntPair v) {           // For IEqualityComparer<T>
    return v.Fst ^ v.Snd;
  }
}

class MyTest {
  public static void Main(String[] args) {
    IntPair[] ips = 
      { new IntPair(15, 15), new IntPair(12, 30), new IntPair(15, 30) };
    Dictionary<IntPair, String> dict0 = new Dictionary<IntPair, String>();
    foreach (IntPair ip in ips) 
      dict0.Add(ip, "meeting");
    foreach (KeyValuePair<IntPair, String> entry in dict0)
      Console.WriteLine("{0}: {1}", entry.Key, entry.Value);
    Dictionary<IntPair, String> dict1 = 
      new Dictionary<IntPair, String>(dict0, new IntPairComparer());
  }
}
