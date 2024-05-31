package main

import (
	"context"
	"fmt"
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
	webSocketURL = "ws://ws.vi-server.org/mirror/" // Replace with your WebSocket server URL
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

	// Connect to WebSocket server
	conn, _, err := websocket.Dial(ctx, webSocketURL, nil)
	if err != nil {
		log.Fatalf("Failed to connect to WebSocket server: %v", err)
	}
	defer conn.Close(websocket.StatusInternalError, "the client is shutting down")

	// WaitGroup to wait for goroutines to finish
	var wg sync.WaitGroup
	wg.Add(2)

	// Start goroutine for sending messages
	go func() {
		defer wg.Done()
		sendMessages(ctx, conn)
	}()

	// Start goroutine for receiving messages
	go func() {
		defer wg.Done()
		receiveMessages(ctx, conn)
	}()

	// Wait for both goroutines to finish
	wg.Wait()

	// Close the connection
	conn.Close(websocket.StatusNormalClosure, "normal closure")
	log.Println("Connection closed gracefully")
}

func sendMessages(ctx context.Context, conn *websocket.Conn) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Println("Context cancelled, stopping sendMessages goroutine")
			return
		case t := <-ticker.C:
			message := fmt.Sprintf("Hello, the time is %s", t)
			err := wsjson.Write(ctx, conn, message)
			if err != nil {
				log.Printf("Failed to send message: %v", err)
				return
			}
			log.Printf("Sent: %s", message)
		}
	}
}

func receiveMessages(ctx context.Context, conn *websocket.Conn) {
	for {
		select {
		case <-ctx.Done():
			log.Println("Context cancelled, stopping receiveMessages goroutine")
			return
		default:
			var message string
			err := wsjson.Read(ctx, conn, &message)
			if err != nil {
				log.Printf("Failed to receive message: %v", err)
				return
			}
			log.Printf("Received: %s", message)
		}
	}
}

/*

go get -v nhooyr.io/websocket
go build -o demo2 demo2.go

$ demo2
2024/05/31 19:36:22 Sent: Hello, the time is 2024-05-31 19:36:22.490033 -0400 EDT m=+5.276632459
2024/05/31 19:36:22 Received: Hello, the time is 2024-05-31 19:36:22.490033 -0400 EDT m=+5.276632459
2024/05/31 19:36:27 Sent: Hello, the time is 2024-05-31 19:36:27.492419 -0400 EDT m=+10.278983542
2024/05/31 19:36:27 Received: Hello, the time is 2024-05-31 19:36:27.492419 -0400 EDT m=+10.278983542
2024/05/31 19:36:32 Sent: Hello, the time is 2024-05-31 19:36:32.489689 -0400 EDT m=+15.276219376
2024/05/31 19:36:32 Received: Hello, the time is 2024-05-31 19:36:32.489689 -0400 EDT m=+15.276219376
^C
2024/05/31 19:36:33 Received interrupt signal, shutting down...
2024/05/31 19:36:33 Context cancelled, stopping sendMessages goroutine
2024/05/31 19:36:33 Failed to receive message: failed to read JSON message: failed to get reader: use of closed network connection
2024/05/31 19:36:33 Connection closed gracefully

*/
