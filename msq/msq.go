package msq

import (
	"log"
	"os"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
)

var (
	clusterID  = os.Getenv("STAN_CLUSTER_ID")
	clientID   = os.Getenv("SERVICE_NAME")
	natsAddr   = os.Getenv("NATS_ADDR")
	natsName   = os.Getenv("NATS_NAME")
	natsConn   *nats.Conn
	stanClient stan.Conn
)

// StartStanClient a connection with a nats instance and return a pointer to it
func StartStanClient() {
	log.Printf("Connecting to NATS client at: %v as %v", natsAddr, natsName)

	NATS, err := nats.Connect(natsAddr, nats.Name(natsName))

	natsConn = NATS

	sc, err := stan.Connect(clusterID, clientID, stan.NatsConn(natsConn))
	if err != nil {
		log.Fatal(err)
	}

	stanClient = sc
	log.Printf("Connected to NATS client at: %v as %v", natsAddr, natsName)
}

// Publish ...
func Publish(s string, m []byte) {
	err := stanClient.Publish(s, m)
	if err != nil {
		log.Fatalf("Error during publish: %v\n", err)
	}
}
