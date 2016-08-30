////////////////////////////////////////////////////////////////////////////
// Porgram: Template-Range
// Purpose: Go template "range" demo
// Authors: Tong Sun (c) 2016, All rights reserved
// Credits: http://stackoverflow.com/questions/24556001
////////////////////////////////////////////////////////////////////////////

package main

import (
	"os"
	"text/template"
)

func main() {
	RangeOverStructOfSlices()
	RangeOverSliceOfStructs()
}

func RangeOverStructOfSlices() {
	print("== Range over struct of slices\n")

	// Define a template.
	const tmpl = `
IDs: {{range .Id}}{{.}} {{end}}
Names: {{range .Name}}
       {{.}}{{end}}
`

	// Prepare some data to insert into the template.
	type UserList struct {
		Id   []int
		Name []string
	}
	users := UserList{
		Id:   []int{0, 1, 2, 3, 4, 5, 6, 7},
		Name: []string{"user0", "user1", "user2", "user3", "user4"},
	}

	// Create a new template and parse the letter into it.
	t := template.Must(template.New("tmpl").Parse(tmpl))

	t.Execute(os.Stdout, users)
}

func RangeOverSliceOfStructs() {
	print("\n== Range over slice of structs\n")

	// Define a template.
	const tmpl = `
{{range .}}
	{{.Id}}
	{{.Name}}
{{end}}
`

	// Prepare some data to insert into the template.
	type User struct {
		Id   int
		Name string
	}

	type UserList []User
	var myuserlist UserList = UserList{
		{1, "a"},
		{2, "b"},
		{3, "c"},
	}

	// Create a new template and parse the letter into it.
	t := template.Must(template.New("tmpl").Parse(tmpl))

	t.Execute(os.Stdout, myuserlist)
}
