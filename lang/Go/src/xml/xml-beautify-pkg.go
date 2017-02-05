package main

import "github.com/go-xmlfmt/xmlfmt"

func main() {
    xml1 := `<root><this><is>a</is><test /><message><org><cn>Some org-or-other</cn><ph>Wouldnt you like to know</ph></org><contact><fn>Pat</fn><ln>Califia</ln></contact></message></this></root>`
    x := xmlfmt.FormatXML(xml1, "\t", "  ")
    print(x)
}