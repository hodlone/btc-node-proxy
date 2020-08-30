package msq

var (
// natsAddr = os.Getenv("NATS_ADDR")
// natsName = os.Getenv("NATS_NAME")
// conn     *nats.Conn
)

// Start a connection with a nats instance and return a pointer to it
func Start() {
	// log.Printf("Connecting to NATS client at: %v as %v", natsAddr, natsName)
	// connecting to a test server, replace the arg for the cluster conn string
	// nats://<kubernetes-svc>.<namespace>:port
	// add a name to the connection for monitoring/debuggin purposes
	// NATS, err := nats.Connect(natsAddr, nats.Name(natsName))

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer connection close
	// defer NATS.Close()l
	// conn = NATS
	// log.Printf("Connected to NATS client at: %v as %v", natsAddr, natsName)
}

// Qpub ...
func Qpub(s string, m []byte) {
	// NATS, err := nats.Connect(natsAddr, nats.Name(natsName))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer connection close
	// defer NATS.Close()
	// conn.Publish(s, m)
	// NATS.Flush()

	// if err := conn.LastError(); err != nil {
	// 	log.Fatal(err)
	// }
}
