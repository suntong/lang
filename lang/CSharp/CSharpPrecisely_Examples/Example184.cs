// Example 184 from page 149 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;
using System.IO;
using System.Runtime.Serialization.Formatters.Binary;  // BinaryFormatter

public class IOExample {
  public static void Main() {
    
    // Write numbers and words on file "f.txt" in human-readable form:
    TextWriter twr = new StreamWriter(new FileStream("f.txt", FileMode.Create));
    twr.Write(4711); twr.Write(' '); twr.Write("cool"); twr.Close();
    
    // Write primitive values to a binary file "p.dat":
    BinaryWriter bwr = new BinaryWriter(new FileStream("p.dat", FileMode.Create));
    bwr.Write(4711); bwr.Write(' '); bwr.Write("cool"); bwr.Close();
    
    // Read primitive values from binary file "p.dat":
    BinaryReader brd = new BinaryReader(new FileStream("p.dat", FileMode.Open));
    Console.WriteLine(brd.ReadInt32() + "|" + brd.ReadChar() + "|" + brd.ReadString());
    
    // Write an object or array to binary file "o.dat":
    FileStream fs1 = new FileStream("o.dat", FileMode.Create);
    BinaryFormatter bf = new BinaryFormatter();
    bf.Serialize(fs1, new int[] { 2, 3, 5, 7, 11 }); fs1.Close();
    
    // Read objects or arrays from binary file "o.dat":
    FileStream fs2 = new FileStream("o.dat", FileMode.Open);
    int[] ia = (int[]) bf.Deserialize(fs2);
    Console.WriteLine("{0} {1} {2} {3} {4}", ia[0], ia[1], ia[2], ia[3], ia[4]); fs2.Close();
    
    // Read and write parts of file "raf.dat" in arbitrary order:
    FileStream fs = new FileStream("raf.dat", FileMode.OpenOrCreate, FileAccess.ReadWrite);
    BinaryWriter bw = new BinaryWriter(fs);
    bw.Write(3.1415); bw.Write(42);
    fs.Seek(0, SeekOrigin.Begin);
    BinaryReader br = new BinaryReader(fs);
    Console.WriteLine("{0} {1}", br.ReadDouble(), br.ReadInt32());
    
    // Read from a String as if it were a text file:
    TextReader tr = new StringReader("abc");
    Console.WriteLine("abc: " + (char)tr.Read() + (char)tr.Read() + (char)tr.Read());
    
    // Write to a StringBuffer as if it were a text file:
    TextWriter tw = new StringWriter();
    tw.Write('d'); tw.Write('e'); tw.Write('f');
    Console.WriteLine(tw.ToString());
    
    // Write characters to standard output and standard error:
    Console.Out.WriteLine("std output"); Console.Error.WriteLine("std error");
    
    // Read characters from standard input (the keyboard):
    Console.WriteLine("Type some characters and press Enter: ");
    TextReader intext = Console.In;
    String response = intext.ReadLine();
    Console.WriteLine("You typed: '{0}'", response);
    
    // Read a character from standard input (the keyboard):
    Console.Write("Type one character and press Enter: ");
    char c = (char)Console.In.Read();
    Console.WriteLine("First character of your input is: " + c);
  }
}
