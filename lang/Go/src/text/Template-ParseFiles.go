package main

import (
	"fmt"
	"os"
	"text/template"
)

type Person string

func main() {
	t := template.New("Template-ParseFiles.a.tmpl")
	fmt.Println("Name: ", t.Name())
	fmt.Printf("Deck: %#v\n", t.Templates())
	if _, err := t.ParseFiles("Template-ParseFiles.a.tmpl", "Template-ParseFiles.b.tmpl"); err != nil {
		panic(err)
	}
	fmt.Println("Name: ", t.Name())
	fmt.Printf("Deck: %#v\n", t.Templates())
	for _, ti := range t.Templates() {
		fmt.Println("  Name: ", ti.Name())
	}
	if err := t.Execute(os.Stdout, nil); err != nil {
		panic(err)
	}

	if err := t.ExecuteTemplate(os.Stdout, "Template-ParseFiles.b.tmpl", nil); err != nil {
		panic(err)
	}

}
