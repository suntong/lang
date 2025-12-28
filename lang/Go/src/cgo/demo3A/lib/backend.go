package main

import "C"

import (
	"log"
	"os"
	"sync"
)

var (
	mu   sync.Mutex
	file *os.File
)

//export InternalGoLogger
func InternalGoLogger(cMsg *C.char) {
	// 1. Convert C String to Go String immediately
	msg := C.GoString(cMsg)

	mu.Lock()
	defer mu.Unlock()

	// 2. Lazy init file
	if file == nil {
		var err error
		file, err = os.OpenFile("system.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println("Go: Error opening log:", err)
			return
		}
	}

	// 3. Write to file
	if _, err := file.WriteString(msg + "\n"); err != nil {
		log.Println("Go: Write error:", err)
	}
}

// Main is required for buildmode=c-shared
func main() {}
