////////////////////////////////////////////////////////////////////////////
// Porgram: xml-beautify-string.go
// Purpose: Go XML Beautify from flat XML
// Authors: Tong Sun (c) 2016, All rights reserved
// Credits: as listed below
////////////////////////////////////////////////////////////////////////////

package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const xml1 = `<root><this><is>a</is><test /><message><org><cn>Some org-or-other</cn><ph>Wouldnt you like to know</ph></org><contact><fn>Pat</fn><ln>Califia</ln></contact></message></this></root>`

const xml2 = `<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ns="http://example.com/ns"><soapenv:Header/><soapenv:Body><ns:request><ns:customer><ns:id>123</ns:id><ns:name type="NCHZ">John Brown</ns:name></ns:customer></ns:request></soapenv:Body></soapenv:Envelope>`

const xml3 = `<s:Envelope xmlns:s="http://schemas.xmlsoap.org/soap/envelope/"><s:Body><Execute xmlns="http://MyCorp/Services/CoreService"><bundle xmlns:i="http://www.w3.org/2001/XMLSchema-instance"><Requests><CoreServiceRequest><ReadableRequestName>Get</ReadableRequestName><RequestName>R2V0</RequestName><Payload xmlns:d7p1="http://www.w3.org/2001/XMLSchema" i:type="d7p1:string"><GeneralGetRequest xmlns:i="http://www.w3.org/2001/XMLSchema-instance" xmlns="http://schemas.datacontract.org/2004/07/MyCorp.Data.DataProcessor"><MethodName>GetUserOrgs</MethodName><Parameters><GeneralGetParameter><Name>UserId</Name><Value xmlns:d4p1="http://www.w3.org/2001/XMLSchema" i:type="d4p1:int">9916</Value></GeneralGetParameter></Parameters><SessionTicket>SessionId</SessionTicket></GeneralGetRequest></Payload></CoreServiceRequest></Requests><RoleId>1008</RoleId><SessionTicket>SessionId</SessionTicket></bundle></Execute></s:Body></s:Envelope>`

func main() {

	// cp, Chris Pushbullet
	cpr, err := formatXML([]byte(xml1))
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("%v\n\n", string(cpr))

	// sw, Sam Whited
	swCustomized()
	fmt.Println()
	swStandard(xml1)
	fmt.Println()
	swStandard(flatxml)
	fmt.Println()
	swStandard(xml2)
	fmt.Println()
	swStandard(xml3)
	fmt.Println()
	swSimple()
}

/*

   <Envelope xmlns="http://schemas.xmlsoap.org/soap/envelope/" xmlns:_xmlns="xmlns" _xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" _xmlns:ns="http://example.com/ns">
    <Header xmlns="http://schemas.xmlsoap.org/soap/envelope/"></Header>
    <Body xmlns="http://schemas.xmlsoap.org/soap/envelope/">
     <request xmlns="http://example.com/ns">
      <customer xmlns="http://example.com/ns">
       <id xmlns="http://example.com/ns">123</id>
       <name xmlns="http://example.com/ns" type="NCHZ">John Brown</name>
      </customer>
     </request>
    </Body>
   </Envelope>

   <Envelope xmlns="http://schemas.xmlsoap.org/soap/envelope/" xmlns:_xmlns="xmlns" _xmlns:s="http://schemas.xmlsoap.org/soap/envelope/">
    <Body xmlns="http://schemas.xmlsoap.org/soap/envelope/">
     <Execute xmlns="http://MyCorp/Services/CoreService" xmlns="http://MyCorp/Services/CoreService">
      <bundle xmlns="http://MyCorp/Services/CoreService" _xmlns:i="http://www.w3.org/2001/XMLSchema-instance">
       <Requests xmlns="http://MyCorp/Services/CoreService">
        <CoreServiceRequest xmlns="http://MyCorp/Services/CoreService">
         <ReadableRequestName xmlns="http://MyCorp/Services/CoreService">Get</ReadableRequestName>
         <RequestName xmlns="http://MyCorp/Services/CoreService">R2V0</RequestName>
         <Payload xmlns="http://MyCorp/Services/CoreService" _xmlns:d7p1="http://www.w3.org/2001/XMLSchema" xmlns:XMLSchema-instance="http://www.w3.org/2001/XMLSchema-instance" XMLSchema-instance:type="d7p1:string">
          <GeneralGetRequest xmlns="http://schemas.datacontract.org/2004/07/MyCorp.Data.DataProcessor" _xmlns:i="http://www.w3.org/2001/XMLSchema-instance" xmlns="http://schemas.datacontract.org/2004/07/MyCorp.Data.DataProcessor">
           <MethodName xmlns="http://schemas.datacontract.org/2004/07/MyCorp.Data.DataProcessor">GetUserOrgs</MethodName>
           <Parameters xmlns="http://schemas.datacontract.org/2004/07/MyCorp.Data.DataProcessor">
            <GeneralGetParameter xmlns="http://schemas.datacontract.org/2004/07/MyCorp.Data.DataProcessor">
             <Name xmlns="http://schemas.datacontract.org/2004/07/MyCorp.Data.DataProcessor">UserId</Name>
             <Value xmlns="http://schemas.datacontract.org/2004/07/MyCorp.Data.DataProcessor" _xmlns:d4p1="http://www.w3.org/2001/XMLSchema" XMLSchema-instance:type="d4p1:int">9916</Value>
            </GeneralGetParameter>
           </Parameters>
           <SessionTicket xmlns="http://schemas.datacontract.org/2004/07/MyCorp.Data.DataProcessor">SessionId</SessionTicket>
          </GeneralGetRequest>
         </Payload>
        </CoreServiceRequest>
       </Requests>
       <RoleId xmlns="http://MyCorp/Services/CoreService">1008</RoleId>
       <SessionTicket xmlns="http://MyCorp/Services/CoreService">SessionId</SessionTicket>
      </bundle>
     </Execute>
    </Body>
   </Envelope>

*/

