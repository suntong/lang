// Example 223 from page 185 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;                           // IComparable, IComparable<T>, IEquatable<T>
using System.Collections.Generic;       // IDictionary<K,V>, SortedDictionary<K,V>

public class Time : IComparable, IComparable<Time>, IEquatable<Time> {
  private readonly int hh, mm;                  // 24-hour clock
  
  public Time(int hh, int mm) { this.hh = hh; this.mm = mm; }

  // Return neg if before that; return pos if after that; return zero if same
  public int CompareTo(Object that) {           // For IComparable
    return CompareTo((Time)that);
  }

  public int CompareTo(Time that) {             // For IComparable<T>
    return hh != that.hh ? hh - that.hh : mm - that.mm;
  }

  public bool Equals(Time that) {               // For IEquatable<T>
    return hh == that.hh && mm == that.mm;
  }

  public override String ToString() { 
    return String.Format("{0:00}:{1:00}", hh, mm); 
  }
}

class TestDatebook {
  public static void Main(String[] args) {
    IDictionary<Time,String> datebook = new SortedDictionary<Time,String>();
    datebook.Add(new Time(12, 30), "Lunch");
    datebook.Add(new Time(15, 30), "Afternoon coffee break");
    datebook.Add(new Time( 9,  0), "Lecture");
    datebook.Add(new Time(13, 15), "Board meeting");
    foreach (KeyValuePair<Time,String> entry in datebook) 
      Console.WriteLine(entry.Key + " " + entry.Value);
  }
}
