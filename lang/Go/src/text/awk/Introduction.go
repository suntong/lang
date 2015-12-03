package main

import (
	"github.com/spakin/awk"
	"os"
)

func main() {
	s := awk.NewScript()
	s.AppendStmt(func(s *awk.Script) bool { return s.F(1).Int()%2 == 1 }, nil)
	if err := s.Run(os.Stdin); err != nil {
		panic(err)
	}
}

/*

Introduction
https://godoc.org/github.com/spakin/awk#hdr-Introduction

For first column is an odd number:

$5 % 2 == 1

$ seq 10 | go run Introduction.go
1
3
5
7
9

*/
