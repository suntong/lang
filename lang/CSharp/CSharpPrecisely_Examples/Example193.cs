// Example 193 from page 157 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;
using System.IO;

public class RandomAccessFileExample {
  public static void Main() {
    for (int i=0; i<4; i++)
      Console.WriteLine(ReadOneString("dna.dat", i));
  }

  static String ReadOneString(String filename, int i) {
    const int IntSize = 4, LongSize = 8;
    FileStream raf = new FileStream(filename, FileMode.Open);
    raf.Seek(raf.Length - IntSize, SeekOrigin.Begin);
    BinaryReader br = new BinaryReader(raf);
    int N = br.ReadInt32();
    raf.Seek(raf.Length - IntSize - LongSize * N + LongSize * i, SeekOrigin.Begin);
    long si = br.ReadInt64();
    raf.Seek(si, SeekOrigin.Begin);
    String s = br.ReadString();
    br.Close();
    return s;
  }
}
