////////////////////////////////////////////////////////////////////////////
// Porgram: FileCgiServer
// Purpose: Go FileServer with CGI support
// Authors: Tong Sun (c) 2013; Jan Newmarch (c) 2012
////////////////////////////////////////////////////////////////////////////

package main

import (
    "fmt"
    "net/http"
    "os"
)

func main() {
    // file handler for most files
    // deliver files from the directory /var/www 
    fileServer := http.FileServer(http.Dir("/var/www"))
    http.Handle("/", fileServer)

    // function handler for /cgi-bin/printenv
    // http://jan.newmarch.name/go/http/chapter-http.html
    http.HandleFunc("/cgi-bin/printenv", printEnv)

    // register the handler and deliver requests to it via the default multiplexer
    err := http.ListenAndServe(":8000", nil)
    // will never get here unless something went wrong
    checkError(err)
    // That's it!
}

func printEnv(writer http.ResponseWriter, req *http.Request) {
    env := os.Environ()
    writer.Write([]byte("<h1>Environment</h1>\n<pre>"))
    for _, v := range env {
        writer.Write([]byte(v + "\n"))
    }
    writer.Write([]byte("</pre>"))
}

func checkError(err error) {
    if err != nil {
        fmt.Println("Fatal error ", err.Error())
        os.Exit(1)
    }
}
