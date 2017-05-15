// Example 91 from page 75 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;
using System.Collections.Generic;

class TestAPoint {
  public static void Main(String[] args) {
    Console.WriteLine("Number of points created: " + APoint.GetSize());
    APoint p = new APoint(12, 123), q = new APoint(200, 10), r = new APoint(99, 12);
    APoint s = p;
    q = null;
    Console.WriteLine("Number of points created: " + APoint.GetSize());
    Console.WriteLine("r is point number " + r.GetIndex());
    for (int i=0; i<APoint.GetSize(); i++)
      Console.WriteLine("APoint number " + i + " is " + APoint.GetPoint(i));
  }
}

public class APoint {
  private static List<APoint> allpoints = new List<APoint>();
  private int x, y;
  public APoint(int x, int y) { 
    allpoints.Add(this); this.x = x; this.y = y; 
  }
  public void Move(int dx, int dy) { 
    x += dx; y += dy; 
  }
  public override String ToString() { 
    return "(" + x + ", " + y + ")"; 
  }
  public int GetIndex() { 
    return allpoints.IndexOf(this); 
  }
  public static int GetSize() { 
    return allpoints.Count; 
  }
  public static APoint GetPoint(int i) { 
    return allpoints[i]; 
  }
}
