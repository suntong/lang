// Example 88 from page 73 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

public class Point {
  protected internal int x, y;

  public Point(int x, int y) { 
    this.x = x; this.y = y; 
  }

  public void Move(int dx, int dy) { 
    x += dx; y += dy; 
  }

  public override String ToString() { 
    return "(" + x + ", " + y + ")"; 
  }
}

class Dummy {
  public static void Main() { }
}
