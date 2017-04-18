////////////////////////////////////////////////////////////////////////////
// Porgram: EvalSymlinks
// Purpose: Go EvalSymlinks demo
// Authors: Tong Sun (c) 2017, All rights reserved
// Credits: http://stackoverflow.com/questions/18062026/resolve-symlinks-in-go
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func show_help_and_exit() {
	fmt.Println(`Usage ./EvalSymlinks path`)
	os.Exit(1)
}

func read_link(path string) {
	exit_code := 0
	defer os.Exit(exit_code)
	ln, err := filepath.EvalSymlinks(path)
	if err != nil {
		fmt.Println("[ERR]", err)
		exit_code = 1
		return
	}
	fmt.Printf("%s => %s\n", path, ln)
}

func main() {
	if len(os.Args) != 2 {
		show_help_and_exit()
	}
	read_link(os.Args[1])
}

/*

$ go run EvalSymlinks.go
Usage ./EvalSymlinks path

$ go run EvalSymlinks.go /bin/sh
/bin/sh => /bin/dash

$ go run EvalSymlinks.go /bin/pidof
/bin/pidof => /sbin/killall5

$ go run EvalSymlinks.go /bin/systemd
/bin/systemd => /lib/systemd/systemd

$ go run EvalSymlinks.go /bin/bash
/bin/bash => /bin/bash

*/
