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

type wsDialer struct {
	url     string
	headers http.Header
}

func (d *wsDialer) Dial(network, address string) (net.Conn, error) {
	wsConn, _, err := websocket.DefaultDialer.Dial(d.url, d.headers)
	if err != nil {
		return nil, err
	}
	return newWebSocketConn(wsConn), nil
}

type webSocketConn struct {
	*websocket.Conn
}

func newWebSocketConn(ws *websocket.Conn) *webSocketConn {
	return &webSocketConn{Conn: ws}
}

func (c *webSocketConn) Read(b []byte) (n int, err error) {
	_, r, err := c.Conn.NextReader()
	if err != nil {
		return 0, err
	}
	return r.Read(b)
}

func (c *webSocketConn) Write(b []byte) (n int, err error) {
	w, err := c.Conn.NextWriter(websocket.BinaryMessage)
	if err != nil {
		return 0, err
	}
	defer w.Close()
	return w.Write(b)
}

func (c *webSocketConn) Close() error {
	return c.Conn.Close()
}

func (c *webSocketConn) LocalAddr() net.Addr {
	return c.Conn.LocalAddr()
}

func (c *webSocketConn) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *webSocketConn) SetDeadline(t time.Time) error {
	return c.Conn.UnderlyingConn().SetDeadline(t)
}

func (c *webSocketConn) SetReadDeadline(t time.Time) error {
	return c.Conn.UnderlyingConn().SetReadDeadline(t)
}

func (c *webSocketConn) SetWriteDeadline(t time.Time) error {
	return c.Conn.UnderlyingConn().SetWriteDeadline(t)
}

func main() {
	// WebSocket server URL
	u := url.URL{Scheme: "ws", Host: "localhost:58220"}
	log.Printf("Connecting to %s", u.String())

	// Create a custom WebSocket dialer
	d := &wsDialer{
		url: u.String(),
		headers: http.Header{
			"Origin": {"http://localhost"},
		},
	}

	// Create a NATS connection using the custom WebSocket dialer
	opts := nats.SetCustomDialer(d)
	nc, err := nats.Connect("", opts)
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
