package bitcoin

import (
	"log"
	"os"

	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/rpcclient"
)

var bitcoinClient *rpcclient.Client

// InitRPCClient initializes a package instance of the rpc client
func InitRPCClient() {
	// Connect to local bitcoin core RPC server using HTTP POST mode.
	connCfg := &rpcclient.ConnConfig{
		Host:         os.Getenv("BITCOIN_RPC_HOST"),
		User:         os.Getenv("BITCOIN_RPC_USER"),
		Pass:         os.Getenv("BITCOIN_RPC_PASSWORD"),
		HTTPPostMode: true, // Bitcoin core only supports HTTP POST mode
		DisableTLS:   true, // Bitcoin core does not provide TLS by default
	}
	// Notice the notification parameter is nil since notifications are
	// not supported in HTTP POST mode.
	client, err := rpcclient.New(connCfg, nil)
	if err != nil {
		log.Fatal(err)
	}
	bitcoinClient = client
	// defer client.Shutdown()
}

func GetBlockCount() int64 {
	// Get the current block count.
	blockCount, err := bitcoinClient.GetBlockCount()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Block count: %d", blockCount)
	return blockCount
}

func ListUnspent() []btcjson.ListUnspentResult {
	// Get the current block count.
	unspents, err := bitcoinClient.ListUnspent()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(unspents)
	return unspents
}
