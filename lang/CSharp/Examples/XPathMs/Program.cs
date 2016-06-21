using System.Xml.XPath; 
using System.Xml; 
using System;
using System.IO; 

class XPathQuery{

public static string PrintError(Exception e, string errStr){

  if(e == null) 
    return errStr; 
  else
    return PrintError(e.InnerException, errStr + e.Message ); 
} 

 public static void Main(string[] args){

   if((args.Length == 0) || (args.Length % 2)!= 0){
     Console.WriteLine("Usage: xpathquery source query <zero or more prefix and namespace pairs>");
      return; 
   }
   
   try{
     
     //Load the file.
     XmlDocument doc = new XmlDocument(); 
     doc.Load(args[0]); 

     //create prefix<->namespace mappings (if any) 
     XmlNamespaceManager  nsMgr = new XmlNamespaceManager(doc.NameTable);

     for(int i=2; i < args.Length; i+= 2)
       nsMgr.AddNamespace(args[i], args[i + 1]); 

     //Query the document 
     XmlNodeList nodes = doc.SelectNodes(args[1], nsMgr); 

     //print output 
     foreach(XmlNode node in nodes)
       Console.WriteLine(node.OuterXml + "\n\n");

   }catch(XmlException xmle){
     Console.WriteLine("ERROR: XML Parse error occured because " + 
PrintError(xmle, null));
   }catch(FileNotFoundException fnfe){
     Console.WriteLine("ERROR: " + PrintError(fnfe, null));
   }catch(XPathException xpath){
     Console.WriteLine("ERROR: The following error occured while querying the document: " 
             + PrintError(xpath, null));
   }catch(Exception e){
     Console.WriteLine("UNEXPECTED ERROR" + PrintError(e, null));
   }
 }
}