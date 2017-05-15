// Example 106 from page 87 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;
using SC = System.Collections;

class MyTest {
  public static void Main(String[] args) {
    StringList ss = new StringList();
    ss.Add("Cop"); ss.Add("en"); ss.Add("cabana");
    ss[2] = "hagen";
    ss[0] += "en" + ss[2];
    Console.WriteLine("A total of {0} strings", ss.Count);
    String last = ss[2];                        // Correct type
    Console.WriteLine(ss["0"] + "/" + last);    // Prints: Copenhagen/hagen
  }
}

class StringList : SC.ArrayList {               // Needs: using SC = System.Collections;
  public new String this[int i] {
    get { return (String)base[i]; }
    set { base[i] = value; }
  }
  public String this[String s] {
    get { return this[int.Parse(s)]; } 
  }
}
