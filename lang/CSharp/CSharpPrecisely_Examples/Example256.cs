// Example 256 from page 219 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

// This is based to a large extent on the generic LinkedList example.

using System;
using System.IO;                        // TextWriter
using System.Collections.Generic;       // IEnumerable<T>, IEnumerator<T>
using SC = System.Collections;          // IEnumerable, IEnumerator
using System.Linq;                      // For Linq query syntax 

public interface IMyList<T> : IEnumerable<T>, IEquatable<IMyList<T>> {
  int Count { get; }                    // Number of elements
  T this[int i] { get; set; }           // Get or set element at index i
  void Add(T item);                     // Add element at end
  void Insert(int i, T item);           // Insert element at index i
  void RemoveAt(int i);                 // Remove element at index i
  IMyList<U> Map<U>(Func<T,U> f);       // Map f over all elements
  void Apply(Action<T> act);		// Apply act to all elements
}

public class LinkedList<T> : IMyList<T> {
  protected int size;			// Number of elements in the list
  protected Node first, last;		// Invariant: first==null iff last==null

  protected class Node {		// Static member class
    public Node prev, next;
    public T item;

    public Node(T item) {
      this.item = item; 
    }

    public Node(T item, Node prev, Node next) {
      this.item = item; this.prev = prev; this.next = next; 
    }
  }

  public LinkedList() {
    first = last = null;
    size = 0;
  }

  public int Count { get { return size; } }     // Property with get accessor

  public T this[int i] {                        // Indexer with get and set accessors
    get { return get(i).item; }
    set { get(i).item = value; }
  }      

  private Node get(int n) {
    if (n < 0 || n >= size)
      throw new IndexOutOfRangeException();
    else if (n < size/2) {              // Closer to front
      Node node = first;
      for (int i=0; i<n; i++)
        node = node.next;
      return node;
    } else {                            // Closer to end
      Node node = last;
      for (int i=size-1; i>n; i--)
        node = node.prev;
      return node;
    }
  }

  public void Add(T item) {                     // Enables collection initializer
    Insert(size, item); 
  }

  public void Insert(int i, T item) { 
    if (i == 0) {
      if (first == null) // and thus last == null
        first = last = new Node(item);
      else {
        Node tmp = new Node(item, null, first);
        first.prev = tmp;
        first = tmp;
      }
      size++;
    } else if (i == size) {
      if (last == null) // and thus first = null
        first = last = new Node(item);
      else {
        Node tmp = new Node(item, last, null);
        last.next = tmp;
        last = tmp;
      }
      size++; 
    } else {
      Node node = get(i);
      // assert node.prev != null;
      Node newnode = new Node(item, node.prev, node);
      node.prev.next = newnode;
      node.prev = newnode;
      size++;
    }
  }

  public void RemoveAt(int i) {
    Node node = get(i);
    if (node.prev == null) 
      first = node.next;
    else
      node.prev.next = node.next;
    if (node.next == null) 
      last = node.prev;
    else
      node.next.prev = node.prev;       
    size--;
  }

  public override bool Equals(Object that) {
    return Equals(that as IMyList<T>);			// Exact runtime type test
  }

  public bool Equals(IMyList<T> that) {
    if (this == that)
      return true;
    if (that == null || this.Count != that.Count) 
      return false;
    Node thisnode = this.first;
    IEnumerator<T> thatenm = that.GetEnumerator();
    while (thisnode != null) {
      if (!thatenm.MoveNext())
	throw new ApplicationException("Impossible: LinkedList<T>.Equals");
      // assert MoveNext() was true (because of the above size test)
      if (!thisnode.item.Equals(thatenm.Current))
	return false;
      thisnode = thisnode.next; 
    }
    // assert !MoveNext(); // because of the size test
    return true;
  }

  public override int GetHashCode() {
    int hash = 0;
    foreach (T x in this)
      hash ^= x.GetHashCode();
    return hash;
  }

  public IEnumerator<T> GetEnumerator() {		// IEnumerable<T> via iterator block
    for (Node curr=first; curr!=null; curr=curr.next) 
      yield return curr.item;
  }

  SC.IEnumerator SC.IEnumerable.GetEnumerator() {
    return GetEnumerator();
  }

  // Explicit conversion from array of T 
  // (Note: <T> is part of the target type name, not not a method type parameter)

  public static explicit operator LinkedList<T>(T[] arr) {
    var res = new LinkedList<T>();
    foreach (T x in arr)
      res.Add(x);
    return res;
  }

  // Overloaded operator

  public static LinkedList<T> operator +(LinkedList<T> xs1, LinkedList<T> xs2) {
    var res = new LinkedList<T>();
    foreach (T x in xs1)
      res.Add(x);
    foreach (T x in xs2)
      res.Add(x);
    return res;
  }

  // Methods with Func and Action arguments 

  public IMyList<U> Map<U>(Func<T,U> f) {
    var res = new LinkedList<U>();
    foreach (T x in this) 
      res.Add(f(x));
    return res;
  }

  public static LinkedList<T> Tabulate(Func<int,T> f, int from, int to) {
    var res = new LinkedList<T>();
    for (int i=from; i<to; i++) 
      res.Add(f(i));
    return res;
  }

  public void Apply(Action<T> act) {                 // Taking delegate argument
    foreach (T x in this) 
      act(x);
  }
}

public class TestLinkedList {
  static void Main(String[] args) {
    LinkedList<int> xs = new LinkedList<int> { 0, 2, 4, 6, 8 };                 // (1)
    Console.WriteLine(xs.Count + " " + xs[2]);                                  // (2)
    xs[2] = 102;                                                                // (3)
    foreach (int k in xs)                                                      
      Console.WriteLine(k);
    LinkedList<int> ys = (LinkedList<int>)(new int[] { 1, 2, 3, 4, 5 });        // (4)
    LinkedList<int> zs = xs + ys;                                               // (5)
    zs.Apply(delegate(int x) { Console.Write(x + " "); });                      // (6)
    Console.WriteLine();
    var vs = LinkedList<double>.Tabulate(x => 1.0/x, 1, 5);                     // (7)
    foreach (var g in from z in zs group z by z/10)                             // (8)
      Console.WriteLine("{0} to {1}: {2} items", g.Key*10, g.Key*10+9, g.Count());
    LinkedList<dynamic> ds = new LinkedList<dynamic>();                         // (9)
    ds.Add(5); ds.Add(0.25); ds.Add(true); ds.Add("foo");                       // (9)
    double d = ds[1];                                                           // (9)
    Console.WriteLine(ds[2] ? ds[3].Length : false);                            // (9)
    Console.WriteLine(xs.Equals(xs));
    Console.WriteLine(xs.Equals(ys));
  }
}
