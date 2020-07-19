package qeue

import (
	"log"

	"github.com/nats-io/nats.go"
)

// NATS export connection pointer to nats
var NATS *nats.Conn

// Start a connection with a nats instance and return a pointer to it
func Start(natsAddr string, natsName string) {
	// connecting to a test server, replace the arg for the cluster conn string
	// nats://<kubernetes-svc>.<namespace>:port
	// add a name to the connection for monitoring/debuggin purposes
	NATS, err := nats.Connect(natsAddr, nats.Name(natsName))
	if err != nil {
		log.Fatal(err)
	}
	// defer connection close
	defer NATS.Close()
	log.Printf("Connected to NATS client at: %v as %v", natsAddr, natsName)
}

// Qpub ...
func Qpub(s string, m []byte) {
	NATS.Publish(s, m)
	NATS.Flush()
	if err := NATS.LastError(); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Published [%s] : '%s'\n", s, m)
	}
}
