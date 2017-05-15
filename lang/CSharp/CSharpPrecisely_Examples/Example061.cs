// Example 61 from page 49 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

public class Test {
  public static void Main() { }
} 

public class Point {
  protected internal int x, y;

  public Point(int x, int y)            // overloaded constructor
  { this.x = x; this.y = y; }

  public Point()			// overloaded constructor
  { }

  public Point(Point p)                 // overloaded constructor
    : this(p.x, p.y) {}                 // calls the first constructor

  public void Move(int dx, int dy)
  { x += dx; y += dy; }

  public override String ToString()
  { return "(" + x + ", " + y + ")"; }
}
