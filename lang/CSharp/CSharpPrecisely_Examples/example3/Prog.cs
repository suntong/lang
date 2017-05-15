// Example 3 (file example3/Prog.cs) from C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

// Compile with:
// csc /target:module Mod.cs
// csc /target:library Lib.cs
// csc /addmodule:Mod.netmodule /reference:Lib.dll Prog.cs

using System;

class Class {
  public static void Main(String[] args) {
    if (args.Length != 1) 
      Console.WriteLine("Usage: Prog <yourname>");
    else {
      Console.WriteLine("Hello, " + args[0]);
      ModClass.Hello(args[0]);
      LibClass.Hello(args[0]);
    }
  }
}
