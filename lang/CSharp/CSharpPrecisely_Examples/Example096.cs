// Example 96 from page 79 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

// Passing an object by value is much the same as passing a struct by
// reference.

using System;

public class Point {
  protected internal int x, y;

  public Point(int x, int y) { this.x = x; this.y = y; }

  public void Move(int dx, int dy) { x += dx; y += dy; }

  public override String ToString() { return "(" + x + ", " + y + ")"; }
}

public struct SPoint {
  internal int x, y;

  public SPoint(int x, int y) { this.x = x; this.y = y; }

  public void Move(int dx, int dy) { x += dx; y += dy; }

  public override String ToString() { return "(" + x + ", " + y + ")"; }
}

class TestClassStruct {
  public static void Main(String[] args) {
    AssignPointClass();
    AssignPointStruct();
    PassClassStruct();
  }

  static void AssignPointClass() {
    Point p = new Point(11, 111);
    Point q = new Point(22, 222);
    Point r = new Point(33, 333);
    r = q;
    r.x = 44;
    Console.WriteLine("p = {0}, q = {1}, r = {2}", p, q, r);
  }

  static void AssignPointStruct() {
    SPoint p = new SPoint(11, 111);
    SPoint q = new SPoint(22, 222);
    SPoint r = new SPoint(33, 333);
    r = q;
    r.x = 44;
    Console.WriteLine("p = {0}, q = {1}, r = {2}", p, q, r);
  }

  static void PassClassStruct() {
    double d1 = 1.1, d2 = 2.2;
    int[] a1 = new int[4], a2 = new int[4];
    M(d1, ref d2, a1, ref a2);
    Console.WriteLine("d1 = {0}, d2 = {1}", d1, d2);
    Console.WriteLine("a1.Length = {0}, a1[0] = {1}", a1.Length, a1[0]);
    Console.WriteLine("a2.Length = {0}, a2[0] = {1}", a2.Length, a2[0]);

    Point pc1 = new Point(55, 555), pc2 = new Point(66, 666);
    SPoint ps1 = new SPoint(77, 777), ps2 = new SPoint(88, 888);
    M(pc1, ref pc2, ps1, ref ps2);
    Console.WriteLine("pc1 = {0}, pc2 = {1}", pc1, pc2);
    Console.WriteLine("ps1 = {0}, ps2 = {1}", ps1, ps2);
  }

  static void M(double dd1, ref double dd2, int[] aa1, ref int[] aa2) {
    dd1 = 3.3; dd2 = 4.4;
    aa1[0] = 17; 
    aa2[0] = 18; 
    aa2 = new int[3];
    aa1 = aa2;
  }

  static void M(Point ppc1, ref Point ppc2, SPoint pps1, ref SPoint pps2) {
    ppc1.x = 97; 
    ppc2 = new Point(16, 17);
    ppc1 = ppc2;
    pps1.x = 98; 
    pps1 = new SPoint(18, 19);
    pps2.x = 99;
  }
}
