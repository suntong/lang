package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

const (
	webSocketTimeout = 3
)

var (
	webSocketURL = os.Getenv("WS_URL") // WebSocket server URL
	webSocketMsg = os.Getenv("WS_MSG") // WebSocket message
)

func main() {
	// Create a context that cancels on interrupt signals
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle interrupt signals for graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Println("Received interrupt signal, shutting down...")
		cancel()
	}()

	// Create a timeout context for the connection
	connCtx, connCancel := context.WithTimeout(ctx, webSocketTimeout*time.Second)
	defer connCancel()

	// Connect to WebSocket server with timeout
	conn, _, err := websocket.Dial(connCtx, webSocketURL, nil)
	if err != nil {
		log.Fatalf("Failed to connect to WebSocket server: %v", err)
	}
	defer conn.Close(websocket.StatusInternalError, "the client is shutting down")

	log.Println("Connected to WebSocket server:", webSocketURL)
	// WaitGroup to wait for goroutines to finish
	var wg sync.WaitGroup
	wg.Add(2)

	// Channel to signal when to send the next message
	sendChan := make(chan struct{})

	// Start goroutine for sending messages
	go func() {
		defer wg.Done()
		sendMessages(ctx, conn, sendChan)
	}()

	// Start goroutine for receiving messages
	go func() {
		defer wg.Done()
		receiveMessages(ctx, conn, sendChan)
	}()

	// Wait for both goroutines to finish
	wg.Wait()

	// Close the connection
	conn.Close(websocket.StatusNormalClosure, "normal closure")
	log.Println("Connection closed gracefully")
}

func sendMessages(ctx context.Context, conn *websocket.Conn, sendChan chan struct{}) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	startTime := time.Now()
	for {
		select {
		case <-ctx.Done():
			log.Println("Context cancelled, stopping sendMessages goroutine")
			return
		case <-ticker.C:
			message := webSocketMsg
			sendCtx, sendCancel := context.WithTimeout(ctx, webSocketTimeout*time.Second)
			defer sendCancel()
			startTime = time.Now() // Record start time
			err := wsjson.Write(sendCtx, conn, message)
			if err != nil {
				log.Printf("Failed to send message: %v", err)
				return
			}
			log.Printf("Sent: %s", message)
			// Notify the receive goroutine to start receiving, after startTime
			sendChan <- struct{}{}
			log.Printf("Time from send to receive: %v", time.Since(startTime))
		}
	}
}

func receiveMessages(ctx context.Context, conn *websocket.Conn, sendChan chan struct{}) {
	for {
		select {
		case <-ctx.Done():
			log.Println("Context cancelled, stopping receiveMessages goroutine")
			return
		case <-sendChan:
			recvCtx, recvCancel := context.WithTimeout(ctx, webSocketTimeout*time.Second)
			defer recvCancel()
			var message string
			err := wsjson.Read(recvCtx, conn, &message)
			if err != nil {
				if websocket.CloseStatus(err) == websocket.StatusNormalClosure {
					log.Println("WebSocket connection closed by server")
					return
				}
				log.Printf("Failed to receive message: %v", err)
				return
			}
			log.Printf("Received: %s", message)
		}
	}
}

/*

go get -v nhooyr.io/websocket
go build -o demor demor.go

export WS_URL="ws://ws.vi-server.org/mirror/" WS_MSG="Request payload"

$ demor
2024/06/01 10:51:59 Connected to WebSocket server: ws://ws.vi-server.org/mirror/
2024/06/01 10:52:04 Sent: Request payload
2024/06/01 10:52:04 Time from send to receive: 421.375µs
2024/06/01 10:52:04 Received: Request payload
2024/06/01 10:52:09 Sent: Request payload
2024/06/01 10:52:09 Time from send to receive: 372.584µs
2024/06/01 10:52:09 Received: Request payload
^C2024/06/01 10:52:10 Received interrupt signal, shutting down...
2024/06/01 10:52:10 Context cancelled, stopping sendMessages goroutine
2024/06/01 10:52:10 Context cancelled, stopping receiveMessages goroutine
2024/06/01 10:52:10 Connection closed gracefully

*/
