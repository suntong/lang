package main

import (
	"context"
	"log"
	"net"
	"net/url"
	"time"

	"github.com/nats-io/nats.go"
	"nhooyr.io/websocket"
)

// wsDialer is a simple struct to hold the WebSocket URL
type wsDialer struct {
	url string
}

// Dial establishes a WebSocket connection and wraps it in a net.Conn
func (d *wsDialer) Dial(network, address string) (net.Conn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	wsConn, _, err := websocket.Dial(ctx, d.url, nil)
	if err != nil {
		return nil, err
	}
	return &webSocketConn{Conn: wsConn}, nil
}

// webSocketConn wraps the nhooyr.io/websocket connection to implement net.Conn
type webSocketConn struct {
	*websocket.Conn
	readCtx  context.Context
	writeCtx context.Context
	cancel   context.CancelFunc
}

func newWebSocketConn(ws *websocket.Conn) *webSocketConn {
	return &webSocketConn{Conn: ws, readCtx: context.Background(), writeCtx: context.Background()}
}

// Read reads from the WebSocket connection
func (c *webSocketConn) Read(b []byte) (n int, err error) {
	_, r, err := c.Conn.Reader(context.Background())
	if err != nil {
		return 0, err
	}
	return r.Read(b)
}

// Write writes to the WebSocket connection
func (c *webSocketConn) Write(b []byte) (n int, err error) {
	w, err := c.Conn.Writer(context.Background(), websocket.MessageBinary)
	if err != nil {
		return 0, err
	}
	defer w.Close()
	return w.Write(b)
}

// Close closes the WebSocket connection
func (c *webSocketConn) Close() error {
	return c.Conn.Close(websocket.StatusNormalClosure, "")
}

// LocalAddr returns the local network address (not implemented)
func (c *webSocketConn) LocalAddr() net.Addr {
	return nil
}

// RemoteAddr returns the remote network address (not implemented)
func (c *webSocketConn) RemoteAddr() net.Addr {
	return nil
}

// SetDeadline sets the read and write deadlines (not implemented)
func (c *webSocketConn) SetDeadline(t time.Time) error {
	c.readCtx, c.cancel = context.WithDeadline(context.Background(), t)
	c.writeCtx, c.cancel = context.WithDeadline(context.Background(), t)
	return nil
}

// SetReadDeadline sets the read deadline (not implemented)
func (c *webSocketConn) SetReadDeadline(t time.Time) error {
	c.readCtx, c.cancel = context.WithDeadline(context.Background(), t)
	return nil
}

// SetWriteDeadline sets the write deadline (not implemented)
func (c *webSocketConn) SetWriteDeadline(t time.Time) error {
	c.writeCtx, c.cancel = context.WithDeadline(context.Background(), t)
	return nil
}

func main() {
	// WebSocket server URL
	u := url.URL{Scheme: "ws", Host: "localhost:58220"}
	log.Printf("Connecting to %s", u.String())

	// Create a custom WebSocket dialer
	d := &wsDialer{
		url: u.String(),
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
