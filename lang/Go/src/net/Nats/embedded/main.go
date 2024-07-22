// See: https://www.youtube.com/watch?v=cdTrl8UfcBo

package main

import (
	"errors"
	"log"
	"time"

	"github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"
)

func main() {
	nc, ns, err := RunEmbeddedServer(true, true)
	if err != nil {
		log.Fatal(err)
	}

	nc.Subscribe("hello.world", func(msg *nats.Msg) {
		log.Println("message received!")
		msg.Respond([]byte("Ahoy there!"))
	})

	ns.WaitForShutdown()
}

func RunEmbeddedServer(inProcess bool, enableLogging bool) (*nats.Conn, *server.Server, error) {
	opts := &server.Options{
		//DontListen: inProcess,
	}

	// log.Println("Starting NATS")
	ns, err := server.NewServer(opts)
	if err != nil {
		return nil, nil, err
	}

	// log.Println("NATS Configure")
	if enableLogging {
		ns.ConfigureLogger()
	}
	go ns.Start()

	if !ns.ReadyForConnections(5 * time.Second) {
		return nil, nil, errors.New("NATS Server timeout")
	}

	clientOpts := []nats.Option{}
	if inProcess {
		clientOpts = append(clientOpts, nats.InProcessServer(ns))
	}

	nc, err := nats.Connect(ns.ClientURL(), clientOpts...)
	if err != nil {
		return nil, nil, err
	}

	return nc, ns, err
}
