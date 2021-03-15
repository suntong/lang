// https://play.golang.org/p/5AdXiGVHPGs

package main

import (
	"encoding/xml"
	"os"
)

func main() {
	str := "\"'&<>"
	xml.Escape(os.Stdout, []byte(str))
}

// Output:
// &#34;&#39;&amp;&lt;&gt;
