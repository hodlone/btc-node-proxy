package msq

import (
	"log"
	"os"
	"sync"

	"github.com/nats-io/nats.go"
)

var (
	natsAddr = os.Getenv("NATS_ADDR")
	natsName = os.Getenv("NATS_NAME")
)

// Start a connection with a nats instance and return a pointer to it
func Start() {
	log.Printf("Connecting to NATS client at: %v as %v", natsAddr, natsName)
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
	wg := new(sync.WaitGroup)

	wg.Add(1)

	NATS, err := nats.Connect(natsAddr, nats.Name(natsName))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Pubbingg!")
	// defer connection close
	defer NATS.Close()
	go NATS.Publish(s, m)
	wg.Done()
	NATS.Flush()
	if err := NATS.LastError(); err != nil {
		log.Fatal(err)
	}
}
