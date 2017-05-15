// Example 71 from page 59 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

public class Point {
  protected internal int x, y;

  public Point(int x, int y) { this.x = x; this.y = y; }
}

struct SPoint {
  public int x, y;

  public SPoint(int x, int y) { this.x = x; this.y = y; }
}

class MyTest {
  public static void M1() {
    Point p = new Point(11, 111), q = new Point(22, 222);
    p = q;
    p.x = 33;
    SPoint r = new SPoint(44, 444), s = new SPoint(55, 555);
    r = s;
    r.x = 66;
    int[] iarr1 = new int[4];
    int[] iarr2 = iarr1;
    iarr1[0] = 77;
    SPoint[] sarr = new SPoint[3];
    sarr[0].x = 88;
    Console.WriteLine("q.x={0} s.x={1} iarr2[0]={2}", p.x, s.x, iarr2[0]);
    M2(2);
  }

  public static void M2(int i) {
    Console.WriteLine(i);
    if (i > 0) 
      M2(i-1);
  }

  public static void Main(String[] args) {
    M1();
  }
}
