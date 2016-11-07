////////////////////////////////////////////////////////////////////////////
// Porgram: xml-beautify-regexp.go
// Purpose: Go XML Beautify from flat XML using pure regexp
// Authors: Tong Sun (c) 2016, All rights reserved
// Credits: diotalevi http://www.perlmonks.org/?node_id=261292
////////////////////////////////////////////////////////////////////////////

package main

import (
	"regexp"
	"strings"
)

const xml1 = `<root><this><is>a</is><test /><message><org><cn>Some org-or-other</cn><ph>Wouldnt you like to know</ph></org><contact><fn>Pat</fn><ln>Califia</ln></contact></message></this></root>`

const xml2 = `<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ns="http://example.com/ns"><soapenv:Header/><soapenv:Body><ns:request><ns:customer><ns:id>123</ns:id><ns:name type="NCHZ">John Brown</ns:name></ns:customer></ns:request></soapenv:Body></soapenv:Envelope>`

const xml3 = `<Envelope xmlns="http://schemas.xmlsoap.org/soap/envelope/" xmlns:_xmlns="xmlns" _xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" _xmlns:ns="http://example.com/ns"><Header xmlns="http://schemas.xmlsoap.org/soap/envelope/"></Header><Body xmlns="http://schemas.xmlsoap.org/soap/envelope/"><request xmlns="http://example.com/ns"><customer xmlns="http://example.com/ns"><id xmlns="http://example.com/ns">123</id><name xmlns="http://example.com/ns" type="NCHZ">John Brown</name></customer></request></Body></Envelope>`

func main() {
	FormatXML(xml1, "\t", "  ")
	FormatXML(xml2, "x ", " ")
	FormatXML(xml3, "", " ")
}

/*

$ go run xml-beautify-regexp.go
        <root>
          <this>
            <is>
              a</is>
            <test />
              <message>
                <org>
                  <cn>
                    Some org-or-other</cn>
                  <ph>
                    Wouldnt you like to know</ph>
                  </org>
                <contact>
                  <fn>
                    Pat</fn>
                  <ln>
                    Califia</ln>
                  </contact>
                </message>
              </this>
            </root>


x <soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ns="http://example.com/ns">
x  <soapenv:Header/>
x   <soapenv:Body>
x    <ns:request>
x     <ns:customer>
x      <ns:id>
x       123</ns:id>
x      <ns:name type="NCHZ">
x       John Brown</ns:name>
x      </ns:customer>
x     </ns:request>
x    </soapenv:Body>
x   </soapenv:Envelope>
x

<Envelope xmlns="http://schemas.xmlsoap.org/soap/envelope/" xmlns:_xmlns="xmlns" _xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" _xmlns:ns="http://example.com/ns">
 <Header xmlns="http://schemas.xmlsoap.org/soap/envelope/">
  </Header>
 <Body xmlns="http://schemas.xmlsoap.org/soap/envelope/">
  <request xmlns="http://example.com/ns">
   <customer xmlns="http://example.com/ns">
    <id xmlns="http://example.com/ns">
     123</id>
    <name xmlns="http://example.com/ns" type="NCHZ">
     John Brown</name>
    </customer>
   </request>
  </Body>
 </Envelope>


*/

var reg = regexp.MustCompile(`<(/?)([^>]+)(/?)>`)

func FormatXML(xmls, prefix, indent string) {
	src := regexp.MustCompile(`>\s+<`).ReplaceAllString(xmls, "><")

	rf := ReplaceTag(prefix, indent)
	println(prefix + reg.ReplaceAllStringFunc(src, rf))
	println()
}

func ReplaceTag(prefix, indent string) func(string) string {
	indentLevel := 0
	return func(m string) string {
		parts := reg.FindStringSubmatch(m)
		// $3: A <foo/> tag. No alteration to indentation.
		// $1: A closing </foo> tag. Drop one indentation level
		// else: An opening <foo> tag. Increase one indentation level
		if len(parts[3]) == 0 && len(parts[1]) != 0 {
			indentLevel--
		} else {
			indentLevel++
		}
		return "<" + parts[1] + parts[2] + parts[3] + ">" +
			"\r\n" + prefix + strings.Repeat(indent, indentLevel)
	}
}
