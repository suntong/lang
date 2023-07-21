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
	"bytes"
	"encoding/gob"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strings"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

const (
	rootUrl = "/"
	chatUrl = "/chat"
)

type session struct {
	Name string
	Text string
}

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

////////////////////////////////////////////////////////////////////////////
// Main

func main() {
	// Importantly, we need to tell the encoding/gob package about the Go type
	// that we want to encode. We do this by passing *an instance* of the type
	// to gob.Register(). In this case we pass a pointer to an initialized (but
	// empty) instance of the session struct.
	gob.Register(&session{})

	// routing
	http.HandleFunc(rootUrl, login)
	http.HandleFunc(chatUrl, chat)
	http.HandleFunc("/logout", logout)
	// server start
	log.Print("Listening...")
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
			t := template.Must(template.New("").Parse(login_html))
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
		//fmt.Printf("%+v\n", r)
		userName = r.FormValue("u") // saveChoice
		if userName == "" {
			// for empty user name, use remote port as anonymous name
			userName = r.RemoteAddr
			re := regexp.MustCompile(`.*:.(.*)$`)
			userName = re.ReplaceAllString(userName, "${1}")
		}
		log.Println("user", userName, "signed in")
	default:
		fmt.Fprintf(w, "Only GET and POST methods supported.")
	}
	var buf bytes.Buffer
	err := gob.NewEncoder(&buf).Encode(session{Name: userName, Text: ""})
	if err != nil {
		log.Println(err)
		http.Error(w, "server error: gob encoding", http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:  "session",
		Value: buf.String(),
		Path:  rootUrl,
	})
	http.Redirect(w, r, chatUrl, http.StatusSeeOther)
}

//==========================================================================
// chat
func chat(w http.ResponseWriter, r *http.Request) {
	userName := getUserName(w, r)
	t := template.Must(template.New("").Parse(chat_html))
	t.Execute(w, session{
		Name: userName,
		Text: "",
	})
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
	c, err := r.Cookie("session")
	if err != nil {
		http.Redirect(w, r, rootUrl, http.StatusSeeOther)
		return ""
	}
	log.Println("session raw", c.Value)

	var s session
	fmt.Printf("%+v\n", s)
	reader := strings.NewReader(c.Value)
	if err := gob.NewDecoder(reader).Decode(&s); err != nil {
		log.Println(err)
		http.Error(w, "server error: gob decoding", http.StatusInternalServerError)
		return ""
	}

	return s.Name // userName
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
var chat_html = `
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
  Hello {{.Name}}!
 </body>
</html>
`
