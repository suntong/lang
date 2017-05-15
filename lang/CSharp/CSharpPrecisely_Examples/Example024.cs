// Example 24 from page 21 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

public class StringEks {
	public static void Main() {
		
Point p1 = new Point(10, 20), p2 = new Point(30, 40);
Console.WriteLine("p1 is " + p1);      // Prints: p1 is (10, 20)
Console.WriteLine("p2 is " + p2);      // Prints: p2 is (30, 40)
p2.Move(7, 7);
Console.WriteLine("p2 is " + p2);      // Prints: p2 is (37, 47)
		
	}
}

class Point {
  protected internal int x, y;

  public Point(int x, int y) { this.x = x; this.y = y; }

  public void Move(int dx, int dy) { x += dx; y += dy; }

  public override String ToString() { return "(" + x + ", " + y + ")"; }
}
