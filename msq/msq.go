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

	if err != nil {
		log.Fatalf("NATS :: %v", err)
	}
	natsConn = NATS

	sc, err := stan.Connect(clusterID, clientID, stan.NatsConn(natsConn))
	if err != nil {
		log.Fatalf("STAN :: %v", err)
	}

	stanClient = sc

	log.Printf("Connected to NATS client at: %v as %v", natsAddr, natsName)
}

// Publish ...
func Publish(s string, m []byte) {
	guid, err := stanClient.PublishAsync(s, m, func(lguid string, err error) {
		if err != nil {
			log.Fatalf("Error in server ack for guid %s: %v\n", lguid, err)
		}
	})
	if err != nil {
		log.Fatalf("Error during publish: %v\n", err)
	}

	if guid == "" {
		log.Fatal("Expected non-empty guid to be returned.")
	}
}
