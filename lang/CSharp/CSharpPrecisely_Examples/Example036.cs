// Example 36 from page 29 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

class MyTest {
  public static void Main(String[] args) {
    Point[] a = new RedPoint[10];     // Length 10, element type RedPoint
    Point p1 = new Point(42, 117);    // Compile-time type Point, class Point
    RedPoint cp = new RedPoint(3,4);  // Compile-time type RedPoint, class RedPoint
    Point p2 = cp;                    // Compile-time type Point, class RedPoint
    a[0] = cp;                        // OK, RedPoint is subclass of RedPoint
    a[1] = p2;                        // OK, RedPoint is subclass of RedPoint
    a[2] = p1;                        // Runtime error: Point not subclass of RedPoint
  }
}

public class Point {
  protected internal int x, y;

  public Point(int x, int y) { this.x = x; this.y = y; }

  public void Move(int dx, int dy) { x += dx; y += dy; }

  public override String ToString() { return "(" + x + ", " + y + ")"; }
}

class RedPoint : Point { 
  private uint rgb;

  public RedPoint(int x, int y) : base(x, y) { this.rgb = 0xFF0000; }

  public override String ToString() { 
    return base.ToString() + "@" + rgb; 
  }
}
