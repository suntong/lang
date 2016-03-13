////////////////////////////////////////////////////////////////////////////
// Porgram: BuildTimeVar.go
// Purpose: Variable at build-time
// Authors: Tong Sun (c) 2016, All rights reserved
// Credits: https://blog.cloudflare.com/setting-go-variables-at-compile-time/
////////////////////////////////////////////////////////////////////////////

/*

Q: Having a constant of build-time would ease my attempt to keep track of my
version-ing. Many of tools I built myself don't following a rigid versioning
scheme, as I use the `rcs` auto-version-increase before, so it'd be good to
know when an executable was built. And a constant of build-time would be the
best answer.

A: The answer is

Setting Go variables from the outside
https://blog.cloudflare.com/setting-go-variables-at-compile-time/

*/

package main

import "fmt"

var date = "2016-03-10"

func main() {
	fmt.Printf("Hello World on %s.\n", date)
}

/*

$ go run BuildTimeVar.go
Hello World on 2016-03-10.

$ go run -ldflags="-X main.date=`date -I`" BuildTimeVar.go
Hello World on 2016-03-12.

$ go run -ldflags="-X main.date=`date -Iseconds`" BuildTimeVar.go
Hello World on 2016-03-12T22:28:31-0500.

*/
