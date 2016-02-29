////////////////////////////////////////////////////////////////////////////
// Porgram: lx2_xpath-search1.go
// Purpose: gokogiri xpath demo
// Authors: Tong Sun (c) 2016, All rights reserved
// Credits: Aleksey Sanin and Daniel Veillard
//          http://veillard.com/XML/examples/xpath2.c
//          XML File for the Complex Data Example
//          http://www.service-architecture.com/articles/object-oriented-databases/xml_file_for_complex_data.html
////////////////////////////////////////////////////////////////////////////

/*

Goal:

Load a document, locate subelements with XPath, modify said elements and save the resulting document.

*/

package main

import (
	"fmt"
	"os"
	//"io/ioutil"

	"github.com/moovweb/gokogiri"
	//"github.com/moovweb/gokogiri/xpath"
)

func main() {
	// content, _ := ioutil.ReadFile("index.html")
	//doc, _ := gokogiri.ParseHtml(content)
	content := `<?xml version="1.0"?>
<?xml-stylesheet href="catalog.xsl" type="text/xsl"?>
<!DOCTYPE catalog SYSTEM "catalog.dtd">
<catalog>
   <product description="Cardigan Sweater" product_image="cardigan.jpg">
      <catalog_item gender="Men's">
         <item_number>QWZ5671</item_number>
         <price>39.95</price>
         <size description="Medium">
            <color_swatch image="red_cardigan.jpg">Red</color_swatch>
            <color_swatch image="burgundy_cardigan.jpg">Burgundy</color_swatch>
         </size>
         <size description="Large">
            <color_swatch image="red_cardigan.jpg">Red</color_swatch>
            <color_swatch image="burgundy_cardigan.jpg">Burgundy</color_swatch>
         </size>
      </catalog_item>
      <catalog_item gender="Women's">
         <item_number>RRX9856</item_number>
         <price>42.50</price>
         <size description="Small">
            <color_swatch image="red_cardigan.jpg">Red</color_swatch>
            <color_swatch image="navy_cardigan.jpg">Navy</color_swatch>
            <color_swatch image="burgundy_cardigan.jpg">Burgundy</color_swatch>
         </size>
         <size description="Medium">
            <color_swatch image="red_cardigan.jpg">Red</color_swatch>
            <color_swatch image="navy_cardigan.jpg">Navy</color_swatch>
            <color_swatch image="burgundy_cardigan.jpg">Burgundy</color_swatch>
            <color_swatch image="black_cardigan.jpg">Black</color_swatch>
         </size>
         <size description="Large">
            <color_swatch image="navy_cardigan.jpg">Navy</color_swatch>
            <color_swatch image="black_cardigan.jpg">Black</color_swatch>
         </size>
         <size description="Extra Large">
            <color_swatch image="burgundy_cardigan.jpg">Burgundy</color_swatch>
            <color_swatch image="black_cardigan.jpg">Black</color_swatch>
         </size>
      </catalog_item>
   </product>
</catalog>`
	doc, _ := gokogiri.ParseXml([]byte(content))
	defer doc.Free()

	if len(os.Args) <= 1 {
		fmt.Println(doc.String())
	} else if len(os.Args) == 3 {
		example4(os.Args[1], os.Args[2])
	} else {
		fmt.Println("Usage: lx2_xpath-search1 xpath-expr new-value")
	}
}

func example4(doc *xml.XmlDocument, xpathExpr, value string) {
	/* Create xpath evaluation context */
	xpathCtx := doc.DocXPathCtx()
	/* Evaluate xpath expression */
	xpathObj := xpathCtx.EvalXPath(xpathExpr, nil) // xpath.ResolveVariable
	/* update selected nodes */
	updateXpathNodes(xpathObj.ResultAsNodeset(), value)
}

func updateXpathNodes() {
}

/*

 */
