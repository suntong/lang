////////////////////////////////////////////////////////////////////////////
// Porgram: demo-redirect.go
// Purpose: Demo http redirect after post
// Authors: Tong Sun (c) 2023, All rights reserved
// Credits: https://stackoverflow.com/questions/35934298/how-to-redirect-to-a-url
////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////
// Program start

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

const chatUrl = "/chat"

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

var userName string

////////////////////////////////////////////////////////////////////////////
// Main

func main() {
	http.HandleFunc("/", login)
	http.HandleFunc(chatUrl, chat)
	err := http.ListenAndServe(":9090", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

////////////////////////////////////////////////////////////////////////////
// Function definitions

//==========================================================================
// login
// https://www.golangprograms.com/example-to-handle-get-and-post-request-in-golang.html
func login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// Check GET params first
		userName = r.URL.Query().Get("u")
		if userName == "" {
			// https://pkg.go.dev/html/template#Template.Parse
			t := template.Must(template.New("l").Parse(login_html))
			t.Execute(w, nil)
			return
		}
		// Else: userName passed as GET param, go to redirect
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		userName = r.FormValue("u") // saveChoice
		fmt.Println("username:", userName, r.Form["u"])
	default:
		fmt.Fprintf(w, "Only GET and POST methods supported.")
	}
	http.Redirect(w, r, chatUrl, http.StatusSeeOther)
}

//==========================================================================
// chat
func chat(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello "+userName+"!") // write data to response
}

////////////////////////////////////////////////////////////////////////////
// Long Constant definitions

var login_html = `
<!DOCTYPE html>
<html>
 <head>
 <title>Login</title>
 </head>
 <body>
  <form action="/" method="post">
   Username:<input type="text" name="u">
   <input type="submit" value="Login">
  </form>
 </body>
</html>
`
