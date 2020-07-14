package nats

import (
	"log"

	"github.com/nats-io/nats.go"
)

// NATS export connection pointer to nats
var NATS *nats.Conn

// Start a connection with a nats instance and return a pointer to it
func Start(natsAddr string, natsName string) {
	// connecting to a test server, replace the arg for the cluster conn string
	// nats://<namespace>.<address>:port
	// add a name to the connection for monitoring/debuggin purposes
	NATS, err := nats.Connect(natsAddr, nats.Name(natsName))
	if err != nil {
		log.Fatal(err)
	}
	// defer connection close
	defer NATS.Close()
}
