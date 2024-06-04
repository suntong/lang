package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

var (
	addr = os.Getenv("WS_ADDR") // "localhost:8080"
)

func main() {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := websocket.Accept(w, r, &websocket.AcceptOptions{
			// Customize the AcceptOptions if needed
		})
		if err != nil {
			log.Printf("Failed to accept WebSocket connection: %v", err)
			return
		}
		defer conn.Close(websocket.StatusInternalError, "the server is shutting down")

		// Handle the WebSocket connection
		handleWebSocketConnection(r.Context(), conn)
	})

	// Create a server with graceful shutdown
	server := &http.Server{
		Addr: addr,
	}

	// Channel to listen for interrupt signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Run the server in a separate goroutine
	go func() {
		log.Printf("Server started at %s", addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on %s: %v\n", addr, err)
		}
	}()

	// Wait for interrupt signal
	<-stop

	// Shutdown the server with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	log.Println("Server gracefully stopped")
}

func handleWebSocketConnection(ctx context.Context, conn *websocket.Conn) {
	for {
		select {
		case <-ctx.Done():
			log.Println("Context cancelled, closing WebSocket connection")
			conn.Close(websocket.StatusNormalClosure, "context cancelled")
			return
		default:
			var message interface{}
			err := wsjson.Read(ctx, conn, &message)
			if err != nil {
				if websocket.CloseStatus(err) == websocket.StatusNormalClosure {
					log.Println("WebSocket connection closed by client")
					return
				}
				log.Printf("Failed to read message: %v", err)
				return
			}

			log.Printf("Received: %v", message)

			// Echo the message back to the client
			err = wsjson.Write(ctx, conn, message)
			if err != nil {
				log.Printf("Failed to write message: %v", err)
				return
			}

			log.Printf("Echoed: %v", message)
		}
	}
}
