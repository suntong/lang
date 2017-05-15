// Example 201 from page 167 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;
using System.Drawing;                           // Color

class MyTest {
  public static void Main(String[] args) {
    Point<String> p1 = new Point<String>(5, 117, "home"), 
      p2 = new Point<String>(2, 3, "work");
    Point<double> p3 = new Point<double>(10, 100, 3.1415);
    ColorPoint<String,uint> p4 = 
      new ColorPoint<String,uint>(20, 30, "foo", 0x0000FF);
    ColorPoint<String,Color> p5 = 
      new ColorPoint<String,Color>(40, 50, "bar", Color.Blue);
    IMovable[] movables = { p1, p2, p3, p4, p5 };
    Point<String>[] stringpoints = { p1, p4, p5 };
  }
}

interface IMovable {
  void Move(int dx, int dy);
}

class Point<Label> : IMovable {
  protected internal int x, y;
  private Label lab;

  public Point(int x, int y, Label lab) {
    this.x = x; this.y = y; this.lab = lab;
  }

  public void Move(int dx, int dy) { 
    x += dx; y += dy;
  }
  
  public Label Lab { 
    get { return Lab; }
  }
}

class ColorPoint<Label, Color> : Point<Label> {
  private Color c; 

  public ColorPoint(int x, int y, Label lab, Color c) : base(x, y, lab) {
    this.c = c;
  }
}
