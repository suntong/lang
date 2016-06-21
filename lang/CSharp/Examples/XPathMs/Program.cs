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
       Console.WriteLine("Usage: xpathms query source <zero or more prefix and namespace pairs>");
      return; 
   }
   
   try{
     
     //Load the file.
     XmlDocument doc = new XmlDocument(); 
     doc.Load(args[1]); 

     //create prefix<->namespace mappings (if any) 
     XmlNamespaceManager  nsMgr = new XmlNamespaceManager(doc.NameTable);

     for(int i=2; i < args.Length; i+= 2)
       nsMgr.AddNamespace(args[i], args[i + 1]); 

     //Query the document 
     XmlNodeList nodes = doc.SelectNodes(args[0], nsMgr); 

     //print output 
     foreach(XmlNode node in nodes)
       Console.WriteLine(node.OuterXml + "\n");

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

/*

Output:

Without Namespaces

> XPathMs.exe /bookstore/book/title bookstore0.xml
   <title>The Autobiography of Benjamin Franklin</title>
   <title>The Confidence Man</title>

> XPathMs.exe //@genre bookstore0.xml
   genre="autobiography"
   genre="novel"

> XPathMs.exe "//title[(../author/first-name = 'Herman')]" bookstore0.xml
<title>The Confidence Man</title> 

With Namespaces

> XPathMs.exe //@genre bookstore1.xml
   genre="autobiography"
   genre="novel"

> XPathMs.exe "//title[(../author/first-name = 'Herman')]" bookstore1.xml
empty

> XPathMs.exe /b:bookstore/b:book/b:title bookstore1.xml b urn:xmlns:25hoursaday-com:bookstore
   <title xmlns="urn:xmlns:25hoursaday-com:bookstore">The Autobiography of Benjamin Franklin</title>
   <bk:title xmlns:bk="urn:xmlns:25hoursaday-com:bookstore">The Confidence Man</bk:title>
  
> XPathMs.exe //@b:genre bookstore1.xml b urn:xmlns:25hoursaday-com:bookstore
  bk:genre="fiction"

> XPathMs.exe "//bk:title[(../bk:author/bk:first-name = 'Herman')]" bookstore1.xml bk urn:xmlns:25hoursaday-com:bookstore
<bk:title xmlns:bk="urn:xmlns:25hoursaday-com:bookstore">The Confidence Man</bk:title>

> XPathMs.exe "//*:title" bookstore1.xml
ERROR: The following error occured while querying the document: '//*:title' has an invalid token.
 
*/
