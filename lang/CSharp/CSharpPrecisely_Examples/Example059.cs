// Example 59 from page 47 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;
using System.Collections.Generic;	// For IEnumerable<T>
using System.Text;			// For StringBuilder

class MyTest {
  public static void Main(String[] args) {
    Console.WriteLine(DateTime.Today.IsoWeek());
    Console.WriteLine(new DateTime(2010, 8, 27).IsoWeek());
    String[] sarr = { "www", "itu", "dk" };
    Console.WriteLine(sarr.ConcatWith("."));
    Console.WriteLine(sarr.IsSorted());
    Console.WriteLine(new int[] { 2, 3, 5 }.IsSorted());
  }
}

public static class DateTimeExtensions {
  // ISO week number: Week 1 of a year contains its first Thursday
  public static int IsoWeek(this DateTime dt) {
    int yday = dt.DayOfYear-1, wday = IsoWeekDay(dt), y = dt.Year;
    const int THU = 3;
    int week = (yday - wday + THU + 7)/7;
    if (week == 0) {
      int prevyear = DateTime.IsLeapYear(y-1) ? 366 : 365;
      return (yday + prevyear - wday + THU + 7)/7;
    } else if (week == 53 && IsoWeekDay(new DateTime(y, 12, 31)) < THU) {
      return 1;
    } else {
      return week;
    }
  }
  
  // Auxiliary method: ISO weekdays: Mon=0, Tue=1, ..., Sun=6
  private static int IsoWeekDay(DateTime dt) {
    return ((int)dt.DayOfWeek + 6) % 7;
  }
}

public static class StringArrayExtensions {
  public static String ConcatWith(this String[] arr, String sep) {
    StringBuilder sb = new StringBuilder();
    if (arr.Length > 0) 
      sb.Append(arr[0]);
    for (int i=1; i<arr.Length; i++)
      sb.Append(sep).Append(arr[i]);
    return sb.ToString();
  }
}

public static class EnumerableExtensions {
  public static bool IsSorted<T>(this IEnumerable<T> xs) where T : IComparable<T> {
    var etor = xs.GetEnumerator(); 
    if (etor.MoveNext()) {
      T prev = etor.Current;
      while (etor.MoveNext())
	if (prev.CompareTo(etor.Current) > 0)
	  return false; 
	else 
	  prev = etor.Current;
    }
    return true;
  }
}

// Extension method scope:

class My { }

namespace Outer {
  static class MyExtensions {
    public static void Extension1(this My my) { }
  }

  namespace Inner {
    static class MyExtensions {
      public static void Extension2(this My my) { }
    }

    class Try2C {
      public static void Try() {
	new My().Extension1(); // In scope here
	new My().Extension2(); // In scope here
      }
    }
  }
  
  class Try1C {
    public static void Try() {
      new My().Extension1();	// In scope here
      // new My().Extension2(); // Not in scope here
    }
  }
}
