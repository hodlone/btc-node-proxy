package bitcoin

import (
	"log"
	"os"

	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/rpcclient"
)

var (
	// Client is the exported RPC interface for the bitcoin client
	Client  *rpcclient.Client
	rpcHost string = os.Getenv("BITCOIN_RPC_HOST")
	rpcUser string = os.Getenv("BITCOIN_RPC_USER")
	rpcPass string = os.Getenv("BITCOIN_RPC_PASSWORD")
)

// InitRPCClient initializes a package instance of the rpc client
func InitRPCClient() {
	// Connect to local bitcoin core RPC server using HTTP POST mode.
	connCfg := &rpcclient.ConnConfig{
		Host:         rpcHost,
		User:         rpcUser,
		Pass:         rpcPass,
		HTTPPostMode: true, // Bitcoin core only supports HTTP POST mode
		DisableTLS:   true, // Bitcoin core does not provide TLS by default
	}

	log.Printf("Connecting Bitcoin RPC interface at %v as %v", rpcHost, rpcUser)
	// Notice the notification parameter is nil since notifications are
	// not supported in HTTP POST mode.
	client, err := rpcclient.New(connCfg, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Succesfully Connected Bitcoin RPC interface")
	Client = client
	// defer client.Shutdown()
	c, _ := Client.GetConnectionCount()
	log.Printf("Conn Count %v", c)
}

// GetBlockCount returns the latest block in the bitcoin blockchain
func GetBlockCount() int64 {
	// Get the current block count.
	blockCount, err := Client.GetBlockCount()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Block count: %d", blockCount)
	return blockCount
}

// ListUnspent returns unspent transactions associated with the target node address
func ListUnspent() []btcjson.ListUnspentResult {
	unspents, err := Client.ListUnspent()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(unspents)
	return unspents
}
