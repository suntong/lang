// Example 133 from page 105 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

class Example133 {
  public static void Main(String[] args) {
    if (args.Length != 1)
      Console.WriteLine("Usage: Example133 <string>\n");
    else {
      String q = args[0];
      Console.WriteLine(q + " substring of hjsdfk: " + Substring1(q, "hjsdfk"));
    }
  }

  // Decide whether query is a substring of target (using goto)

  static bool Substring1(String query, String target) {
    for (int j=0, n=target.Length-query.Length; j<=n; j++) {
      for (int k=0, m=query.Length; k<m; k++)
	if (target[j+k] != query[k])
	  goto nextPos;
      return true;
    nextPos: { }		// Label on empty statement
    }
    return false;
  }
}


