////////////////////////////////////////////////////////////////////////////
// Porgram: Defer1.go
// Purpose: Demo a proper way to defer
// Authors: Tong Sun (c) 2016, All rights reserved
// Credits: Bleve
// http://www.blevesearch.com/news/Deferred-Cleanup,-Checking-Errors,-and-Potential-Problems/
////////////////////////////////////////////////////////////////////////////

package main

import (
	"io"
	"log"
)

func main() {
	test1()
	log.Printf("==== \n")
	test2()
	log.Printf("==== \n")
	test3()
	log.Printf("==== \n")
	test4()
	log.Printf("==== \n")
	test5()
	log.Printf("==== \n")
	test6()
}

//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
func test1() {
	r, err := Open("a")
	if err != nil {
		log.Fatalf("error opening 'a'\n")
	}
	defer r.Close()

	r.Use()
}

//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
func test2() {
	r := subc()
	r.Use()
}

func subc() *Resource {
	r, err := Open("a")
	if err != nil {
		log.Fatalf("error opening 'a'\n")
	}
	defer r.Close()
	return r
}

//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
func test3() {
	r, err := Open("a")
	if err != nil {
		log.Fatalf("error opening 'a'\n")
	}
	defer r.Close()

	r, err = Open("b")
	if err != nil {
		log.Fatalf("error opening 'b'\n")
	}
	defer r.Close()
}

//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
/*
Now, after closing the resource, we will open another one (with a different
name). Can we reuse the same variable?
*/

func test4() {
	{
		r, err := Open("a")
		if err != nil {
			log.Fatalf("error opening 'a'\n")
		}
		defer r.Close()

		r.Use()
	}
	{
		r, err := Open("b")
		if err != nil {
			log.Fatalf("error opening 'b'\n")
		}
		defer r.Close()

		r.Use()
	}
}

func test5() {
	r, err := Open("a")
	if err != nil {
		log.Fatalf("error opening 'a'\n")
	}
	defer func() {
		err := r.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	r2, err := Open("b")
	if err != nil {
		log.Fatalf("error opening 'b'\n")
	}
	defer func(r *Resource) {
		err := r2.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(r2)
}

//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
// Create a named cleanup function and use the existing io.Closer interface:

func test6() {
	r, err := Open("a")
	if err != nil {
		log.Fatalf("error opening 'a'\n")
	}
	defer Close(r)

	r, err = Open("b")
	if err != nil {
		log.Fatalf("error opening 'b'\n")
	}
	defer Close(r)
}

func Close(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Fatal(err)
	}
}

//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
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

2016/12/03 17:21:39 opening a
2016/12/03 17:21:39 using a
2016/12/03 17:21:39 closing a
2016/12/03 17:21:39 ====
2016/12/03 17:21:39 opening a
2016/12/03 17:21:39 closing a
2016/12/03 17:21:39 using a
2016/12/03 22:13:46 ====
2016/12/03 22:13:46 opening a
2016/12/03 22:13:46 opening b
2016/12/03 22:13:46 closing b
2016/12/03 22:13:46 closing a
2016/12/03 22:13:46 ====
2016/12/03 22:13:46 opening a
2016/12/03 22:13:46 using a
2016/12/03 22:13:46 opening b
2016/12/03 22:13:46 using b
2016/12/03 22:13:46 closing b
2016/12/03 22:13:46 closing a
2016/12/03 22:25:37 ====
2016/12/03 22:25:37 opening a
2016/12/03 22:25:37 opening b
2016/12/03 22:25:37 closing b
2016/12/03 22:25:37 closing a
2016/12/03 22:25:37 ====
2016/12/03 22:25:37 opening a
2016/12/03 22:25:37 opening b
2016/12/03 22:25:37 closing b
2016/12/03 22:25:37 closing a

*/
