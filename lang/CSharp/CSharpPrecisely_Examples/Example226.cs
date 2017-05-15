// Example 226 from page 187 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

// Representing sets of ints using a HashSet.
// sestoft@itu.dk 2001, 2011-08-05

// In contrast to the .NET 4.0 HashSet<T> implementation, the
// GetHashCode for this set representation is based on the set's
// elements, so Set<Set<int>> and Dictionary<Set<int>>,V> work as one
// would expect from mathematics.  We cache the hash code between
// calls to GetHashCode() to avoid needlessly recomputing it, and of
// course invalidate it when the collection is modified by Add, Clear
// or Remove.

// Note that we use composition (private field HashSet<T> inner).  It
// does not work to use subclassing (class Set<T> : HashSet<T>)
// because the Add and Remove methods on HashSet<T> are non-virtual,
// so the validity of the cached hash code can be undermined by
// calling ((HashSet<T>)set).Add(item).  Dubious .NET 4.0 design,
// probably to save some nanoseconds.

// Computing the intersection closure to illustrate a worklist algorithm. 

using System;                           // Console
using System.Text;                      // StringBuilder
using SC = System.Collections;          // IEnumerable, IEnumerator
using System.Collections.Generic;       // Dictionary<K,V>, IEnumerable<T>

class Set<T> : IEquatable<Set<T>>, ICollection<T> where T : IEquatable<T> {
  private readonly HashSet<T> inner = new HashSet<T>();
  private int? cachedHash = null;   // Cached hash code is valid if non-null

  public Set() { }

  public Set(T x) : this() { 
    Add(x);
  }

  public Set(IEnumerable<T> coll) : this() {
    foreach (T x in coll) 
      Add(x);
  }

  public bool Contains(T x) {
    return inner.Contains(x);
  }

  public void Add(T x) {
    if (!Contains(x)) {
      inner.Add(x);
      cachedHash = null;
    }
  }

  public bool Remove(T x) {
    bool removed = inner.Remove(x);
    if (removed)
      cachedHash = null;
    return removed;
  }

  public IEnumerator<T> GetEnumerator() {
    return inner.GetEnumerator();
  }

  SC.IEnumerator SC.IEnumerable.GetEnumerator() {
    return GetEnumerator();
  }

  public int Count {
    get { return inner.Count; }
  }
  
  public void CopyTo(T[] arr, int i) {
    inner.CopyTo(arr, i);
  }

  public void Clear() {
    inner.Clear();
    cachedHash = null;
  }

  public bool IsReadOnly {
    get { return false; }
  }

  // Is this set a subset of that?
  public bool IsSubsetOf(Set<T> that) { 
    foreach (T x in this)
      if (!that.Contains(x))
        return false;
    return true;            
  }

  // Create new set as intersection of this and that
  public Set<T> Intersection(Set<T> that) { 
    Set<T> res = new Set<T>();
    foreach (T x in this)
      if (that.Contains(x))
        res.Add(x);
    return res;
  }

  // Create new set as union of this and that
  public Set<T> Union(Set<T> that) { 
    Set<T> res = new Set<T>(this);
    foreach (T x in that)
      res.Add(x);
    return res;
  }

  // Create new set as difference between this and that
  public Set<T> Difference(Set<T> that) { 
    Set<T> res = new Set<T>();
    foreach (T x in this)
      if (!that.Contains(x))
        res.Add(x);
    return res;
  }

  // Create new set as symmetric difference between this and that
  public Set<T> SymmetricDifference(Set<T> that) { 
    Set<T> res = new Set<T>();
    foreach (T x in this)
      if (!that.Contains(x))
        res.Add(x);
    foreach (T x in that)
      if (!this.Contains(x))
        res.Add(x);
    return res;
  }

  // Compute hash code based on set contents, and cache it
  public override int GetHashCode() { 
    if (!cachedHash.HasValue) {
      int res = 0;
      foreach (T x in this)
        res ^= x.GetHashCode();
      cachedHash = res;
    }
    return cachedHash.Value;
  }

  public bool Equals(Set<T> that) { 
    return that != null && that.Count == this.Count && that.IsSubsetOf(this);
  }

  public override String ToString() {
    StringBuilder res = new StringBuilder();
    res.Append("{ ");
    bool first = true;
    foreach (T x in this) {
      if (!first) 
        res.Append(", ");
      res.Append(x);
      first = false;
    }
    res.Append(" }");
    return res.ToString();
  }
}

class IntersectionClosure {
  public static void Main(String[] args) {
    Set<Set<int>> SS = new Set<Set<int>>();
    SS.Add(new Set<int>(new int[] { 2, 3 }));
    SS.Add(new Set<int>(new int[] { 1, 3 }));
    SS.Add(new Set<int>(new int[] { 1, 2 }));
    Console.WriteLine("SS = " + SS);
    Set<Set<int>> TT = IntersectionClose(SS);
    Console.WriteLine("TT = " + TT);
  }

  // Given a set SS of sets of Integers, compute its intersection
  // closure, that is, the least set TT such that SS is a subset of TT
  // and such that for any two sets t1 and t2 in TT, their
  // intersection is also in TT.  

  // For instance, if SS is {{2,3}, {1,3}, {1,2}}, 
  // then TT is {{2,3}, {1,3}, {1,2}, {3}, {2}, {1}, {}}.

  // Both the argument and the result is a Set<Set<int>>

  static Set<Set<T>> IntersectionClose<T>(Set<Set<T>> ss) where T : IEquatable<T> {
    Queue<Set<T>> worklist = new Queue<Set<T>>(ss);
    Set<Set<T>> tt = new Set<Set<T>>();
    while (worklist.Count != 0) {
      Set<T> s = worklist.Dequeue();
      foreach (Set<T> t in tt) {
        Set<T> ts = t.Intersection(s);
        if (!tt.Contains(ts)) 
          worklist.Enqueue(ts);
      }
      tt.Add(s);
    }
    return tt;
  }
}
