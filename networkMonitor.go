package main

import (
	"btc-node-proxy/bitcoin"
	"fmt"
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
	fmt.Println("Ticker stopped")
}

// StartNetworkMonitor ...
func StartNetworkMonitor() {
	testTask := task{job: func() {
		fmt.Println("Rutine test task")
	},
		interval: 120 * time.Second,
	}

	taskTwo := task{job: func() { bitcoin.GetBlockCount() },
		interval: 5 * time.Second,
	}

	wg := new(sync.WaitGroup)

	wg.Add(2)
	go testTask.schedule()
	go taskTwo.schedule()
	wg.Wait()
}
