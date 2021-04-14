using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Xml;
using System.Xml.XPath;
using Dayforce.Common;
using System.IO;
using System.IO.Compression;

using ICSharpCode.SharpZipLib;
using ICSharpCode.SharpZipLib.Zip.Compression.Streams;
using ICSharpCode.SharpZipLib.Zip.Compression;

namespace ConsoleApplication1
{
    class XPath2
    {
        static void Main(string[] args)
        {
            TestXPathRequest();
            TestXPathReqPayloads();
            TestXPathRespond();

            // Keep the console window open in debug mode.
            Console.WriteLine("\nPress any key to exit.");
            Console.Read();
        }


        static void TestXPathRequest()
        {
            Console.WriteLine("== XML Request Processing");

            XmlDocument doc = new XmlDocument();
            doc.Load(@"..\..\Request.xml");
            XmlNode root = doc.DocumentElement;

            string xpathExpression;

            // Select and display binary 
            xpathExpression = "//StringHttpBody";

            XmlNode compressedNode = root.SelectSingleNode(xpathExpression, null);
            Console.WriteLine(compressedNode.InnerText);
            //Console.Read();

            byte[] bytes = Convert.FromBase64String(compressedNode.InnerXml);
            string xmlData = System.Text.Encoding.Unicode.GetString(bytes);

            Console.WriteLine(xmlData);
            Console.Read();
        }

        static void TestXPathReqPayloads()
        {
            Console.WriteLine("\n== XML Request Payloads Processing");

            XmlDocument doc = new XmlDocument();
            doc.Load(@"..\..\ReqPayloads.xml");
            XmlNode root = doc.DocumentElement;

            XmlNode binaryNode = root.SelectSingleNode(@"//*[local-name() = 'base64Binary']"); // Found
            binaryNode = root.SelectSingleNode("//base64Binary");       // No found => null

            // Add the namespace.
            XmlNamespaceManager nsmgr = new XmlNamespaceManager(doc.NameTable);
            nsmgr.AddNamespace("d4p1", "http://Dayforce/Data/CoreService");
            nsmgr.AddNamespace("d5p1", "http://schemas.microsoft.com/2003/10/Serialization/Arrays");

            string xpathExpression;

            // Select and display binary 
            xpathExpression = "//d4p1:Payloads/d5p1:base64Binary";

            XmlNode compressedNode = root.SelectSingleNode(xpathExpression, nsmgr);
            Console.WriteLine(compressedNode.InnerText);
            //Console.Read();

            //(2)===========================================================================================================

            string xmlData = Compressor.CompressedBinaryToXml(compressedNode.InnerXml);

            Console.WriteLine(xmlData);
            Console.Read();

            compressedNode.InnerXml = xmlData;

            byte[] encodedXml = System.Text.Encoding.UTF8.GetBytes(root.InnerXml);
            Console.WriteLine(Convert.ToBase64String(encodedXml, Base64FormattingOptions.None));

            doc.Save(Console.Error);
        }

        static void TestXPathRespond()
        {

            // Output text to the screen.
            Console.WriteLine("\n== XML Respond Processing");

            // Get the current directory.
            string path = Directory.GetCurrentDirectory();
            //Console.WriteLine("The current directory is {0}", path);

            XmlDocument doc = new XmlDocument();
            doc.Load(@"..\..\CompressedXML2.xml");
            XmlNode root = doc.DocumentElement;

            //(1)===========================================================================================================

            /*
   <s:Envelope xmlns:s="http://schemas.xmlsoap.org/soap/envelope/">
     <s:Body>
       <Execute xmlns="http://Dayforce/Services/CoreService">
         <b xmlns:d4p1="http://Dayforce/Data/CoreService" xmlns:i="http://www.w3.org/2001/XMLSchema-instance">
           <d4p1:SessionTicket>{{sessionId}}</d4p1:SessionTicket>
           <d4p1:Infos>
             <d4p1:CoreServiceRequestInfo>
               <d4p1:RequestName>Get</d4p1:RequestName>
             </d4p1:CoreServiceRequestInfo>
           </d4p1:Infos>
           <d4p1:Payloads xmlns:d5p1="http://schemas.microsoft.com/2003/10/Serialization/Arrays">
             <d5p1:base64Binary>AQJyAAAALco7DsIwEAXAHomCQ1BQ4A3QcYE0JEI4Be3KPLCV+MN6wdenYerZ9EgQXnroDe8Pqm69ajkTWc9SplzMCCUL+QaHSteF9Zkl7v6rtWbayWR50bHrDnQfLtZ5RN6HVJWTw2qA+vwYOWJtUWvIaQpuhv4ABcJLCoAgFABAJCKqdceQNPNTSwlq0SUsFB6JVrjpKN42ZjRqihZBqavcrTZt1t87hGuxDgIkiEHXuX/Mbb43es9ndtpDOHVgyqXBo3QST2oieGCGUUIUG5VAPw==</d5p1:base64Binary>
           </d4p1:Payloads>
         </b>
       </Execute>
     </s:Body>
   </s:Envelope>
            */
            // Add the namespace.
            XmlNamespaceManager nsmgr = new XmlNamespaceManager(doc.NameTable);
            nsmgr.AddNamespace("d4p1", "http://Dayforce/Data/CoreService");
            nsmgr.AddNamespace("d5p1", "http://schemas.microsoft.com/2003/10/Serialization/Arrays");

            string xpathExpression;
            xpathExpression = "//d4p1:RequestName";

            XmlNode xnRecord = root.SelectSingleNode(xpathExpression, nsmgr);
            Console.WriteLine(xnRecord.InnerXml);

            // Select and display binary 
            xpathExpression = "//d4p1:Payloads/d5p1:base64Binary";

            XmlNode compressedNode = root.SelectSingleNode(xpathExpression, nsmgr);
            Console.WriteLine(compressedNode.InnerText);
            //Console.Read();

            //(2)===========================================================================================================

            string xmlData = Compressor.CompressedBinaryToXml(compressedNode.InnerXml);

            Console.WriteLine(xmlData);
            //Console.Read();
        }

    }
}

