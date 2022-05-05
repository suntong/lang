// https://go.dev/play/p/2mcEhmjUKtk

package main

import (
	"github.com/davecgh/go-spew/spew"
)

type Demo struct {
	Id   int
	Name string
	Desc string
	Pids []string
}

func main() {
	demo := Demo{1, "Gordon", "This is my name", []string{"pic1", "pic2"}}
	spew.Dump(demo)

}

/*

(main.Demo) {
 Id: (int) 1,
 Name: (string) (len=6) "Gordon",
 Desc: (string) (len=15) "This is my name",
 Pids: ([]string) (len=2 cap=2) {
  (string) (len=4) "pic1",
  (string) (len=4) "pic2"
 }
}

*/
