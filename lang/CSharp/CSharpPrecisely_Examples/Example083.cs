// Example 83 from page 69 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;
using System.Collections.Generic;

class MyTest {
  public static void Main(String[] args) {
    List<int> list1 = new List<int> { { 2 }, { 3 }, { 2+3 }, { 5+2 } };
    List<int> list2 = new List<int>();
    list2.Add(2); list2.Add(3); list2.Add(2+3); list2.Add(5+2);
    List<int> list3 = new List<int> { 2, 3, 2+3, 5+2 };
    int cinq, sept;
    List<int> list4 = new List<int> { 2, 3, { cinq = 2+3 }, { sept = cinq+2 } };
    Print(list1);
    Print(list2);
    Print(list3);
    Print(list4);

    Dictionary<int,String> numerals 
      = new Dictionary<int,String> { { 1, "one" }, { 2, "two" }, { 5, "five" } };
    Console.WriteLine(numerals[2]);

    Polygon poly = new Polygon { { 1, 1 }, { 1, 4 }, { 4, 4  } };
  }

  private static void Print(List<int> xs) {
    foreach (int x in xs) 
      Console.Write(x + " ");
    Console.WriteLine();
  }
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

public class Polygon : List<Point> {
  public void Add(int x, int y) { base.Add(new Point { x = x, y = y }); }
}
