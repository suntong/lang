// Example 114 from page 95 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

// All the ways a statement may terminate: normally, throw exception,
// return, etc.

using System;

class Example114 {
  public static void Main(String[] args) {
    if (args.Length != 1) 
      Console.WriteLine("Usage: Example114 <integer 1-7>\n");
    else
      Statement(int.Parse(args[0]));
  }

  public static void Statement(int choice) {
    bool again = true;
    while (Again(again)) {
      again = !again;
      if (choice == 1)                  // Terminate normally
        Console.WriteLine("Choice 1");
      else if (choice == 2)             // Throw exception
        throw new Exception();
      else if (choice == 3)             // Return from method
        return;
      else if (choice == 4)             // Break out of loop
        break;
      else if (choice == 5)             // Continue at loop test
        continue; 
      else if (choice == 6)             // Jump out of loop
        goto end;
      else                              // Loop forever
        while (true) { }
      Console.WriteLine("At end of loop");
    } 
    Console.WriteLine("After loop");
  end:
    Console.WriteLine("At end of method");
  }

  private static bool Again(bool again) {
    Console.WriteLine("Loop test");
    return again;
  }
}
