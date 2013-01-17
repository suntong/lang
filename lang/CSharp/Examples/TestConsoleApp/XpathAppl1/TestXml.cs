using System;
using System.IO;        // for StringWriter
using System.Text;

using System.Xml;
using System.Xml.XPath; // For XPathNavigator, XPathDocument & XPathNodeIterator 

namespace XpathAppl1
{
    class TestXml
    {
        static void Main(string[] args)
        {
            // TestCNN1();            TestCNN2();
            TestXpathFile();
            TestXmlNamespace();
            TestXpaths();

            // Keep the console window open in debug mode.
            Console.WriteLine("\nPress any key to exit.");
            Console.ReadKey();
        }

        // From: http://csharp.net-tutorials.com/xml/using-xpath-with-the-xmldocument-class/

        /// <summary>
        /// Test Xml from CNN, Test 1
        /// </summary>
        static void TestCNN1()
        {
            Console.WriteLine("== Test CNN 1");
            XmlDocument xmlDoc = new XmlDocument();
            xmlDoc.Load("http://rss.cnn.com/rss/edition_world.rss");
            XmlNode titleNode = xmlDoc.SelectSingleNode("//rss/channel/title");
            if (titleNode != null)
                Console.WriteLine(titleNode.InnerText);
        }

        /// <summary>
        /// Test Xml from CNN, Test 2
        /// </summary>
        static void TestCNN2()
        {
            Console.WriteLine("== Test CNN 2");

            XmlDocument xmlDoc = new XmlDocument();
            xmlDoc.Load("http://rss.cnn.com/rss/edition_world.rss");
            XmlNodeList itemNodes = xmlDoc.SelectNodes("//rss/channel/item");
            foreach (XmlNode itemNode in itemNodes)
            {
                XmlNode titleNode = itemNode.SelectSingleNode("title");
                XmlNode dateNode = itemNode.SelectSingleNode("pubDate");
                if ((titleNode != null) && (dateNode != null))
                    Console.WriteLine(dateNode.InnerText + ": " + titleNode.InnerText);
            }

        }

        // From: http://support.microsoft.com/kb/308333

        /// <summary>
        /// Test Xml file handling
        /// </summary>

        static void TestXpathFile()
        {
            Console.WriteLine("== Test xpath with file");

            XPathNavigator nav;
            XPathDocument docNav;
            XPathNodeIterator NodeIter;
            String strExpression;

            // Open the XML.
            docNav = new XPathDocument(@"d:\books1.xml");

            // Create a navigator to query with XPath.
            nav = docNav.CreateNavigator();

            // Find the average cost of a book.
            // This expression uses standard XPath syntax.
            strExpression = "sum(/bookstore/book/price) div count(/bookstore/book/price)";

            // Use the Evaluate method to return the evaluated expression.
            Console.WriteLine("The average cost of the books are {0}", nav.Evaluate(strExpression));

            // Find the title of the books that are greater then $10.00.
            strExpression = "/bookstore/book/title[../price>10.00]";

            // Select the node and place the results in an iterator.
            NodeIter = nav.Select(strExpression);

            Console.WriteLine("List of expensive books:");
            //Iterate through the results showing the element value.
            while (NodeIter.MoveNext())
            {
                Console.WriteLine("Book Title: {0}", NodeIter.Current.Value);
            };
        }

        /// <summary>
        /// Test Xml Namespace
        /// </summary>

        static void TestXmlNamespace()
        {
            // From: http://msdn.microsoft.com/en-us/library/6k4x060d.aspx#Y0
            // Select a node set using the Select method with the XmlNamespaceManager object specified to resolve namespace prefixes in the XPath expression.

            Console.WriteLine("== Test Xml Namespace");

            // Open the XML.
            XPathDocument document = new XPathDocument(@"d:\books2.xml");
            XPathNavigator navigator = document.CreateNavigator();
            XmlNamespaceManager manager = new XmlNamespaceManager(navigator.NameTable);
            manager.AddNamespace("bk", "http://www.contoso.com/books");

            XPathNodeIterator nodes = navigator.Select("/bk:bookstore/bk:book/bk:price", manager);
            while (nodes.MoveNext())
            {
                Console.WriteLine(nodes.Current.Value);
            }

            nodes = navigator.Select("/bk:bookstore/bk:book/bk:price", manager);
            // Move to the first node bk:price node
            if (nodes.MoveNext())
            {
                // now nodes.Current points to the first selected node
                XPathNavigator nodesNavigator = nodes.Current;

                //select all the descendants of the current price node
                XPathNodeIterator nodesText =
                   nodesNavigator.SelectDescendants(XPathNodeType.Text, false);

                while (nodesText.MoveNext())
                {
                    Console.WriteLine(nodesText.Current.Value);
                }
            }
        }


