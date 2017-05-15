// Example 22 from page 21 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

public class StringEks {
  static bool Sorted(String[] a) {
    for (int i=1; i<a.Length; i++)
      if (a[i-1].CompareTo(a[i]) > 0)
	return false;
    return true;
  }
}
