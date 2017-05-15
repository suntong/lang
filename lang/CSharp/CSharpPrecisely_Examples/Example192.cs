// Example 192 from page 157 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;
using System.Collections.Generic;
using System.IO;

public class RandomAccessFileExample {
  public static void Main() {
    String[] dna = { "TTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTT",
                     "CCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCC",
                     "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
                     "GGGGGGGGGGGGGGGGGGGGGGG" };
    WriteStrings("dna.dat", dna);
  }

  static void WriteStrings(String filename, String[] dna) {
    FileStream raf = new FileStream(filename, FileMode.Create);
    BinaryWriter sw = new BinaryWriter(raf);
    raf.SetLength(0);                                   // Truncate the file
    List<long> offsettable = new List<long>();
    foreach (String s in dna) {
      offsettable.Add(raf.Position);                    // Store string offset
      sw.Write(s);                                      // Write string
    }

    foreach (long v in offsettable) {                   // Write string offsets
      Console.WriteLine(v);
      sw.Write(v);
    }
    sw.Write(offsettable.Count);                        // Write string count
    sw.Close();
  }
}
