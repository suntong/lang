////////////////////////////////////////////////////////////////////////////
// Porgram: lfgrp
// Purpose: list files from the given path in groups with total sizes
// Authors: Tong Sun (c) 2015, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

type CurrentGrp struct {
	Prefix   string
	MatchPre int
	Files    []string
	Size     int64
}

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

////////////////////////////////////////////////////////////////////////////
// Main

func main() {

	// need one command line arguments
	if len(os.Args) <= 1 {
		fmt.Printf("Usage: lfgrp path\nList files from the given path in groups with total sizes\n\nExample: lfgrp . | sort -n\n")
		os.Exit(0)
	}

	// changes the current working directory to the given directory
	os.Chdir(os.Args[1])

	files, _ := ioutil.ReadDir("./")
	reportByGroup(files)
}

////////////////////////////////////////////////////////////////////////////
// Function definitions

func reportByGroup(files []os.FileInfo) {
	var cg CurrentGrp
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		s, _ := os.Stat(f.Name()) // because f.Size() do not properly
		// reflect symbolic link file size but s.Size() does

		if cg.Prefix == "" {
			cg.init(f.Name(), s.Size())
			continue
		}

		if m := matchTo(cg.Prefix, f.Name()); m >= len(f.Name())*3/4 ||
			m >= cg.MatchPre*3/4 {
			// match up to 3/4 of file name lengh, same group
			//fmt.Printf("] %v\n", cg)
			cg.Prefix = cg.Prefix[:m]
			cg.Files = append(cg.Files, f.Name())
			cg.Size += s.Size()
		} else {
			// different group, output previous group then reset
			cg.dump()
			cg.init(f.Name(), s.Size())
		}
	}
	// the last file, regardless if it is alone or within group
	cg.dump()
}

func (cg *CurrentGrp) init(fname string, size int64) {
	cg.Prefix = fname
	cg.MatchPre = len(fname)
	cg.Files = append([]string{}, fname)
	cg.Size = size
}

func (cg CurrentGrp) dump() {
	fmt.Printf(" %12d %6dM [%s]: %s\n",
		cg.Size, cg.Size/1024/1024, cg.Prefix, strings.Join(cg.Files, ", "))
}

func matchTo(s1, s2 string) int {
	i := 0
	for n := len(s1); i < n; i++ {
		if s1[i] != s2[i] {
			break
		}
	}
	return i
}

func debug(input string) {
	print("] ")
	print(input)
	print("\n")
}
