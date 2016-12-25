////////////////////////////////////////////////////////////////////////////
// Porgram: jsonfile.go
// Purpose: GO jsonfile package demo
// Authors: Tong Sun (c) 2016, All rights reserved
// Credits: http://www.dotnetperls.com/json-go
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"

	// json http stream
	jhs "github.com/go-jsonfile/jsonfile/ext"
)

func main() {
	var data interface{}
	jhs.GetJSON("http://headers.jsontest.com/", &data)

	fmt.Printf("%+v\n", data)
}

/*


 */
