package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/nats-io/nats.go"
	"nhooyr.io/websocket"
)

type ConnectionWrapper struct {
	ws *websocket.Conn
}

var _ nats.CustomDialer = (*ConnectionWrapper)(nil) // Verify the implementation

func (cw ConnectionWrapper) Dial(network, address string) (net.Conn, error) {
	log.Printf("Got Request for Network: %q / Address: %q", network, address)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	log.Printf("Dialing Address: ws://%s", address)
	c, _, err := websocket.Dial(ctx, "ws://"+address, nil)
	if err != nil {
		log.Printf("websocket.Dial failed %#v", err)
		return nil, fmt.Errorf("websocket.Dial failed %w", err)
	}
	cw.ws = c
	nconn := websocket.NetConn(context.Background(), c, websocket.MessageBinary)
	return nconn, nil
}

func (cw ConnectionWrapper) SkipTLSHandshake() bool {
	return true
}

type chatType int

const (
	CSignIn chatType = iota
	CText
)

type chatData struct {
	CFrom, CMsg string
	CType       chatType
}

func main() {
	// NCT: Nat Client Talker
	whoami := os.Getenv("NCT_ME")
	subject := os.Getenv("NCT_SBJ")
	message := os.Getenv("NCT_MSG")
	server_url := os.Getenv("NCT_SVR")
	// the username and password for NATS authentication
	username := os.Getenv("NCT_USER")
	password := os.Getenv("NCT_PASS")

	var nc *nats.Conn
	var ncw ConnectionWrapper
	log.Println("NatsConnect")

	// Create a NATS connection using the custom WebSocket dialer with username and password
	opts := nats.Options{
		//AllowReconnect: true,
		Url:          server_url,
		User:         username,
		Password:     password,
		CustomDialer: ncw,
	}
	var err error
	nc, err = opts.Connect()
	if err != nil {
		log.Printf("Native Go Connect did fail: %#v", err)
		return
	}
	log.Println("Nats connected through websocket and netConn wrapper!")
	log.Println("Connected to: " + nc.ConnectedServerName())

	// Subscribe to a subject
	sub, err := nc.SubscribeSync(subject)
	if err != nil {
		log.Fatalf("Error subscribing to subject %s: %v", subject, err)
	}
	log.Printf("Subscribed to subject %s", subject)

	// Publish the message to the subject
	c := chatData{CFrom: whoami, CType: CText, CMsg: message}
	b, _ := json.Marshal(c)
	err = nc.Publish(subject, b)
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
