// Example 170 from page 137 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;
using System.Collections.Generic;       // IList, List, IEnumerable
using System.Net;                       // WebClient
using System.Linq;                      // from ... select ... syntax
using System.Threading.Tasks;           // Parallel
using System.Text;                      // ASCIIEncoding

class MyTest {
  // Entrez E-utilities at the US National Center for Biotechnology Information:
  static readonly String server = "http://www.ncbi.nlm.nih.gov/entrez/eutils/";

  public static void Main(String[] args) {
    ShowResult(NcbiProtein("P01308"));
    ShowResult(NcbiProteinParallel("P01308", "P01315", "P01317"));
    ShowResult(NcbiProteinParallel2("P01308", "P01315", "P01317"));
  }

  private static void ShowResult(String s) {
    Console.WriteLine("\n------------------------------------------------------------");
    Console.WriteLine(s);
  }

  private static void ShowResult(IEnumerable<String> ss) {
    Console.WriteLine("\n------------------------------------------------------------");
    foreach (var s in ss) 
      Console.WriteLine(s);
  }

  public static String NcbiEntrez(String query) {
    byte[] bytes = new WebClient().DownloadData(new Uri(server + query));
    return ASCIIEncoding.ASCII.GetString(bytes);
  }

  public static String NcbiProtein(String id) {
    return NcbiEntrez("efetch.fcgi?rettype=fasta&retmode=text&db=protein&id=" + id);
  }

  public static String[] NcbiProteinParallel(params String[] ids) {
    String[] results = new String[ids.Length];
    Parallel.For(0, ids.Length, i => { results[i] = NcbiProtein(ids[i]); });
    return results;
  }

  public static String[] NcbiProteinParallel2(params String[] ids) {
    IList<String> results = new List<String>();
    Parallel.For(0, ids.Length, 
                 i => { String res = NcbiProtein(ids[i]); 
                        lock (results) results.Add(res); });
    return results.ToArray();
  }
}
