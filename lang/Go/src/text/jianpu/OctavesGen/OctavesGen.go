////////////////////////////////////////////////////////////////////////////
// Porgram: Octaves.go
// Purpose: Octaves jianpu generation
// Authors: Tong Sun (c) 2023, All rights reserved
////////////////////////////////////////////////////////////////////////////

//go:generate stringer -type=solfa

////////////////////////////////////////////////////////////////////////////
// Program start

package main

import (
	"fmt"
	"os"
	"text/template"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

type solfa int

const (
	_ solfa = iota
	Do
	Re
	Mi
	Fa
	So
	La
	Ti
)

type scale int

const (
	C3 scale = iota
	D3
	E3
	F3
	G3
	A3
	B3
	C4
	D4
	E4
	F4
	G4
	A4
	B4
	C5
	D5
	E5
	F5
	G5
	A5
	B5
)

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

// jot 7 | ssuf '\,' | sed 's/.*/"&", /' | tr -d '\n'
var noteJP = []string{
	"1,", "2,", "3,", "4,", "5,", "6,", "7,",
	"1", "2", "3", "4", "5", "6", "7",
	"1'", "2'", "3'", "4'", "5'", "6'", "7'",
}

var (
	tmpl1 = template.Must(template.New("1").Parse(template1))
	tmpl2 = template.Must(template.New("2").Parse(template2))
	tmpl3 = template.Must(template.New("3").Parse(template3))
)

////////////////////////////////////////////////////////////////////////////
// Main

func main() {

	vars := make(map[string]interface{})

	for n := Do; n <= Ti; n++ {
		fn := fmt.Sprintf("Oct8-%d%s", n, n)
		//print(fn)
		file, _ := os.Create("/tmp/" + fn + ".jp")
		defer file.Close()
		vars["Name"] = fn

		// file header
		tmpl1.Execute(file, vars)

		cur := scale(n) + C4 - 1
		nn := cur - 7
		vars["Cur"] = noteJP[cur]
		// == upwards
		// from one less octave up to cur
		for ; nn < cur; nn++ {
			//fmt.Println(noteJP[nn])
			vars["Inc"] = noteJP[nn]
			tmpl2.Execute(file, vars)
		}
		fmt.Fprintf(file, "%s  - 0 \\break \n", vars["Cur"])
		nn++ // skip cur
		// up to next octave
		for ; nn < cur+7; nn++ {
			vars["Inc"] = noteJP[nn]
			tmpl2.Execute(file, vars)
		}
		vars["Inc"] = noteJP[nn]
		fmt.Fprintf(file, "%s %s ( - %[2]s - ) 0 \\break \n", vars["Cur"], vars["Inc"])

		// == downwards
		for ; nn > cur; nn-- {
			vars["Inc"] = noteJP[nn]
			tmpl3.Execute(file, vars)
		}
		fmt.Fprintf(file, "%s  - 0 \\break \n", vars["Cur"])
		nn-- // skip cur
		for ; nn >= cur-7; nn-- {
			vars["Inc"] = noteJP[nn]
			tmpl3.Execute(file, vars)
		}
		fmt.Fprintf(file, "%s  - 0\n", vars["Cur"])
	}

}

////////////////////////////////////////////////////////////////////////////
// Long Constant definitions

const (
	template1 = `% {{.Name}}
%% tempo: 4=120

1=C
3/4

`
	template2 = "{{.Cur}} {{.Inc}} -  "
	template3 = "{{.Inc}} - {{.Cur}}  "
)
