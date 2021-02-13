package cache

import (
	"strconv"
)

var bitcoinBlockNumberKey = "bitcoin:block:number"

//PutBitcoinBlockNumber ...
func (redis *Redis) PutBitcoinBlockNumber(blockN int64) error {
	var err error

	encodedData := strconv.FormatInt(blockN, 10)

	if err = redis.Set(bitcoinBlockNumberKey, encodedData); err != nil {
		return err
	}

	return nil
}

//GetBitcoinBlockNumber ...
func (redis *Redis) GetBitcoinBlockNumber() (int64, error) {
	var err error

	var encodedData string
	if encodedData, err = redis.Get(bitcoinBlockNumberKey); err != nil {
		return 0, err
	}

	// convert encoded block number to int64
	var blockN int64
	if blockN, err = strconv.ParseInt(encodedData, 10, 64); err != nil {
		return 0, err
	}

	return blockN, nil
}
