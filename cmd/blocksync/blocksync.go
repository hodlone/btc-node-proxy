package blocksync

import (
	"github.com/NodeHodl/btc-node-proxy/pkg/bitcoin"
	"github.com/NodeHodl/btc-node-proxy/pkg/cache"
	"github.com/NodeHodl/btc-node-proxy/pkg/logger"
	"github.com/NodeHodl/btc-node-proxy/pkg/msq"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/jasonlvhit/gocron"
	"github.com/sirupsen/logrus"
)

var log = logger.New("blocksync")

// Start ...
func Start(interval uint64) {
	s := gocron.NewScheduler()

	s.Every(interval).Seconds().Do(syncBlock)

	<-s.Start()
}

// syncBlock ...
func syncBlock() {
	var err error

	// Make RPC call to get block number
	var NodeBlockN int64
	var bitcoinRPCClient = bitcoin.NewBitcoinRPCClient()
	if NodeBlockN, err = bitcoinRPCClient.GetBlockCount(); err != nil {
		log.WithFields(logrus.Fields{
			"error": err,
		}).
			Errorln("Could not fetch block height from bitcoind")

		return
	}

	// Get system blockN from redis
	var SystemBlockN int64
	var redisClient = cache.NewRedisClient()
	if SystemBlockN, err = redisClient.GetBitcoinBlockNumber(); err != nil {
		log.WithFields(logrus.Fields{
			"error":           err,
			"nodeBlockNumber": NodeBlockN,
		}).
			Errorln("Could not fetch block height from redis")

		return
	}

	// Compare system blockN with node blockN,
	if SystemBlockN >= NodeBlockN {
		log.WithFields(logrus.Fields{
			"systemBlockNumber": SystemBlockN,
			"nodeBlockNumber":   NodeBlockN,
		}).
			Errorln("Skipping, blocks already in sync")

		return
	}

	// Get block hashes and publish
	for SystemBlockN < NodeBlockN {
		SystemBlockN++

		var blockHash *chainhash.Hash
		if blockHash, err = bitcoinRPCClient.GetBlockHash(SystemBlockN); err != nil {
			// Failure maybe occur because node is down,
			// exit for loop and wait for next blocksync interval
			log.WithFields(logrus.Fields{
				"error":             err,
				"systemBlockNumber": SystemBlockN,
				"nodeBlockNumber":   NodeBlockN,
			}).
				Errorln("Could not fetch blockhash")

			return
		}

		// Convert blockhash into bytes slice and publish
		msq.New().Publish("btc.node.block.hash", []byte(blockHash.CloneBytes()))

		if err = redisClient.PutBitcoinBlockNumber(SystemBlockN); err != nil {
			// Failure maybe occur because redis is not available,
			// in which case just exit and wait for next blocksync interval
			log.WithFields(logrus.Fields{
				"error":             err,
				"systemBlockNumber": SystemBlockN,
				"nodeBlockNumber":   NodeBlockN,
			}).
				Errorln("Could not store latest block number")

			return
		}
	}
}
