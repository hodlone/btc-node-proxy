package main

import (
	"github.com/NodeHodl/btc-node-proxy/bitcoin"
	"log"
	"sync"
	"time"
)

type task struct {
	job      func()
	interval time.Duration
}

func (t *task) schedule() {
	ticker := time.NewTicker(t.interval)
	done := make(chan bool)

	func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				t.job()
			}
		}
	}()

	done <- true
	log.Println("Ticker stopped")
}

// StartNetworkMonitor ...
func StartNetworkMonitor() {
	testTask := task{job: func() {
		log.Printf("Rutine test task")
	},
		interval: 240 * time.Second,
	}

	taskTwo := task{job: func() {
		diff, _ := bitcoin.Client.GetDifficulty()
		bc, _ := bitcoin.Client.ListUnspent()
		log.Printf("Difficulty %v", diff)
		log.Printf("Block Count %v", bc)
	},
		interval: 1000 * time.Second,
	}

	wg := new(sync.WaitGroup)

	wg.Add(2)
	go testTask.schedule()
	go taskTwo.schedule()
	wg.Wait()
}
