// Example 81 from page 69 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;
using System.Collections.Generic;

class MyTest {
  public static void Main(String[] args) {
    Point p1 = new Point { x = 10, y = 12 };
    Point p2 = new Point();                     // Equivalent to p1 initialization
    p2.x = 10;
    p2.y = 12;
    Point p3 = new Point(p1) { y = 17 };
    Point p4 = new Point(p1);                   // Equivalent to p3 initialization
    p4.y = 17;
    Console.WriteLine(p1);
    Console.WriteLine(p2);
    Console.WriteLine(p3);
    Console.WriteLine(p4);

    Rectangle r1 = new Rectangle { Ul = new Point { x = 10, y = 12 }, 
                                   Lr = new Point { x = 14, y = 20 } };
    Rectangle r2 = new Rectangle { Ul = { x = 10, y = 12 }, Lr = { x = 14, y = 20 } };
    Rectangle r3 = new Rectangle();
    r3.Ul.x = 10;
    r3.Ul.y = 12;
    r3.Lr.x = 14;
    r3.Lr.y = 20;
    Console.WriteLine(r1);
    Console.WriteLine(r2);
    Console.WriteLine(r3);
  }
}

public class Rectangle {
  public Point Ul { get; set; }         // Upper left corner
  public Point Lr { get; set; }         // Lower right corner
  public Rectangle() {
    Ul = new Point();
    Lr = new Point();
  }
  public override String ToString() { 
    return "{ Ul=" + Ul + ", Lr=" + Lr + "}"; 
  }
}

public class Point {
  protected internal int x, y;

  public Point(int x, int y)           
  { this.x = x; this.y = y; }

  public Point()                        
  { }

  public Point(Point p)                 
    : this(p.x, p.y) {}                 

  public void Move(int dx, int dy)
  { x += dx; y += dy; }

  public override String ToString()
  { return "(" + x + ", " + y + ")"; }
}
