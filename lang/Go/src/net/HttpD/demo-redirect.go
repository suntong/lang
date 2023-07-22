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
	"encoding/base64"
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

type sessionT struct {
	Name string
	Text string
}

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

////////////////////////////////////////////////////////////////////////////
// Main

func main() {
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
	session := sessionT{Name: userName, Text: ""}
	err := gob.NewEncoder(&buf).Encode(session)
	if err != nil {
		log.Println(err)
		http.Error(w, "server error: gob encoding", http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:  "session",
		Value: base64.StdEncoding.EncodeToString(buf.Bytes()),
		Path:  rootUrl,
	})
	http.Redirect(w, r, chatUrl, http.StatusSeeOther)
}

//==========================================================================
// chat
func chat(w http.ResponseWriter, r *http.Request) {
	session := getSession(w, r)
	t := template.Must(template.New("").Parse(chat_html))
	t.Execute(w, session)
}

//==========================================================================
// logout
func logout(w http.ResponseWriter, r *http.Request) {
	session := getSession(w, r)
	log.Println("user", session.Name, "signed out")
	// https://stackoverflow.com/a/59736764/2125837
	http.SetCookie(w, &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   rootUrl,
		MaxAge: -1,
	})
	http.Redirect(w, r, rootUrl, http.StatusSeeOther)
}

//==========================================================================
// getSession
func getSession(w http.ResponseWriter, r *http.Request) sessionT {
	emptySession := sessionT{}

	c, err := r.Cookie("session")
	if err != nil {
		http.Redirect(w, r, rootUrl, http.StatusSeeOther)
		return emptySession
	}
	sr, err := base64.StdEncoding.DecodeString(c.Value) // session raw

	var s sessionT
	reader := strings.NewReader(string(sr))
	if err := gob.NewDecoder(reader).Decode(&s); err != nil {
		log.Println(err)
		http.Error(w, "server error: gob decoding", http.StatusInternalServerError)
		return emptySession
	}

	return s
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
