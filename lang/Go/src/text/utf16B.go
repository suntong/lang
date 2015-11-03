////////////////////////////////////////////////////////////////////////////
// Porgram: utf16
// Purpose: Convert base64 encoded utf16 string to ascii string
// Authors: Tong Sun (c) 2015, All rights reserved
// Credits: https://golang.org/pkg/encoding/base64/#example_Encoding_DecodeString
//          http://www.pocketsoap.com/pocketsoap/docs/master/headers.htm
//          http://play.golang.org/p/cCufRfEQso by Jakob Borg
////////////////////////////////////////////////////////////////////////////

package main

import (
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"unicode/utf16"
)

func main() {

	uEnc :=
		`PABTADoARQBuAHYAZQBsAG8AcABlAAoAIAAgACAAIAAgACAAIAAgAHgAbQBsAG4AcwA6AFMAPQAn
AGgAdAB0AHAAOgAvAC8AcwBjAGgAZQBtAGEAcwAuAHgAbQBsAHMAbwBhAHAALgBvAHIAZwAvAHMA
bwBhAHAALwBlAG4AdgBlAGwAbwBwAGUALwAnAD4ACgA8AFMAOgBCAG8AZAB5AD4APABkADoARABv
AFMAdAB1AGYAZgA+ADwALwBkADoARABvAFMAdAB1AGYAZgA+AAoAPAAvAFMAOgBCAG8AZAB5AD4A
PAAvAFMAOgBFAG4AdgBlAGwAbwBwAGUAPgA=`

	uDec, _ := base64.StdEncoding.DecodeString(uEnc)
	fmt.Printf("%q\n", uDec)

	u16s := make([]uint16, len(uDec)/2)
	for i := range u16s {
		u16s[i] = binary.LittleEndian.Uint16([]byte(uDec[i*2:]))
	}

	str := string(utf16.Decode(u16s))
	fmt.Printf("%q\n", str)
}

/*

<S:Envelope
        xmlns:S='http://schemas.xmlsoap.org/soap/envelope/'>
<S:Body><d:DoStuff></d:DoStuff>
</S:Body></S:Envelope>

*/
