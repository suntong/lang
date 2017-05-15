// Example 190 from page 155 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;
using System.IO;

public class BinaryIOExample {
  public static void Main() {
    BinaryWriter bw = 
      new BinaryWriter(new FileStream("tmp1.dat", FileMode.Create));
    WriteData(bw); bw.Close();
    BinaryReader br = 
      new BinaryReader(new FileStream("tmp1.dat", FileMode.Open));
    ReadData(br);
  }

  static void WriteData(BinaryWriter bw) {
    bw.Write(true);                             // Write 1 byte
    bw.Write((byte)120);                        // Write 1 byte
    bw.Write('A');                              // Write 1 byte (UTF-8)
    bw.Write("foo");                            // Write 1+3 bytes (UTF-8)
    bw.Write("Rhône");                          // Write 1+6 bytes (UTF-8)
    bw.Write(300.1);                            // Write 8 bytes
    bw.Write(300.2F);                           // Write 4 bytes
    bw.Write(1234);                             // Write 4 bytes
    bw.Write(12345L);                           // Write 8 bytes
    bw.Write((short)32000);                     // Write 2 bytes
    bw.Write((sbyte)-1);                        // Write 1 byte
    bw.Write((short)-1);                        // Write 2 bytes
  }

  static void ReadData(BinaryReader br) {
    Console.Write(      br.ReadBoolean());      // Read 1 byte
    Console.Write(" " + br.ReadByte());         // Read 1 byte
    Console.Write(" " + br.ReadChar());         // Read 1 byte
    Console.Write(" " + br.ReadString());       // Read 1+3 bytes
    Console.Write(" " + br.ReadString());       // Read 1+6 bytes
    Console.Write(" " + br.ReadDouble());       // Read 8 bytes
    Console.Write(" " + br.ReadSingle());       // Read 4 bytes
    Console.Write(" " + br.ReadInt32());        // Read 4 bytes
    Console.Write(" " + br.ReadInt64());        // Read 8 bytes
    Console.Write(" " + br.ReadInt16());        // Read 2 bytes
    Console.Write(" " + br.ReadSByte());        // Read 1 byte
    Console.Write(" " + br.ReadUInt16());       // Read 2 bytes
    Console.WriteLine();
  }
}
