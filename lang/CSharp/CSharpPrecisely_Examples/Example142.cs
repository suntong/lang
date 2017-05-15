// Example 142 from page 113 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

public struct SPoint {
  internal int x, y;

  public SPoint(int x, int y) { this.x = x; this.y = y; }

  public SPoint Move(int dx, int dy) { x += dx; y += dy; return this; }

  public override String ToString() { return "(" + x + ", " + y + ")"; }
}

class Dummy {
  public static void Main() { }
}
