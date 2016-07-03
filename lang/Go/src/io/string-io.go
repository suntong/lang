package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func main() {
	// String Conversion to io.Reader
	// Use strings.Reader, which implements io.Reader
	r := strings.NewReader("This is just a long\t long string.\n")
	// Use bytes.Buffer, which implements io.Writer, as the sink

	// Pass it to the function that expect io.Reader
	fmt.Print(StreamToString(r))
}

func StreamToByte(stream io.Reader) []byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.Bytes()
}

func StreamToString(stream io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.String()
}
