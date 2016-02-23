////////////////////////////////////////////////////////////////////////////
// Porgram: lx2_element-names.go
// Purpose: parse xml element names, gokogiri node iteration
// Authors: Tong Sun (c) 2016, All rights reserved
// Credits: Libxml2: Everything You Need in an XML Library (Victor Volkman)
//          http://www.codeguru.com/cpp/data/data-misc/xml/article.php/c14893/Libxml2-Everything-You-Need-in-an-XML-Library.htm
//          XSD Complex Elements
//          http://www.w3schools.com/xml/schema_complex.asp
//          XML File for the Complex Data Example
//          http://www.service-architecture.com/articles/object-oriented-databases/xml_file_for_complex_data.html
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"github.com/moovweb/gokogiri"
	"github.com/moovweb/gokogiri/xml"
)

// https://godoc.org/github.com/moovweb/gokogiri/xml

func printElementNames(aNode xml.Node) {
	for curNode := aNode; curNode != nil; curNode = curNode.NextSibling() {
		if curNode.NodeType() == xml.XML_ELEMENT_NODE {
			fmt.Printf("node type: Element, name: %s\n", curNode.Name())
		} else if curNode.NodeType() == xml.XML_TEXT_NODE {
			fmt.Printf("node type: Text, string: %s\n", curNode.ToUnformattedXml())
		} else {
			fmt.Printf("node type: %v\n", curNode.NodeType())
		}
		printElementNames(curNode.FirstChild())
	}
}

func main() {
	xml := `<?xml version="1.0"?>
<!DOCTYPE doc [
<!ELEMENT doc (src | dest)*>
<!ELEMENT src EMPTY>
<!ELEMENT dest EMPTY>
<!ATTLIST src ref IDREF #IMPLIED>
<!ATTLIST dest id ID #IMPLIED>
]>
<doc>
  <src ref="foo"/>
  <dest id="foo"/>
  <src ref="foo"/>
  <employee>
    <firstname>John</firstname>
    <lastname>Smith</lastname>
    <food type="dessert">Ice cream</food>
    <description>
      Born on <date lang="norwegian">03.03.99</date> ....
    </description>
  </employee>
</doc>
`
	doc, _ := gokogiri.ParseXml([]byte(xml))
	defer doc.Free()

	rootElement := doc.Root()
	printElementNames(rootElement)
}

/*

Output:

node type: Element, name: doc
node type: Text, string:

node type: Element, name: src
node type: Text, string:

node type: Element, name: dest
node type: Text, string:

node type: Element, name: src
node type: Text, string:

node type: Element, name: employee
node type: Text, string:

node type: Element, name: firstname
node type: Text, string: John
node type: Text, string:

node type: Element, name: lastname
node type: Text, string: Smith
node type: Text, string:

node type: Element, name: food
node type: Text, string: Ice cream
node type: Text, string:

node type: Element, name: description
node type: Text, string:
      Born on
node type: Element, name: date
node type: Text, string: 03.03.99
node type: Text, string:  ....

node type: Text, string:

node type: Text, string:


Ref:

Libxml2: Everything You Need in an XML Library
http://www.codeguru.com/cpp/data/data-misc/xml/article.php/c14893

#include <stdio.h>
#include <libxml/parser.h>
#include <libxml/tree.h>

static void print_element_names(xmlNode * a_node)
{
   xmlNode *cur_node = NULL;

   for (cur_node = a_node; cur_node; cur_node =
      cur_node->next) {
      if (cur_node->type == XML_ELEMENT_NODE) {
         printf("node type: Element, name: %s\n",
            cur_node->name);
      }
      print_element_names(cur_node->children);
   }
}

int main(int argc, char **argv)
{
   xmlDoc *doc = NULL;
   xmlNode *root_element = NULL;

   if (argc != 2)  return(1);

   LIBXML_TEST_VERSION    // Macro to check API for match with
                          // the DLL we are using

   /*parse the file and get the DOM/
   if (doc = xmlReadFile(argv[1], NULL, 0)) == NULL){
      printf("error: could not parse file %s\n", argv[1]);
      exit(-1);
      }

   /*Get the root element nod/
   root_element = xmlDocGetRootElement(doc);
   print_element_names(root_element);
   xmlFreeDoc(doc);       // free document
   xmlCleanupParser();    // Free globals
   return 0;
}


*/
