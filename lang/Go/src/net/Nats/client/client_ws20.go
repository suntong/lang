package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
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
	var nc *nats.Conn
	var ncw ConnectionWrapper
	log.Println("NatsConnect")
	var err error
	nc, err = nats.Connect("localhost:58220",
		nats.Name("PWA-anonymous"),
		nats.SetCustomDialer(ncw),
	)
	if err != nil {
		log.Printf("Native Go Connect did fail: %#v", err)
		return
	}
	log.Println("Nats connected through websocket and netConn wrapper!")

	// Subscribe to a subject
	subject := "chat.say" // "chat.room"
	// _, err = nc.Subscribe(subject, func(msg *nats.Msg) {
	// 	// Print message data
	// 	var c chatData
	// 	json.Unmarshal(msg.Data, &c)
	// 	fmt.Printf("Got %+v\n", c)
	// })
	sub, err := nc.SubscribeSync(subject)
	if err != nil {
		log.Fatalf("Error subscribing to subject %s: %v", subject, err)
	}
	log.Printf("Subscribed to subject %s", subject)

	// Publish a message to the subject
	message := "Hello, NATS over WebSocket!"
	c := chatData{CFrom: "anonymousA", CType: CText, CMsg: message}
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
