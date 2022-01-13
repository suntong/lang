package main

import (
	"fmt"

	"github.com/jessevdk/go-flags"
)

var opts struct {
	// Slice of bool will append 'true' each time the option
	// is encountered (can be set multiple times, like -vvv)
	Verbose []bool `short:"v" long:"verbose" description:"Show verbose debug information"`

	// Example of automatic marshalling to desired type (uint)
	Offset uint `long:"offset" description:"Offset"`

	// Example of a callback, called each time the option is found.
	Call func(string) `short:"c" description:"Call phone number"`

	// Example of a required flag
	Name string `short:"n" long:"name" description:"A name" required:"true"`

	// Example of a value name
	File string `short:"f" long:"file" description:"A file" value-name:"FILE"`

	// Example of a pointer
	Ptr *int `short:"p" description:"A pointer to an integer"`

	// Example of a slice of strings
	StringSlice []string `short:"s" description:"A slice of strings"`

	// Example of a slice of pointers
	PtrSlice []*string `long:"ptrslice" description:"A slice of pointers to string"`

	// Example of a map
	IntMap map[string]int `long:"intmap" description:"A map from string to int"`

	// Example of positional arguments
	Args struct {
		ID   string
		Num  int
		Rest []string
	} `positional-args:"yes" required:"yes"`
}

func main() {

	// Callback which will invoke callto:<argument> to call a number.
	// Note that this works just on OS X (and probably only with
	// Skype) but it shows the idea.
	opts.Call = func(num string) {
		// cmd := exec.Command("open", "callto:"+num)
		// cmd.Start()
		// cmd.Process.Release()
		fmt.Printf("Calling: %s\n", num)
	}

	_, err := flags.Parse(&opts)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Verbosity: %v\n", opts.Verbose)
	fmt.Printf("Offset: %d\n", opts.Offset)
	fmt.Printf("Name: %s\n", opts.Name)
	fmt.Printf("Ptr: %d\n", *opts.Ptr)
	fmt.Printf("StringSlice: %v\n", opts.StringSlice)
	fmt.Printf("PtrSlice: [%v %v]\n", *opts.PtrSlice[0], *opts.PtrSlice[1])
	fmt.Printf("IntMap: [a:%v b:%v]\n", opts.IntMap["a"], opts.IntMap["b"])
	fmt.Printf("Args.ID: %s\n", opts.Args.ID)
	fmt.Printf("Args.Num: %d\n", opts.Args.Num)
	fmt.Printf("Args.Rest: %v\n", opts.Args.Rest)

}

/*

$ go run readme.go -vv --offset=5 -n Me -p 3 -s hello -s world --ptrslice hello --ptrslice world --intmap a:1 --intmap b:5 id 10 remaining1 remaining2
Verbosity: [true true]
Offset: 5
Name: Me
Ptr: 3
StringSlice: [hello world]
PtrSlice: [hello world]
IntMap: [a:1 b:5]
Args.ID: id
Args.Num: 10
Args.Rest: [remaining1 remaining2]

$ go run readme.go -vv --offset=5 -n Me -p 3 -s hello -s world --ptrslice hello --ptrslice world --intmap a:1 --intmap b:5 -c 123 -c 678 id 10 remaining1 remaining2
Calling: 123
Calling: 678
Verbosity: [true true]
Offset: 5
Name: Me
Ptr: 3
StringSlice: [hello world]
PtrSlice: [hello world]
IntMap: [a:1 b:5]
Args.ID: id
Args.Num: 10
Args.Rest: [remaining1 remaining2]

*/
