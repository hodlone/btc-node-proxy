package txsync

import (
	"github.com/NodeHodl/btc-node-proxy/pkg/bitcoin"
	"github.com/NodeHodl/btc-node-proxy/pkg/logger"
	"github.com/NodeHodl/btc-node-proxy/pkg/msq"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
)

var log = logger.New("txsync")

//Start ...
func Start() {
	msqClient := msq.New()
	msqClient.AddHandler("btc.node.block.hash", blockHashHandler)
	msqClient.Run()
}

func blockHashHandler(m *stan.Msg) {
	var err error
	var bitcoinRPCClient = bitcoin.NewBitcoinRPCClient()

	var blockHash *chainhash.Hash
	blockHash.SetBytes(m.Data)

	var block *wire.MsgBlock
	if block, err = bitcoinRPCClient.GetBlock(blockHash); err != nil {
		log.WithFields(logrus.Fields{
			"blockhash": blockHash.String(),
			"error":     err,
		}).
			Errorln("Could not fetch block from bitcoind")

		return
	}

	log.Infof("BlockHash received %v", block.Header.BlockHash().String())
}
