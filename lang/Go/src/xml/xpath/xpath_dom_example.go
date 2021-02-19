// https://play.golang.org/p/PoxZ0b_zSat

package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/santhosh-tekuri/dom"
)

func main() {
	str := `
<developer xmlns="www.jroller.com/santhosh/">
 <name>Santhosh Kumar Tekuri</name>
 <email>santhosh.tekuri@gmail.com</email>
</developer>
`

	doc, err := dom.Unmarshal(xml.NewDecoder(strings.NewReader(str)))
	if err != nil {
		fmt.Println(err)
		return
	}

	root := doc.RootElement()
	fmt.Printf("rootElement: {%s}%s\n", root.URI, root.Local)
	buf := new(bytes.Buffer)
	if err = dom.Marshal(doc, buf); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("xml:\n%s", buf.String())
}
