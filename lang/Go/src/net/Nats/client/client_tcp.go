package main

import (
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	// Define the NATS server URL
	natsURL := "nats://localhost:54222"

	// Connect to NATS server
	nc, err := nats.Connect(natsURL)
	if err != nil {
		log.Fatalf("Error connecting to NATS server: %v", err)
	}
	defer nc.Close()

	log.Printf("Connected to NATS server at %s", natsURL)

	// Subscribe to a subject
	subject := "updates"
	sub, err := nc.SubscribeSync(subject)
	if err != nil {
		log.Fatalf("Error subscribing to subject %s: %v", subject, err)
	}
	log.Printf("Subscribed to subject %s", subject)

	// Publish a message to the subject
	message := "Hello, NATS!"
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
