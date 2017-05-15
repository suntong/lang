// Example 178 from page 143 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;
using System.Collections.Generic;       // IList, List, IEnumerable
using System.Net;                       // WebClient
using System.Linq;                      // from ... select ... syntax
using System.Threading.Tasks;           // Task<T>
using System.Xml;                       // XmlDocument, XmlNode
using System.Text;                      // ASCIIEncoding

class MyTest {
  // Entrez E-utilities at the US National Center for Biotechnology Information:
  static readonly String server = "http://www.ncbi.nlm.nih.gov/entrez/eutils/";

  public static void Main(String[] args) {
    ShowResult(NcbiProteinTask("P01308").Result);
    ShowResult(NcbiProteinParallelTasks("P01308", "P01315", "P01317").Result);
    // ShowResult(NcbiProteinAsync("P01308").Result);
    // ShowResult(NcbiProteinParallelAsync("P01308", "P01315", "P01317").Result);
    // ShowResult(NcbiSomeProteinAsync("P01308", "P01315", "P01317").Result);
    // ShowResult(NcbiPubmedAsync("molin+s[au]").Result);
    // ShowResult(NcbiPubmedParallelAsync("molin+s[au]", "ingmer+h[au]").Result);
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

  public static Task<String> NcbiEntrezTask(String query) {
    return new WebClient().DownloadDataTaskAsync(new Uri(server + query))
           .ContinueWith((Task<byte[]> task) => 
                         ASCIIEncoding.ASCII.GetString(task.Result));
  }

  public static Task<String> NcbiProteinTask(String id) {
    return NcbiEntrezTask("efetch.fcgi?rettype=fasta&retmode=text&db=protein&id=" + id);
  }

  public static Task<String[]> NcbiProteinParallelTasks(params String[] ids) {
    IEnumerable<Task<String>> tasks = from id in ids select NcbiProteinTask(id);
    return TaskEx.WhenAll(tasks);
  }
 
  public static async Task<String> NcbiEntrezAsync(String query) {
    Console.WriteLine(">>>" + query + ">>>");
    byte[] bytes = await new WebClient().DownloadDataTaskAsync(new Uri(server + query));
    Console.WriteLine("<<<" + query + "<<<");
    return ASCIIEncoding.ASCII.GetString(bytes);
  }

  public static async Task<String> NcbiProteinAsync(String id) {
    return await NcbiEntrezAsync("efetch.fcgi?rettype=fasta&retmode=text&db=protein&id=" + id);
  }

  public static async Task<String[]> NcbiProteinParallelAsync(params String[] ids) {
    IEnumerable<Task<String>> tasks = from id in ids select NcbiProteinAsync(id);
    return await TaskEx.WhenAll(tasks);
  }

  public static async Task<String> NcbiSomeProteinAsync(params String[] ids) {
    IEnumerable<Task<String>> tasks = from id in ids select NcbiProteinAsync(id);
    return await TaskEx.WhenAny(tasks).Result;
  }

  public static async Task<String> NcbiPubmedAsync(String term) {
    String search = String.Format("esearch.fcgi?db=Pubmed&retmax=1&usehistory=y&term={0}", term);
    XmlDocument xml = new XmlDocument();
    xml.LoadXml(await NcbiEntrezAsync(search));
    XmlNode node = xml["eSearchResult"]; 
    String fetch = String.Format("retmax=3&db=Pubmed&query_key={0}&WebEnv={1}", 
                                 node["QueryKey"].InnerText, node["WebEnv"].InnerText);
    return await NcbiEntrezAsync("efetch.fcgi?rettype=abstract&retmode=text&" + fetch);
  }

  public static async Task<String[]> NcbiPubmedParallelAsync(params String[] terms) {
    IEnumerable<Task<String>> tasks = from term in terms select NcbiPubmedAsync(term);
    return await TaskEx.WhenAll(tasks);
  }
}
