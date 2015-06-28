////////////////////////////////////////////////////////////////////////////
// Porgram: Template-Html
// Purpose: Go html template usage demo
// Authors: Tong Sun (c) 2015, All rights reserved
// Based on https://www.socketloop.com/references/golang-html-template-parsefiles-function-examples
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"html/template"
	"os"
)

func main() {

	data := map[string]interface{}{
		"Title": "Hello World!",
	}

	t := template.New("HELLO")

	err := t.ExecuteTemplate(os.Stdout, "Template-Html.tmpl", data)
	// not working, empty output.
	// So the template file MUST be specified twice!
	var templates = template.Must(t.ParseFiles("Template-Html.tmpl"))
	err = templates.ExecuteTemplate(os.Stdout, "Template-Html.tmpl", data)
	// working
	fmt.Println("\n\n")

	// This option, need to use the template file twice, once for
	// `ParseFiles`, then for `ExecuteTemplate`. That looks awkward and
	// unnecessary. The following one looks straightforward to me.

	t = template.New("Test template")
	t, err = t.Parse("<title>{{ .Title }}</title>")
	// working
	t, err = template.ParseFiles("Template-Html.tmpl")
	// working now! Note the difference!!
	checkError(err)

	err = t.Execute(os.Stdout, data)

	fmt.Println("\n\nTemplate name is : ", t.Name())

}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
