package main

import (
	"github.com/mkideal/cli"
)

type helloT struct {
	cli.Helper
	Name string `cli:"name" usage:"tell me your name" dft:"world"`
	Age  uint8  `cli:"a,age" usage:"tell me your age" dft:"100"`
}

func main() {
	cli.Run(new(helloT), func(ctx *cli.Context) error {
		argv := ctx.Argv().(*helloT)
		ctx.String("Hello, %s! Your age is %d?\n", argv.Name, argv.Age)
		return nil
	})
}

/*


To run it as a scripting language
https://dev.to/ignatk/using-go-as-a-scripting-language-in-linux-4c8c/

One time:

  go get -v -u github.com/erning/gorun
  sudo mv $GOPATH/bin/gorun /usr/local/bin/gorun

After every reboot:

  echo ':golang:E::go::/usr/local/bin/gorun:OC' | sudo tee /proc/sys/fs/binfmt_misc/register

then,

  $ 001-hello.go -a 75
  Hello, world! Your age is 75?

  $ 001-hello.go -a 75 --name Tom
  Hello, Tom! Your age is 75?


*/
