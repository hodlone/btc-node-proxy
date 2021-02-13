package bitcoin

import (
	"os"

	"github.com/NodeHodl/btc-node-proxy/pkg/logger"
	"github.com/btcsuite/btcd/rpcclient"
)

var (
	// Client is the exported RPC interface for the bitcoin client
	Client  *rpcclient.Client
	rpcHost string = os.Getenv("BITCOIN_RPC_HOST")
	rpcUser string = os.Getenv("BITCOIN_RPC_USER")
	rpcPass string = os.Getenv("BITCOIN_RPC_PASSWORD")
)

var log = logger.New("bitcoin-rpc-client")

//NewBitcoinRPCClient initializes a package instance of the rpc client
func NewBitcoinRPCClient() *rpcclient.Client {
	// Connect to local bitcoin core RPC server using HTTP POST mode.
	connCfg := &rpcclient.ConnConfig{
		Host: "bitcoin-core.bitcoin:8332",
		User: rpcUser,
		Pass: "ZCSK5R51slCs0I4LfLbfr7Ultc8l1Lu9taqYK-QIYDc=", // "fIiet_JLr2Dz-MywqmDF2D3vGZCoIu8P2TZ6FpTjoUE="
		// CookiePath:   "bitcoin:50c188112c466e431bd2a2e66ae4013b$d216ce45ed23cac07e9e98b2d20703a0eafb70b0bd677ea3786c876977617689",
		HTTPPostMode: true, // Bitcoin core only supports HTTP POST mode
		DisableTLS:   true, // Bitcoin core does not provide TLS by default
	}

	log.Infof("Connecting Bitcoin RPC interface at %v as %v pw=%v", rpcHost, rpcUser, rpcPass)
	// Notice the notification parameter is nil since notifications are
	// not supported in HTTP POST mode.
	client, err := rpcclient.New(connCfg, nil)
	if err != nil {
		log.Fatal(err)
	}

	return client
	// defer client.Shutdown()
}
