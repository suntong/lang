////////////////////////////////////////////////////////////////////////////
// Porgram: SocketFileServer
// Purpose: Use io.Copy to service a file to socket
// Authors: Tong Sun (c) 2013; Jan Newmarch (c) 2012
////////////////////////////////////////////////////////////////////////////

package main

import (
    "fmt"
    "net"
    "os"
    "io"
)

func main() {
    service := ":1200"
    tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
    checkError(err)
    listener, err := net.ListenTCP("tcp4", tcpAddr)
    checkError(err)
    for {
        conn, err := listener.Accept()
        if err != nil {
            continue
        }
        go handleClient(conn)
    }
}

func handleClient(conn net.Conn) {
    defer conn.Close()

    srcName := "/tmp/SocketFileServer"
    src, err := os.Open(srcName)
    if err != nil {
        return
    }
    defer src.Close()

    dst := conn

    io.Copy(dst, src)
    // we're finished with this client
}
func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}
