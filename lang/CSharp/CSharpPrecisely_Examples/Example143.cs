// Example 143 from page 113 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

public struct SPoint {
  internal int x, y;

  public SPoint(int x, int y) { this.x = x; this.y = y; }

  public void Move(int dx, int dy) { x += dx; y += dy; }

  public override String ToString() { return "(" + x + ", " + y + ")"; }
}

class MyTest {
  public static void Main(String[] args) {
    SPoint p = new SPoint(11, 22);              // Create a struct value in p
    SPoint[] arr = { p, p };                    // Two more copies of p
    arr[0].x = 33;
    Console.WriteLine(arr[0] + " " + arr[1]);   // Prints (33, 22) (11, 22)
    Object o = p;                               // Another copy of p, in heap
    p.x = 44;
    Console.WriteLine(p + " " + o);             // Prints (44, 22) (11, 22)
    Console.WriteLine(o is SPoint);             // Prints True
    Console.WriteLine(o is int);                // Prints False
  }

  public static void TryReverse() {
    SPoint[] arr = new SPoint[5];
    for (int i=0; i<arr.Length; i++)
      arr[i] = new SPoint(i*11, i*22); 
    // Reverse((ValueType[])arr); // Cannot convert SPoint[] to ValueType[]
    Console.WriteLine("Reversed input:");
    Console.WriteLine("--------------------");
    foreach (SPoint q in arr) 
      Console.WriteLine(q);
    Console.WriteLine("--------------------");
  }

  public static void Reverse(ValueType[] arr) {
    for (int s=0, t=arr.Length-1; s<t; s++, t--) {
      ValueType tmp = arr[s]; arr[s] = arr[t]; arr[t] = tmp;
    }
  }
}
