////////////////////////////////////////////////////////////////////////////
// Porgram: Defer1.go
// Purpose: Demo a proper way to defer
// Authors: Tong Sun (c) 2016, All rights reserved
// Credits: Andy Schweig http://stackoverflow.com/questions/40953241/
////////////////////////////////////////////////////////////////////////////

/*

Q: is it possible to put defer in sub func?

I.e., for line 12~16 of func test1() in https://play.golang.org/p/evabhcjvNs

Is there any possibility to put them in sub function?

The reason I'm asking is that, for line 12~16 of func test1(), my actual
code is to restore the variable from persistent data, then use defer to save
it when test1() is done. However, there are cases that the whole
restore/save is not necessary, so I'm thinking a better way to control it.

A: You want to know if a function can place a function on the defer stack of
the caller. The answer to that is no.
One possible solution to this is to have the function that wants to defer
something return that function to the caller and have the caller do the
defer.

*/

package main

import "log"

func main() {
	test1()
	log.Printf("==== \n")
	test2()
}

func test1() {
	r, err := Open("a")
	if err != nil {
		log.Fatalf("error opening 'a'\n")
	}
	defer r.Close()

	r.Use()
}

func test2() {
	r, cleanup := subc()
	defer cleanup()
	r.Use()
}

func subc() (*Resource, func()) {
	r, err := Open("a")
	if err != nil {
		log.Fatalf("error opening 'a'\n")
	}
	return r, func() { r.Close() }
}

type Resource struct {
	name string
}

func Open(name string) (*Resource, error) {
	log.Printf("opening %s\n", name)
	return &Resource{name}, nil
}

func (r *Resource) Use() error {
	log.Printf("using %s\n", r.name)
	return nil
}

func (r *Resource) Close() error {
	log.Printf("closing %s\n", r.name)
	return nil
}

/*

2016/12/03 22:03:21 opening a
2016/12/03 22:03:21 using a
2016/12/03 22:03:21 closing a
2016/12/03 22:03:21 ====
2016/12/03 22:03:21 opening a
2016/12/03 22:03:21 using a
2016/12/03 22:03:21 closing a

Comparing to Defer0.go, the above sequence is correct.

*/
