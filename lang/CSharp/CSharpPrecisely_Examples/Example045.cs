// Example 45 from page 37 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

public abstract class Vessel {
  private double contents;
  public abstract double Capacity();
  public void Fill(double amount) { contents = Math.Min(contents + amount, Capacity()); }
  public double Contents { get { return contents; } }
}
public class Tank : Vessel {
  protected readonly double length, width, height;
  public Tank(double length, double width, double height)
  { this.length = length; this.width = width; this.height = height; }
  public override double Capacity() { return length * width * height; }
  public override String ToString()
  { return "tank (" + length + ", " + width + ", " + height + ")"; }
}
public class Cube : Tank {
  public Cube(double side) : base(side, side, side) {}
  public override String ToString() { return "cube (" + length + ")"; }
}
public class Barrel : Vessel {
  private readonly double radius, height;
  public Barrel(double radius, double height) { this.radius = radius; this.height = height; }
  public override double Capacity() { return height * Math.PI * radius * radius; }
  public override String ToString() { return "barrel (" + radius + ", " + height + ")"; }
}

public class UseVesselHierarchy {
  public static void Main(String[] args) {
    Vessel v1 = new Barrel(3, 10);
    Vessel v2 = new Tank(10, 20, 12);
    Vessel v3 = new Cube(4);
    Vessel[] vs = { v1, v2, v3 };
    v1.Fill(90); v1.Fill(10); v2.Fill(100); v3.Fill(80);
    double sum = 0;
    for (int i=0; i<vs.Length; i++)
      sum += vs[i].Capacity();
    Console.WriteLine("Total capacity is " + sum);
    sum = 0;
    for (int i=0; i<vs.Length; i++)
      sum += vs[i].Contents;
    Console.WriteLine("Total contents is " + sum);
    for (int i=0; i<vs.Length; i++)
      Console.WriteLine("vessel number " + i + ": " + vs[i]);
  }
}
