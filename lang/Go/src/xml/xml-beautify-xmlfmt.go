////////////////////////////////////////////////////////////////////////////
// Porgram: xml-beautify-xmlfmt.go
// Purpose: Go XML Beautify using the xmlfmt package
// Authors: Tong Sun (c) 2016, All rights reserved
// Credits: https://github.com/go-xmlfmt/xmlfmt
////////////////////////////////////////////////////////////////////////////

package main

import "github.com/go-xmlfmt/xmlfmt"

func main() {
	xml1 := `<root><this><is>a</is><test /><message><org><cn>Some org-or-other</cn><ph>Wouldnt you like to know</ph></org><contact><fn>Pat</fn><ln>Califia</ln></contact></message></this></root>`
	x := xmlfmt.FormatXML(xml1, "\t", "  ")
	print(x)
}

/*

$ go run xml-beautify-xmlfmt.go
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

*/
