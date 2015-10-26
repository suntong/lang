////////////////////////////////////////////////////////////////////////////
// Porgram: utf16
// Purpose: Convert base64 encoded utf16 string to ascii string
// Authors: Tong Sun (c) 2015, All rights reserved
// Credits: https://golang.org/pkg/encoding/base64/#example_Encoding_DecodeString
//          http://www.pocketsoap.com/pocketsoap/docs/master/headers.htm
////////////////////////////////////////////////////////////////////////////

package main

import (
	"encoding/base64"
	"fmt"
)

func main() {

	uEnc :=
		`PABTADoARQBuAHYAZQBsAG8AcABlAAoAIAAgACAAIAAgACAAIAAgAHgAbQBsAG4AcwA6AFMAPQAn
AGgAdAB0AHAAOgAvAC8AcwBjAGgAZQBtAGEAcwAuAHgAbQBsAHMAbwBhAHAALgBvAHIAZwAvAHMA
bwBhAHAALwBlAG4AdgBlAGwAbwBwAGUALwAnAD4ACgA8AFMAOgBCAG8AZAB5AD4APABkADoARABv
AFMAdAB1AGYAZgA+ADwALwBkADoARABvAFMAdAB1AGYAZgA+AAoAPAAvAFMAOgBCAG8AZAB5AD4A
PAAvAFMAOgBFAG4AdgBlAGwAbwBwAGUAPgA=`

	uDec, _ := base64.StdEncoding.DecodeString(uEnc)
	fmt.Println(uDec)
	fmt.Println(string(uDec))

}

/*

<S:Envelope
        xmlns:S='http://schemas.xmlsoap.org/soap/envelope/'>
<S:Body><d:DoStuff></d:DoStuff>
</S:Body></S:Envelope>

*/
