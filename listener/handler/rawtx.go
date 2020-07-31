package handler

import "btc-node-proxy/qeue"

// RawTx handles the zmq messages published to the "rawtx" socket topic.
func RawTx(msg []byte) {
	qeue.Qpub("btc.rawtx", msg)
	qeue.Qpub("tele.test", []byte("Moon Soon!"))
}

// Example of Transaction decoding using "github.com/btcsuite/btcd/wire"
//
// func RawTx(msg []byte) {
// 	// qeue.Qpub("btc.rawtx", msg)
//
// 	//RawBlock deserialize the block and return it.
// 	var msgTx wire.MsgTx
//
// 	err := msgTx.Deserialize(bytes.NewReader(msg))
//
// 	printTx(msgTx)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// func printTx(tx wire.MsgTx) {
// 	log.Printf("### TX ###")
// 	log.Printf("Tx Version:            %v", tx.Version)
// 	log.Printf("Tx LockTime:           %v", tx.LockTime)
//
// 	for j, out := range tx.TxOut {
// 		log.Printf("### OUT%v ###", j)
// 		log.Printf("Tx Value BTC:          %v", out.Value)
// 	}
// }
