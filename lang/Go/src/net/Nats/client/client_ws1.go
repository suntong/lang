package main

import (
	"log"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
	"github.com/nats-io/nats.go"
)

func main() {
	// WebSocket server URL
	u := url.URL{Scheme: "wss", Host: "localhost:58220", Path: "/"}
	log.Printf("Connecting to %s", u.String())

	// Create a WebSocket dialer
	dialer := websocket.DefaultDialer

	// Dial the WebSocket server
	wsConn, _, err := dialer.Dial(u.String(), http.Header{})
	if err != nil {
		log.Fatal("Error connecting to WebSocket server:", err)
	}
	defer wsConn.Close()

	// Create a NATS encoded connection over the WebSocket
	nc, err := nats.Connect("", nats.CustomDialer(func(_, _ string) (net.Conn, error) {
		return wsConn.UnderlyingConn(), nil
	}))
	if err != nil {
		log.Fatal("Error creating NATS connection:", err)
	}
	defer nc.Close()

	log.Println("Connected to NATS server via WebSocket")

	// Subscribe to a subject
	subject := "updates"
	sub, err := nc.SubscribeSync(subject)
	if err != nil {
		log.Fatalf("Error subscribing to subject %s: %v", subject, err)
	}
	log.Printf("Subscribed to subject %s", subject)

	// Publish a message to the subject
	message := "Hello, NATS over WebSocket!"
	err = nc.Publish(subject, []byte(message))
	if err != nil {
		log.Fatalf("Error publishing message to subject %s: %v", subject, err)
	}
	log.Printf("Published message to subject %s: %s", subject, message)

	// Wait for a message
	msg, err := sub.NextMsg(10 * time.Second)
	if err != nil {
		log.Fatalf("Error receiving message: %v", err)
	}
	log.Printf("Received message from subject %s: %s", msg.Subject, string(msg.Data))
}
