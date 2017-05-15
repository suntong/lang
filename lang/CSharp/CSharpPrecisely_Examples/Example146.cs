// Example 146 from page 117 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;
using System.Windows.Forms;
using System.Drawing;           // Color, Graphics, SolidBrush, Pen, ...

class UseColored : System.Windows.Forms.Form {
  private static IColoredDrawable[] cs;

  static void PrintColors(IColored[] cs) {
    for (int i=0; i<cs.Length; i++)
      Console.WriteLine(cs[i].GetColor);
  } 

  static void Draw(Graphics g, IColoredDrawable[] cs) {
    for (int i=0; i<cs.Length; i++) {
      Console.WriteLine(cs[i].GetColor);
      cs[i].Draw(g);
    }
  } 

  public static void Main(String[] args) {
    cs = new IColoredDrawable[]
          { new ColoredDrawablePoint(3, 4, Color.Red), 
            new ColoredRectangle(50, 100, 60, 110, Color.Green) };
    PrintColors(cs);
    Application.Run(new UseColored());
  }

  protected override void OnPaint(PaintEventArgs e) {
    Graphics g = e.Graphics;
    Draw(g, cs);
  }    
}

public class Point {
  protected internal int x, y;

  public Point(int x, int y) { this.x = x; this.y = y; }

  public void Move(int dx, int dy) { x += dx; y += dy; }

  public override String ToString() { return "(" + x + ", " + y + ")"; }
}

interface IColored { Color GetColor { get; } }
interface IDrawable { void Draw(Graphics g); }
interface IColoredDrawable : IColored, IDrawable {}

class ColoredPoint : Point, IColored {
  protected Color c;
  public ColoredPoint(int x, int y, Color c) : base(x, y) { this.c = c; }
  public Color GetColor { get { return c; } }
}

class ColoredDrawablePoint : ColoredPoint, IColoredDrawable {
  public ColoredDrawablePoint(int x, int y, Color c) : base(x, y, c) { }
  public void Draw(Graphics g) { 
    g.FillRectangle(new SolidBrush(c), x, y, 2, 2);
  }    
}

class ColoredRectangle : IColoredDrawable {
  private int x1, x2, y1, y2;   // (x1, y1) upper left, (x2, y2) lower right
  protected Color c;

  public ColoredRectangle(int x1, int y1, int x2, int y2, Color c) 
  { this.x1 = x1; this.y1 = y1; this.x2 = x2; this.y2 = y2; this.c = c; }
  public Color GetColor { get { return c; } }
  public void Draw(Graphics g) { 
    g.DrawRectangle(new Pen(c), x1, y1, x2, y2); 
  }
}
