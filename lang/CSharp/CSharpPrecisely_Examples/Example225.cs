// Example 225 from page 187 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

// File index: read a text file, build and print a list of all words
// and the line numbers (possibly with duplicates) at which they occur.

using System;                           // Console
using System.Collections.Generic;       // Dictionary, List
using System.IO;                        // StreamReader, TextReader
using System.Text.RegularExpressions;   // Regex

class Example225 {
  static void Main(String[] args) {
    if (args.Length != 1)
      Console.WriteLine("Usage: Example225 <filename>\n");
    else { 
      IDictionary<String, List<int>> index = IndexFile(args[0]);
      PrintIndex(index);
    }
  }

  static IDictionary<String, List<int>> IndexFile(String filename) {
    IDictionary<String, List<int>> index = new Dictionary<String, List<int>>();
    Regex delim = new Regex("[^a-zA-Z0-9]+");
    TextReader rd = new StreamReader(filename);
    int lineno = 0;
    String line;
    while (null != (line = rd.ReadLine())) {
      String[] res = delim.Split(line);
      lineno++;
      foreach (String s in res)
        if (s != "") {
          if (!index.ContainsKey(s)) 
            index[s] = new List<int>();
          index[s].Add(lineno);
        }
    }
    rd.Close();
    return index;
  }

  static void PrintIndex(IDictionary<String, List<int>> index) {
    List<String> words = new List<String>(index.Keys);
    words.Sort();
    foreach (String word in words) {
      Console.Write("{0}: ", word);
      foreach (int ln in index[word])
        Console.Write("{0} ", ln);
      Console.WriteLine();
    }
  }
}
