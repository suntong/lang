// Example 68 from page 55 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

// The struct type of integer sequences

// Overloaded operators and indexers, enumerators

using System;                           // For Console, String
using System.Text;                      // For StringBuilder
using SC = System.Collections;          // For IEnumerator, IEnumerable
using System.Collections.Generic;       // For IEnumerator<T>

struct Seq : ISeq {
  private readonly int b, k, n;                      // Sequence b+k*[0..n-1]

  // Default constructor Seq() creates an empty sequence with n=0 

  public Seq(int m, int n) : this(m, 1, n-m+1) { }   // Sequence [m..n]
  
  public Seq(int b, int k, int n) {
    this.b = b; this.k = k; this.n = n;
  }

  // Add b to sequence
  public static Seq operator +(int b, Seq seq) {
    return new Seq(seq.b+b, seq.k, seq.n);
  }

  // Add b to sequence
  public static Seq operator +(Seq seq, int b) {
    return new Seq(seq.b+b, seq.k, seq.n);
  }
  
  // Multiply all members of the sequence by k
  public static Seq operator *(int k, Seq seq) {
    return new Seq(seq.b*k, seq.k*k, seq.n);
  }

  // Multiply all members of the sequence by k
  public static Seq operator *(Seq seq, int k) {
    return new Seq(seq.b*k, seq.k*k, seq.n);
  }

  // Reverse the sequence
  public static Seq operator !(Seq seq) {
    return new Seq(seq.b+(seq.n-1)*seq.k, -seq.k, seq.n);
  }

  // Equality and inequality
  public static bool operator ==(Seq s1, Seq s2) {
    return s1.n==s2.n && (s1.n==0 || s1.b==s2.b && (s1.n==1 || s1.k==s2.k));
  }

  public static bool operator !=(Seq s1, Seq s2) { return !(s1==s2); }

  public override bool Equals(Object that) {
    return that is Seq && this==(Seq)that;
  }

  public override int GetHashCode() { 
    return n==0 ? 0 : n==1 ? b : b^k^n;
  }

  // Get enumerator for the sequence
  public IEnumerator<int> GetEnumerator() {
    return new SeqEnumerator(this);
  }

  // Get enumerator for the sequence
  SC.IEnumerator SC.IEnumerable.GetEnumerator() {
    return GetEnumerator();
  }

  // An enumerator for a sequence, used in foreach statements
  private class SeqEnumerator : IEnumerator<int> {    // Static member class
    private readonly Seq seq;
    private int i;

    public SeqEnumerator(Seq seq) {
      this.seq = seq; Reset();
    }
    
    public int Current {		// For IEnumerator<int>
      get { 
        if (0 <= i && i < seq.n) 
          return seq.b + seq.k * i; 
        else
          throw new InvalidOperationException();
      }
    }
    
    Object SC.IEnumerator.Current { get { return Current; } }

    public bool MoveNext() {		// For IEnumerator<int> and IEnumerator
      return ++i < seq.n;
    }

    public void Reset() {		// For IEnumerator
      i = -1;
    }

    public void Dispose() { }		// For IEnumerator<int>
  }

  public int Count {
    get { return n; } 
  }

  public int this[int i] {
    get { 
      if (0 <= i && i < n) 
        return b + k * i;
      else
        throw new ArgumentOutOfRangeException("Seq indexer: " + i);
    }
  }

  public int[] this[params int[] ii] {
    get { 
      int[] res = new int[ii.Length];
      for (int h=0; h<res.Length; h++)
        res[h] = this[ii[h]];
      return res;
    }
  }

  public void Print() {
    IEnumerator<int> etor = GetEnumerator();
    while (etor.MoveNext()) 
      Console.Write(etor.Current + " ");
  }

  public override String ToString() {
    StringBuilder sb = new StringBuilder();
    foreach (int i in this) 
      sb.Append(i).Append(" ");
    return sb.ToString();
  }
}

class TestSeq {
  public static void Main(String[] args) {
    Seq s1 = new Seq(1, 3);             // 1 2 3
    Seq s2 = 2 * s1 + 5;                // 7 9 11
    Seq s3 = s2 * 3;                    // 21 27 33
    Seq s4 = !s3;                       // 33 27 21
    Console.WriteLine(s1); 
    Console.WriteLine(s2);             
    Console.WriteLine(s3);
    Console.WriteLine(s4);
    Console.WriteLine(s1==s2);                          // False
    Console.WriteLine(s3==!s4);                         // True
    Console.WriteLine(new Seq()==new Seq(5,7,0));       // True    
    Console.WriteLine(new Seq(17,17)==new Seq(17,5,1)); // True    
    s4.Print();                         // 33 27 21
    Console.WriteLine();                // 33 27 21
    for (int i=0, stop=s4.Count; i<stop; i++)
      Console.Write(s4[i] + " ");       
    Console.WriteLine();
    int[] r = s4[2, 2, 1, 2, 0];
    for (int i=0, stop=r.Length; i<stop; i++)
      Console.Write(r[i] + " ");        // 21 21 27 21 33
    Console.WriteLine();
  }
}

interface ISeq : IEnumerable<int> { 
  int Count { get; }
  int this[int i] { get; }
  int[] this[params int[] ii] { get; }
}