////////////////////////////////////////////////////////////////////////////
// Chris Pushbullet http://stackoverflow.com/a/27141132/2125837

func formatXML(data []byte) ([]byte, error) {
	b := &bytes.Buffer{}
	decoder := xml.NewDecoder(bytes.NewReader(data))
	encoder := xml.NewEncoder(b)
	encoder.Indent("", "  ")
	for {
		token, err := decoder.Token()
		if err == io.EOF {
			encoder.Flush()
			return b.Bytes(), nil
		}
		if err != nil {
			return nil, err
		}
		err = encoder.EncodeToken(token)
		if err != nil {
			return nil, err
		}
	}
}

/*

<root>
  <this>
    <is>a</is>
    <test></test>
    <message>
      <org>
        <cn>Some org-or-other</cn>
        <ph>Wouldnt you like to know</ph>
      </org>
      <contact>
        <fn>Pat</fn>
        <ln>Califia</ln>
      </contact>
    </message>
  </this>
</root>

*/

////////////////////////////////////////////////////////////////////////////
// Sam Whited, https://groups.google.com/forum/#!msg/golang-nuts/lHPOHD-8qio

const flatxml = `<xml><inner><text>Some text</text><text>More text</text><ns1:Element/></inner></xml>`

//==========================================================================

func swCustomized() {
	//println("\nswCustomized")
	d := xml.NewDecoder(strings.NewReader(flatxml))
	e := xml.NewEncoder(os.Stdout)

	nesting := -1
	for {
		t, err := d.Token()
		if err == io.EOF {
			e.Flush()
			return
		}
		if err != nil {
			log.Fatal(err)
		}
		if _, ok := t.(xml.StartElement); ok {
			nesting++
		}
		if nesting > 0 {
			e.EncodeToken(xml.CharData(bytes.Repeat([]byte("  "), nesting)))
		}
		e.EncodeToken(t)
		e.EncodeToken(xml.CharData([]byte{'\n'}))
		if _, ok := t.(xml.EndElement); ok {
			nesting--
		}
		e.Flush()
	}
}

/*

<xml>
  <inner>
    <text>
    Some text
    </text>
    <text>
    More text
    </text>
    <Element xmlns="ns1">
    </Element>
  </inner>
</xml>

*/

//==========================================================================

func swStandard(flatxml string) {
	buf := new(bytes.Buffer)
	d := xml.NewDecoder(strings.NewReader(flatxml))
	e := xml.NewEncoder(buf)
	e.Indent("\t", " ")

tokenize:
	for {
		tok, err := d.Token()
		switch {
		case err == io.EOF:
			e.Flush()
			break tokenize
		case err != nil:
			log.Fatal(err)
		}
		e.EncodeToken(tok)
	}

	newxml := buf.String()
	fmt.Println(newxml)
}

/*

	<xml>
	 <inner>
	  <text>Some text</text>
	  <text>More text</text>
	  <Element xmlns="ns1"></Element>
	 </inner>
	</xml>

*/

//==========================================================================

// "in this example you could of course just use MarshalIndent"

func swSimple() {
	d := xml.NewDecoder(strings.NewReader(flatxml))
	output, err := xml.MarshalIndent(d, "xx ", " ")
	if err != nil {
		//fmt.Printf("error: %v\n", err)
		// error: xml: unsupported type: map[string]string
	}

	os.Stdout.Write(output)

}
