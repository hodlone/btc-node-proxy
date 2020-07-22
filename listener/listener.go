package listener

import (
	handler "btc-node-proxy/listener/handler"
	"log"

	"github.com/pebbe/zmq4"
)

func Start(btcNodeZmqAddr string) {
	subscriber, _ := zmq4.NewSocket(zmq4.SUB)
	subscriber.SetLinger(0)

	subscriber.Connect(btcNodeZmqAddr)

	subscriber.SetSubscribe("hashblock")
	subscriber.SetSubscribe("hashtx")
	subscriber.SetSubscribe("rawblock")
	subscriber.SetSubscribe("rawtx")

	log.Println("Listening")
	for {
		frames, err := subscriber.RecvMessageBytes(0)
		if err != nil {
			log.Println(err)
			continue
		}

		topic := string(frames[0])
		body := frames[1:]

		// log.Printf("body: %v", body[1])
		// decode(body[1])

		log.Printf("topic: %v", topic)
		switch topic {
		case "hashblock":
			handler.HashBlock(body[0])
		case "rawblock":
			handler.RawBlock(body[0])
			// case "hashtx":
			// hexi := hex.EncodeToString(body[0])
			// log.Printf("From HashTx HEX: %v", hexi)
			// handler.HashTx(topic)
			// case "rawtx":
			// tx, _ := btc.NewTxIn(body[0])
			// log.Printf("From RawTx Hash: %v", tx.Sequence)
			// handler.RawTx(topic)
		}
		// if msg == "rawblock" {
		// 	continue
		// }

		// addrs := []byte(msg)
		// log.Printf("BYTE: %v", addrs)
		// decode(addrs)

		// log.Printf("Block Raw: %v", bl.Raw)
		// log.Printf("Block TxCount: %v", bl.TxCount)
		// log.Printf("Block Txs: %v", bl.Txs)

		// for _, tx := range bl.Txs {
		// 	log.Printf("Tx Hash: %v", tx.Hash)
		// 	for _, output := range tx.TxOut {
		// 		s := output.String(true)
		// 		log.Printf("Tx Out Value: %v", s)
		// 	}
		// type Tx struct {
		// 	Version   uint32
		// 	TxIn      []*TxIn
		// 	TxOut     []*TxOut
		// 	SegWit    [][][]byte
		// 	Lock_time uint32

		// 	// These three fields should be set in block.go:
		// 	Raw             []byte
		// 	Size, NoWitSize uint32
		// 	Hash            Uint256

		// 	// This field is only set in chain's ProcessBlockTransactions:
		// 	Fee uint64

		// 	wTxID Uint256

		// 	hash_lock    sync.Mutex
		// 	hashPrevouts []byte
		// 	hashSequence []byte
		// 	hashOutputs  []byte
		// }
		// }
		// log.Printf("Block time: %v", bl.BlockTime)
		// log.Printf("Block version: %v", bl.Version)

		// type Block struct {
		// 	Raw               []byte
		// 	Hash              *Uint256
		// 	Txs               []*Tx
		// 	TxCount, TxOffset int  // Number of transactions and byte offset to the first one
		// 	Trusted           bool // if the block is trusted, we do not check signatures and some other things...
		// 	LastKnownHeight   uint32

		// 	BlockExtraInfo // If we cache block on disk (between downloading and comitting), this data has to be preserved

		// 	MedianPastTime uint32 // Set in PreCheckBlock() .. last used in PostCheckBlock()

		// 	// These flags are set in BuildTxList() used later (e.g. by script.VerifyTxScript):
		// 	NoWitnessSize int
		// 	BlockWeight   uint
		// 	TotalInputs   int

		// 	NoWitnessData []byte // This is set by BuildNoWitnessData()
		// }
	}
}
