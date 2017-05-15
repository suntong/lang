// Example 3 (file example3/Lib.cs) from C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

// Compile with:
// csc /target:module Mod.cs
// csc /target:library Lib.cs
// csc /addmodule:Mod.netmodule /reference:Lib.dll Prog.cs

using System;

public class LibClass {
  public static void Hello(Object name) {
    Console.WriteLine("Hello to " + name + " from Lib");
  }    
}
