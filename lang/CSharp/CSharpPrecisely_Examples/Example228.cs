// Example 228 from page 189 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

// This example requires:

//  * MySQL-ODBC driver 3.51
//  * A MySQL database server on host ellemose with a database 
//    test containing a table Message declared like this:
//       CREATE TABLE Message 
//              (name VARCHAR(80), 
//               msg VARCHAR(200), 
//               severity INT);
//  * Due to the default security restrictions it cannot be run from a
//    network drive.

using System;
using System.Data.Odbc;           // OdbcConnection OdbcCommand OdbcDataReader
using System.Collections.Generic; // List<T>

class Example228 {
  static Record[] GetMessages(OdbcConnection conn) {
    String query = "SELECT name, msg, severity FROM Message ORDER BY name";
    OdbcCommand cmd = new OdbcCommand(query, conn);
    OdbcDataReader r = cmd.ExecuteReader();
    List<Record> results = new List<Record>();
    while (r.Read()) 
      results.Add(new Record(r.GetString(0), r.GetString(1), r.GetInt32(2)));
    r.Close();
    return results.ToArray();
  }

  struct Record {
    public readonly String name, msg;
    public readonly int severity;
    public Record(String name, String msg, int severity) {
      this.name = name; this.msg = msg; this.severity = severity;
    }
    public override String ToString() {
      return String.Format("{0}: {1} ({2})", name, msg, severity);
    }
  }

  public static void Main(String[] args) {
    if (args.Length != 1) 
      Console.WriteLine("Usage: Example228 <password>\n");
    else {    
      String setup = 
        "DRIVER={MySQL ODBC 3.51 Driver};" 
        + "SERVER=sql.dina.kvl.dk;" 
        + "DATABASE=test;" 
        + "UID=sestoft;" 
        + "PASSWORD=" + args[0] + ";";
      using (OdbcConnection conn = new OdbcConnection(setup)) {
        conn.Open();
        Record[] results = GetMessages(conn);
        foreach (Record rec in results) 
          Console.WriteLine(rec);
      }
    }
  }
}

