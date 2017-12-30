////////////////////////////////////////////////////////////////////////////
// Porgram: FileExist.go
// Purpose: Go check file exists demo
// Authors: Tong Sun (c) 2017, All rights reserved
// Credits: As listed below
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"os"

	"github.com/labstack/gommon/color"
)

func main() {
	test0()
	fmt.Println("======")
	test1()
	fmt.Println("======")
	test2()
	fmt.Println("======")
	test3()
}

func test0() {
	color.Println(color.Black("black"))
	color.Println(color.Red("red"))
	color.Println(color.Green("green"))
	color.Println(color.Yellow("yellow"))
	color.Println(color.Blue("blue"))
	color.Println(color.Magenta("magenta"))
	color.Println(color.Cyan("cyan"))
	color.Println(color.White("white"))
	color.Println(color.Grey("grey"))
}

func test1() {
	fmt.Println(color.Black("black"))
	fmt.Println(color.Red("red"))
	fmt.Println(color.Green("green"))
	fmt.Println(color.Yellow("yellow"))
	fmt.Println(color.Blue("blue"))
	fmt.Println(color.Magenta("magenta"))
	fmt.Println(color.Cyan("cyan"))
	fmt.Println(color.White("white"))
	fmt.Println(color.Grey("grey"))
}

func test2() {
	fmt.Println(color.Black("black"),
		color.Red("red"),
		color.Green("green"),
		color.Yellow("yellow"),
		color.Blue("blue"),
		color.Magenta("magenta"),
		color.Cyan("cyan"),
		color.White("white"),
		color.Grey("grey"))
}

func test3() {
	fmt.Fprintln(os.Stderr, color.Black("black"),
		color.Red("red"),
		color.Green("green"),
		color.Yellow("yellow"),
		color.Blue("blue"),
		color.Magenta("magenta"),
		color.Cyan("cyan"),
		color.White("white"),
		color.Grey("grey"))
}
