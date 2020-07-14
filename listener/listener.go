package listener

import (
	"encoding/hex"
	"log"

	"github.com/pebbe/zmq4"
)

func decode(src []byte) {
	// log.Printf("decoding length: %v", len(src))
	dst := make([]byte, hex.DecodedLen(len(src)))
	log.Printf("dst: %v", dst)
	log.Printf("HEX: %v", hex.EncodeToString(src))
	// n, err := hex.Decode(dst, src)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("%s\n", dst[:n])
}

func Start(btcNodeZmqAddr string) {
	subscriber, _ := zmq4.NewSocket(zmq4.SUB)
	subscriber.SetLinger(0)

	subscriber.Connect(btcNodeZmqAddr)

	subscriber.SetSubscribe("hashblock")
	// subscriber.SetSubscribe("hashtx")
	// subscriber.SetSubscribe("rawblock")
	// subscriber.SetSubscribe("rawtx")

	log.Println("Listening")
	for {
		e, addr, v, err := subscriber.RecvEvent(0)
		if err != nil {
			log.Println(err)
			continue
		}

		log.Printf("Event: %v", e)
		addrs := []byte(addr)
		log.Printf("Address: %v", addrs)
		decode(addrs)
		log.Printf("Value: %v", v)
		// bl, er := btc.NewBlock([]byte(addr))

		// if er != nil {
		// 	log.Println(err)
		// 	continue
		// }

		// log.Printf("Block hash: %v", bl.Hash)
		// log.Printf("Block time: %v", bl.BlockTime)
		// log.Printf("Block version: %v", bl.Version)
		// log.Printf("Block parent: %v", btc.NewUint256(bl.Parent).String())
		// log.Printf("Block merkle root: %v", hex.EncodeToString(bl.MerkleRoot))
		// log.Printf("Block bits: %v", bl.Bits)
		// log.Printf("Block size: %v", len(bl.Raw))
		// decode([]byte(s))
		// log.Println("rec", s)
	}
}
