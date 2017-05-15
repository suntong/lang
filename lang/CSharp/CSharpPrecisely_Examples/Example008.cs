// Example 8 from page 9 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

class MyTest {
  public static void Main(String[] args) {
    Object o1 = new Object(), o2 = new Object(), o3 = o1;
    Console.WriteLine(o1.Equals(o3) + " " + o1.Equals(o2));       // True False
    Console.WriteLine(o1.GetHashCode() == o3.GetHashCode());      // True
    Console.WriteLine(o1.GetHashCode() == o2.GetHashCode());      // Usually False
    Console.WriteLine(o1.GetHashCode() + " " + o2.GetHashCode()); // Usually distinct
    Console.WriteLine(o1.GetType());                              // System.Object
    String s1 = "abc", s2 = "ABC", s3 = s1 + "";
    Console.WriteLine(s1.Equals(s3) + " " + s1.Equals(s2));       // True False
    Console.WriteLine(s1.GetHashCode() == s3.GetHashCode());      // True
    Console.WriteLine(s1.GetHashCode() == s2.GetHashCode());      // Usually False
    Console.WriteLine(s1.GetHashCode() + " " + s2.GetHashCode()); // Usually distinct
    Console.WriteLine(s1.GetType());                              // System.String
    Console.WriteLine(117.GetHashCode());                         // 117
    Console.WriteLine(5.GetType());                               // System.Int32
    Console.WriteLine(5.0.GetType());                             // System.Double
    int[] ia1 = { 7, 9, 13 }, ia2 = { 7, 9, 13 };
    Console.WriteLine(ia1.GetType());                             // System.Int32[]
    Console.WriteLine(ia1.Equals(ia2));                           // False
    Console.WriteLine(Object.ReferenceEquals(ia1,ia2));           // False
    Console.WriteLine(ia1.GetHashCode() == ia2.GetHashCode());    // Usually False
    int[,] ia3 = new int[6,7];          
    Console.WriteLine(ia3.GetType());                             // System.Int32[,]
    int[][] ia4 = new int[6][];         
    Console.WriteLine(ia4.GetType());                             // System.Int32[][]
  }
}
