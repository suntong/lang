// Example 89 from page 73 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;
using System.Collections.Generic;

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

class Dummy {
  public static void Main() { }
}