        /// <summary>
        /// Test various xpath expressionss
        /// </summary>

        static void TestXpaths()
        {
            Console.WriteLine("== Test various xpath expressionss");

            // From: http://www.csharp-examples.net/xpath-top-xml-nodes/
            // get top 2 nodes <book> nodes use XPath expression
            // http://www.java2s.com/Code/CSharp/XML/FindElementswithanXPathSearch.htm

            String xmlStr = "<Names>" +
                "<Name>James</Name>" +
                "<Name>John</Name>" +
                "<Name>Robert</Name>" +
                "<Name>Michael</Name>" +
                "<Name>William</Name>" +
                "<Name>David</Name>" +
                "<Name>Richard</Name>" +
                "</Names>";

            XmlDocument xml = new XmlDocument();
            xml.LoadXml(xmlStr);

            XmlNodeList xnList = xml.SelectNodes("/Names/Name[position() <= 3]");
            foreach (XmlNode xn in xnList)
            {
                Console.WriteLine(xn.InnerText);
            }

            Console.WriteLine("== Test Write XML To String");

            // From: http://www.fincher.org/tips/Languages/csharp.shtml
            //read an xml file and then write to a string 
            //(no error checking)
            // http://www.java2s.com/Code/CSharp/XML/XPathQueryDemo.htm

            // XmlDocument xmlDocument = new XmlDocument();
            // xmlDocument.Load(args[0]);
            StringWriter stringWriter = new StringWriter();
            XmlTextWriter xmlTextWriter = new XmlTextWriter(stringWriter);
            xmlTextWriter.Formatting = Formatting.Indented;

            xml.WriteTo(xmlTextWriter); // xmlDocument
            xmlTextWriter.Flush();
            Console.WriteLine(stringWriter.ToString());


            Console.WriteLine("== Test more xpath expressionss");

            // From: http://mydotnet.wordpress.com/2008/05/29/worlds-smallest-xml-xpath-tutorial/

            XmlDocument xmlDocument = new XmlDocument();
            xmlDocument.Load(@"d:\books2.xml");
            XmlElement root = xmlDocument.DocumentElement;

            // define name space manager
            XmlNameTable xmlNameTable = new NameTable();
            XmlNamespaceManager manager = new XmlNamespaceManager(xmlNameTable);
            manager.AddNamespace("bk", "http://www.contoso.com/books");

            // select the first Node under the root node ("bookstore") which matches the name "book"
            XmlNode xnRecord = root.SelectSingleNode("bk:book", manager);
            Console.WriteLine(xnRecord.InnerText);

            // Searching XMLNode by Attribute value: book genre="novel"
            xnRecord = root.SelectSingleNode("bk:book[@genre='novel']", manager);
            Console.WriteLine(xnRecord.InnerText);

            // Searching XMLNode by ChildNode’s inner text: <price>9.99</price>
            xnRecord = root.SelectSingleNode("bk:book[bk:price='9.99']", manager);
            Console.WriteLine(xnRecord.InnerXml);

            // Searching XMLNodeList
            XmlNodeList xnlRecords = root.SelectNodes("bk:book", manager);

            // Sorting
            Console.WriteLine("== Test Xml Sorting");

            XPathNavigator navigator = xml.CreateNavigator();
            //XPathExpression selectExpression = navigator.Compile(navigateNode);
            //selectExpression.SetContext(manager);
            //selectExpression.AddSort("@EmpId", XmlSortOrder.Ascending, XmlCaseOrder.None, "", XmlDataType.Number);
            XPathNodeIterator nodeIterator = navigator.Select("/Names/Name");
            while (nodeIterator.MoveNext())
            {
                XmlElement xnm = (XmlElement)((IHasXmlNode)nodeIterator.Current).GetNode();
                //Do other required operations here ??
                Console.WriteLine(xnm.InnerText);
            }
        }


    }
}
