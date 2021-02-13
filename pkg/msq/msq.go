package msq

import (
	"log"
	"sync"

	"github.com/NodeHodl/btc-node-proxy/config"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
)

// MSQ ...
type MSQ struct {
	Nats            *nats.Conn
	Stan            stan.Conn
	SubjectHandlers map[string]func(*stan.Msg)
}

// New contains access to nats and stan types
func New() *MSQ {
	log.Printf("Connecting to NATS client at: %v as %v", config.NatsAddr, config.NatsName)

	NATS, err := nats.Connect(config.NatsAddr, nats.Name(config.NatsName))

	if err != nil {
		log.Fatalf("NATS :: %v", err)
	}

	STAN, err := stan.Connect(config.StanClusterID, config.ServiceName, stan.NatsConn(NATS))
	if err != nil {
		log.Fatalf("STAN :: %v", err)
	}

	log.Printf("Connected to NATS client at: %v as %v", config.NatsAddr, config.NatsName)

	return &MSQ{
		Nats:            NATS,
		Stan:            STAN,
		SubjectHandlers: map[string]func(*stan.Msg){},
	}
}

// Publish ...
func (msq *MSQ) Publish(s string, m []byte) {
	guid, err := msq.Stan.PublishAsync(s, m, func(lguid string, err error) {
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

// AddHandler ...
func (msq *MSQ) AddHandler(subject string, handler func(m *stan.Msg)) {
	msq.SubjectHandlers[subject] = handler
}

// Run ...
func (msq *MSQ) Run() {
	// Use a WaitGroup to wait for a message to arrive
	wg := sync.WaitGroup{}
	defer wg.Done()

	for subject, handler := range msq.SubjectHandlers {
		wg.Add(1)

		// Implement support for groups?
		_, err := msq.Stan.QueueSubscribe(subject, "group1", handler)
		if err != nil {
			log.Fatal(err)
		}
	}
	// Wait for a message to come in
	wg.Wait()
}
