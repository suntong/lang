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

const (
	rootUrl = "/"
	chatUrl = "/chat"
)

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

////////////////////////////////////////////////////////////////////////////
// Main

func main() {
	http.HandleFunc(rootUrl, login)
	http.HandleFunc(chatUrl, chat)
	http.HandleFunc("/logout", logout)
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
	userName := ""
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
		log.Println("user", userName, "signed in")
	default:
		fmt.Fprintf(w, "Only GET and POST methods supported.")
	}
	http.SetCookie(w, &http.Cookie{
		Name:  "username",
		Value: userName,
		Path:  rootUrl,
	})
	http.Redirect(w, r, chatUrl, http.StatusSeeOther)
}

//==========================================================================
// chat
func chat(w http.ResponseWriter, r *http.Request) {
	userName := getUserName(w, r)
	fmt.Fprintf(w, chat_html_beg+"Hello "+userName+"!"+chat_html_end)
}

//==========================================================================
// logout
func logout(w http.ResponseWriter, r *http.Request) {
	userName := getUserName(w, r)
	log.Println("user", userName, "signed out")
	// https://stackoverflow.com/a/59736764/2125837
	http.SetCookie(w, &http.Cookie{
		Name:   "username",
		Value:  "",
		Path:   rootUrl,
		MaxAge: -1,
	})
	http.Redirect(w, r, rootUrl, http.StatusSeeOther)
}

//==========================================================================
// getUserName
func getUserName(w http.ResponseWriter, r *http.Request) string {
	c, err := r.Cookie("username")
	if err != nil {
		http.Redirect(w, r, rootUrl, http.StatusSeeOther)
		return ""
	}

	return c.Value // userName
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

// https://stackoverflow.com/a/2906586/2125837
var chat_html_beg = `
<!DOCTYPE html>
<html>
 <head>
 <title>Chat</title>
 </head>
 <body>
  <form action="/logout" method="post">
   <input type="submit" value="Logout">
  </form>
  <hr>
`

var chat_html_end = `
 </body>
</html>
`
