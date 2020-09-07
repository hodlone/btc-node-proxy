package bitcoin

import (
	"log"
	"os"

	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/rpcclient"
)

var Client *rpcclient.Client

// InitRPCClient initializes a package instance of the rpc client
func InitRPCClient() {
	log.Printf("HOST :: %v", os.Getenv("BITCOIN_RPC_HOST"))
	log.Printf("USER :: %v", os.Getenv("BITCOIN_RPC_USER"))
	log.Printf("PASSWORD :: %v", os.Getenv("BITCOIN_RPC_PASSWORD"))
	// Connect to local bitcoin core RPC server using HTTP POST mode.
	connCfg := &rpcclient.ConnConfig{
		Host:         "bitcoin-core:8332",
		User:         os.Getenv("BITCOIN_RPC_USER"),
		Pass:         os.Getenv("BITCOIN_RPC_PASSWORD"),
		HTTPPostMode: true, // Bitcoin core only supports HTTP POST mode
		DisableTLS:   true, // Bitcoin core does not provide TLS by default
	}
	// log.Printf("COnfig %v", connCfg)
	// Notice the notification parameter is nil since notifications are
	// not supported in HTTP POST mode.
	client, err := rpcclient.New(connCfg, nil)
	if err != nil {
		log.Fatal(err)
	}
	Client = client
	// defer client.Shutdown()
	c, _ := Client.GetConnectionCount()
	log.Printf("Conn Count %v", c)
}

func GetBlockCount() int64 {
	// Get the current block count.
	blockCount, err := Client.GetBlockCount()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Block count: %d", blockCount)
	return blockCount
}

func ListUnspent() []btcjson.ListUnspentResult {
	// Get the current block count.
	unspents, err := Client.ListUnspent()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(unspents)
	return unspents
}
