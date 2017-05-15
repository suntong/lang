// Example 230 from page 191 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;
using System.Collections.Generic;
using System.Diagnostics;

class Example230 {
  static readonly String[] keywordarray = 
    { "abstract", "as", "base", "bool", "break", "byte", "case", "catch",
      "char", "checked", "class", "const", "continue", "decimal", "default",
      "delegate", "do", "double", "else", "enum", "event", "explicit",
      "extern", "false", "finally", "fixed", "float", "for", "foreach",
      "goto", "if", "implicit", "in", "int", "interface", "internal", "is",
      "lock", "long", "namespace", "new", "null", "object", "operator",
      "out", "override", "params", "private", "protected", "public",
      "readonly", "ref", "return", "sbyte", "sealed", "short", "sizeof",
      "stackalloc", "static", "string", "struct", "switch", "this", "throw",
      "true", "try", "typeof", "uint", "ulong", "unchecked", "unsafe",
      "ushort", "using", "virtual", "void", "volatile", "while" };

  static readonly ISet<String> keywords = new HashSet<String>();
  
  static Example230() {
    foreach (String keyword in keywordarray) 
      keywords.Add(keyword);
  }
    
  static bool IsKeyword1(String id) { 
    return keywords.Contains(id); 
  }

  static bool IsKeyword2(String id) { 
    return Array.BinarySearch(keywordarray, id) >= 0; 
  }
      
  public static void Main(String[] args) {
    if (args.Length != 2) 
      Console.WriteLine("Usage: Example230 <iterations> <word>\n");
    else {
      int count = int.Parse(args[0]);
      String id = args[1];
      for (int i=0; i<keywordarray.Length; i++)
        if (IsKeyword1(keywordarray[i]) != IsKeyword2(keywordarray[i]))
          Console.WriteLine("Error at i = " + i);
      if (IsKeyword1(id) != IsKeyword2(id))
        Console.WriteLine("Error at id = " + id);
      
      Console.Write("HashSet.Contains ");
      Stopwatch sw = new Stopwatch();
      sw.Start();
      for (int i=0; i<count; i++)
        IsKeyword1(id);
      Console.WriteLine("{0} ms", sw.ElapsedMilliseconds);      
      
      Console.Write("Array.BinarySearch  ");
      sw.Reset();
      sw.Start();
      for (int i=0; i<count; i++)
        IsKeyword2(id);
      Console.WriteLine("{0} ms", sw.ElapsedMilliseconds); 
    }
  }
}
