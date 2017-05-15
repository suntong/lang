// Example 126 from page 101 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

class Example126 {
  public static void Main(String[] args) {
    if (args.Length != 1)
      Console.WriteLine("Usage: Example126 <string>\n");
    else {
      String q = args[0];
      Console.WriteLine(q + " substring of hjsdfk: " + Substring1(q, "hjsdfk"));
    }
  }

  // Decide whether query is a substring of target (using for and while);
  // recommended

  static bool Substring1(String query, String target) {
    for (int j=0, n=target.Length-query.Length; j<=n; j++) {
      int k=0, m=query.Length;
      while (k<m && target[j+k] == query[k])
	      k++;
      // Now k>=m (and target[j..]==query[0..m-1]) or target[j+k] != query[k]
      if (k>=m)
	      return true;
    }
    return false;
  }
}
