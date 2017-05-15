// Example 254 from page 217 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

// Serialize a data structure to a file "objects".  Strangely,
// SoapFormatter and 

using System;
using System.IO;                                       // File, FileMode, Stream
using System.Runtime.Serialization;                    // IFormatter
using System.Runtime.Serialization.Formatters.Soap;    // SoapFormatter
using System.Runtime.Serialization.Formatters.Binary;  // BinaryFormatter

[Serializable()]
class SC { public int ci; }

[Serializable()]
class SO {
  public int i;
  public SC c;
  [NonSerialized()] public String s; 

  public SO(int i, SC c) { this.i = i; this.c = c; s = i.ToString(); }
  public void CPrint() { 
    Console.WriteLine("i{0}c{1}({2})", i, c.ci, s); }
}

class SerializeUnshared {
  public static void Main(String[] args) {
    IFormatter fmtr = new SoapFormatter();
    // IFormatter fmtr = new BinaryFormatter();        // Alternative
    if (!File.Exists("objects")) {
      Console.WriteLine("Creating objects and writing them to file:"); 
      SC c = new SC();
      SO o1 = new SO(1, c), o2 = new SO(2, c); 
      Console.WriteLine("The SC object is shared between o1 and o2:");
      o1.c.ci = 3; o2.c.ci = 4;                 // Update the shared c twice
      o1.CPrint(); o2.CPrint();                 // Prints i1c4 i2c4
      // Open file and serialize objects to it:
      Stream strm = File.Open("objects", FileMode.Create);
      fmtr.Serialize(strm, o1); fmtr.Serialize(strm, o2);
      strm.Close();
      Console.WriteLine("\nRun the example again to read objects from file");
    } else {
      Console.WriteLine("Reading objects from file (unshared c):");
      Stream strm = File.Open("objects", FileMode.Open);
      SO o1i = (SO)(fmtr.Deserialize(strm)), o2i = (SO)(fmtr.Deserialize(strm));
      strm.Close();
      o1i.CPrint(); o2i.CPrint();               // Prints i1c4() i2c4()
      Console.WriteLine("The sharing of the SC object is lost:");
      o1i.c.ci = 5; o2i.c.ci = 6;               // Update two different c's
      o1i.CPrint(); o2i.CPrint();               // Prints i1c5() i2c6()
      File.Delete("objects");
    }
    Console.WriteLine();
  }
}
