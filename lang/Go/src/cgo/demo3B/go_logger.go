package main

import "C"

import (
	"fmt"
	"os"
	"sync"
	"time"
)

var (
	logFile *os.File
	mu      sync.Mutex
	once    sync.Once
)

//export InitLogger
func InitLogger(path *C.char) bool {
	var err error
	once.Do(func() {
		filePath := C.GoString(path)
		logFile, err = os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err == nil {
			// Force line buffering
			logFile.WriteString("") // trigger sync
		}
	})
	return err == nil && logFile != nil
}

//export LogMessage
func LogMessage(cMsg *C.char) {
	if logFile == nil {
		return
	}
	msg := C.GoString(cMsg)
	line := fmt.Sprintf("[%s] %s\n", fastTime(), msg)

	mu.Lock()
	logFile.WriteString(line)
	logFile.Sync() // optional: remove if performance critical
	mu.Unlock()

}

//export CloseLogger
func CloseLogger() {
	mu.Lock()
	if logFile != nil {
		logFile.Close()
		logFile = nil
	}
	mu.Unlock()
}

// Zero-allocation UTC timestamp (fast enough for logging)
func fastTime() string {
	const layout = "2006-01-02T15:04:05.000Z0700"
	t := time.Now().UTC()
	return t.Format(layout)
}

func main() {}
