using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Xml;
using System.Xml.XPath;
using Dayforce.Common;
using System.IO;

namespace ConsoleApplication1
{
    class XPath
    {
        static void Main(string[] args)
        {
         // Output text to the screen.
         Console.WriteLine("XML Processing");

         // Get the current directory.
         string path = Directory.GetCurrentDirectory();
         Console.WriteLine("The current directory is {0}", path);

         XmlDocument doc = new XmlDocument();
         doc.Load(@"..\..\CompressedXML.xml");
         XmlNode root = doc.DocumentElement;

        //(1)===========================================================================================================

         /*
         <s:Envelope xmlns:s="http://schemas.xmlsoap.org/soap/envelope/">
            <s:Body>
                <GeneralGetCompressedResponse xmlns="http://Dayforce/Services">
                    <GeneralGetCompressedResult>AQLGAAAARYy9TsMwFIUHJErVChh4BQYG4gIbG6oVyRINVV0E6yU+aQKObV0brL49HlBYjs6n83P5xEzHl07WG8RIB1z3KYVHIXRPHPY+VA2S0OCfoUUUW0up8zze/LVyzlV+qDwfxP1qdSfeN8+67THS7eBiItdiPj2fre0Al5RZrBmUYGSRxRQrczH5/TEUXsp6hxi8iwVO3GCvZP3m+auzPjfeoOxJmbmKEhblb1YcU5dOVdyBzFLDGfBrBCtz/g8NjZjp749PtOkXvdl9aBtlHAfwXGI629C1NX3ZK31GJ0VWbRHF6nAuZ7ZJX7Z0TR1Fy3JJnpmz6V12uaS23eLLTGMWCnGyztSXhk6nVeeKTnDVscJAfGm3IOhQhJZBsS5jtHagMt283BO69KcyoeT5J/nv93m+v3ue391xLJOnMzC8nl3G5kZmdazhyEg18+DoW+duZ/P7Dl3WaNiCsI4tKjb2MmxJhNGxZUF2ZZBdHWLX9h33aDVseaxgG+dzYrkKmUVJlF1d7LpYjckh86KAduK9Pl7CzoeQxY05L0YS9vO4E8kujDhZ5hwu7ER7RKmDuelH7W2/niL+y9C/rs2+b3ntsFhK/Cj0r1Hwx14dMl86qvoR6F+h4G/+pnZyhvgHoZ9c8INZ8xPDGx5vJH4Y+jMU/KldF55oIP5L0J+m4BfeZQ7WEz8E/YsU/Jrl75fXEb8X+pMU/OhcHWsmfhD6P9HoP79a3kj8F6H/A4XzV1NhbF5H/GehP04hv+YvPmAkfgD6X1HwUXDzA2l/P/S/oOBb+noq9g+p/j7on6Ux/wdOFnYTvwf6n1LIH28MF3uJ3w39Tyjk9/yo2yQRvwv6H1HwE/duHBWJ/wz0T9C4/60pGROI3wn9D2jsf8elM2nfD/1hCv7IwPqpHcT3Qn+Qxv677cPrd8ZVfy/xtQt+f6YfOWJQ9PwGzoGrkFXyObtQY8q2KsVTMPbKqAFzfozEPchk92LBgZGL8yI7xgLiPB5J9GPnPSjDtl0+aygjNg+zP7+o971NulT2PJMgVKFmXzePlxh7tuAMyie0HdJCdmnNdOVRPaE5SLdnl7bVMhxD6DZI22gc9q8vvn1tUPVbod9Cw58rufs34rdAv47Gw4Y/cWqO+E3Qf4SCb7v6VHkb8bdBfz2FYYO+dV3dRXwz9NfQeNi1tleuIn4t9P9kKPjjLfrn3lT9CuifpuDPrjXv+/wN1V8J/SEm+9d/6v5AT9pf8Y+bHQU/PlE6N0r8Mui/TqH/8VDzRNovhX6Mgm8ZTO4+SfwS6PfT8C9oj39M/GLov0LBR4OfhdO+EfpRCr7N+l3RYeLfAf0+Cn7ivoc7o8Qvgv5BGvPn9+/9ab8Q+iEa87d/etN7A6q/DPoihfmTyDl9Ypj4OdB307j+Y39E0r4e+jwFf8Qz9vMh4mvhy8ZuxQ9l8WXDUx1xRol9IwmytzJZfez23Ij3zMdUeiYJYq/4H7GrU7G3cx2KIDhRIyfxMudGzTIn+7y3SD0rNSTfJfR5QKu/GfYvOsXOedTk7OAFBc1lJZFzOjivnDnA5yej6XLn1HIHFle7eYAqr6gbqMhk59w+bxWq55QOSbyQ6mXuVs4vKiFw5koLA6b8dOmJ/y4dvEXpnK28W8ZS5syXj7XNH1Drxv61AwuFYQf0j2G3W1z0rYBfPk5KtYMtFF6VGf6d8+oKjSbBiSWvqGyjeuVfVi7SlqXeQUfzLC0vqEt4Ei6hMLNJgQ3qJi4wsewOq1VZgGnnlu3kEC1lgDx97MsmwrdC3pAtnvkb
                    </GeneralGetCompressedResult>
                </GeneralGetCompressedResponse>
         </s:Body>
         </s:Envelope>
         */
         // Add the namespace.
         XmlNamespaceManager nsmgr = new XmlNamespaceManager(doc.NameTable);
         nsmgr.AddNamespace("s","http://schemas.xmlsoap.org/soap/envelope/");
         nsmgr.AddNamespace("df", "http://Dayforce/Services");

         nsmgr.AddNamespace("i", "http://www.w3.org/2001/XMLSchema-instance");
         nsmgr.AddNamespace("df1", "http://SharpTop.Net/Services/Platform");

         // Select and display the first node  GeneralGetCompressedResult
         
         string xpathExpression;
         xpathExpression = "//s:Envelope/s:Body/df:GeneralGetCompressedResponse/df:GeneralGetCompressedResult";


         XmlNode compressedNode = root.SelectSingleNode(xpathExpression, nsmgr);


         Console.WriteLine(compressedNode.InnerText);
         Console.ReadKey();

         //(2)===========================================================================================================

         //Compressor compressor = new Compressor();
         //Console.WriteLine(System.Convert.FromBase64String(compressedNode.InnerText));

  
         //doc.LoadXml(Compressor.CompressedBinaryToXml(compressedNode.InnerText));

         string xmlData = Compressor.CompressedBinaryToXml(compressedNode.InnerXml);

         Console.WriteLine(xmlData);
         Console.ReadKey();

         doc.Load(new StringReader(xmlData));



         //(3)=======================================================================================================

        //XmlDocument doc = new XmlDocument();
         //doc.Load("D:\\Project Scripts\\DecompressedXML.xml");
         //root = doc.DocumentElement;
         //Console.WriteLine(root.InnerXml);
         //Console.ReadKey();
   
         //xpathExpression = "//s:Envelope/df1:ArrayOfDFMessage/df1:DFMessage/df1:DFMessageId";
         xpathExpression = "//df1:ArrayOfDFMessage/df1:DFMessage/df1:DFMessageId";

         XmlNode valueNode = doc.SelectSingleNode(xpathExpression, nsmgr);
         //XmlNode valueNode = root.SelectSingleNode(xpathExpression, nsmgr);
         Console.WriteLine(valueNode.InnerText);
         Console.ReadKey();


        }
    }
}

