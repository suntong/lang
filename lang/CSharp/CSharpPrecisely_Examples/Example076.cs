// Example 76 from page 65 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

[Flags]                               // Print enum combinations symbolically
public enum FileAccess {
  Read = 1 << 0, 
  Write = 1 << 1
}

class MyTest {
  public static void Write(FileAccess access) {
    if (0 != (access & FileAccess.Write))
      Console.WriteLine("You have write permission");
  }

  public static void Main(String[] args) {
    FileAccess access = FileAccess.Read | FileAccess.Write;
    Console.WriteLine(access);                              // Prints: Read, Write
    Write(access);
  }
}
