// Example 144 from page 113 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

public struct SPoint {
  internal int x, y;

  public SPoint(int x, int y) { this.x = x; this.y = y; }

  public SPoint Move(int dx, int dy) { return this = new SPoint(x+dx, y+dy); }

  public override String ToString() { return "(" + x + ", " + y + ")"; }
}

class MyTest {
  static readonly SPoint q = new SPoint(33, 44);

  public static void Main(String[] args) {
    SPoint p = new SPoint(11, 22);
    Console.WriteLine("p = {0}", p);           // Now p = (11, 22)
    p.Move(9,8);
    Console.WriteLine("p = {0}", p);           // Now p = (20, 30)
    p.Move(5,5).Move(6,6);
    Console.WriteLine("p = {0}", p);           // Now p = (25, 35) not (31, 41)
    q.Move(5,5);
    Console.WriteLine("q = {0}", q);           // Now q = (33, 44) not (38, 49)
  }
}
