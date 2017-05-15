// Example 86 from page 71 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

class MyTest {
  public static void Main(String[] args) {
    IsNearPoint(new Point(30, 20));
    IsNearPoint(new Point(4, 5));
    IsNearPoint("foo");
    IsNearPoint(null);
  }

  public static void IsNearPoint(Object o) {
    Point p = o as Point;
    if (p != null && p.x*p.x + p.y*p.y <= 100)
      Console.WriteLine(p + " is a Point near (0,0)");
    else
      Console.WriteLine(o + " is not a Point or not near (0,0)");
  }
}

public class Point {
  protected internal int x, y;

  public Point(int x, int y) { this.x = x; this.y = y; }

  public void Move(int dx, int dy) { x += dx; y += dy; }

  public override String ToString() { return "(" + x + ", " + y + ")"; }
}
