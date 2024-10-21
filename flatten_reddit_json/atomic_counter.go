package main

import (
	"sync/atomic"
	"time"
	log "github.com/sirupsen/logrus"
)

type AtomicCounter struct {
	counter uint64
	paused uint64
}

func NewCounter() *AtomicCounter {
	ac := AtomicCounter{}
	ac.counter = 0
	ac.paused = 0
	return &ac
}

func (ac *AtomicCounter) isAllowed() bool {
	currPaused := atomic.LoadUint64(&ac.paused)
	if currPaused == 1 {
		return false
	} else {
		return true
	}
}

func (ac *AtomicCounter) setPaused () {
	log.Info("set paused, going to sleep...")
	atomic.StoreUint64(&ac.paused, 1)
	time.Sleep(60 * time.Second)
	log.Info("waking up from pause...")
	atomic.StoreUint64(&ac.paused, 0)
}

func (ac *AtomicCounter) addCounter () {
	atomic.AddUint64(&ac.counter, 1)
	if atomic.LoadUint64(&ac.counter) % 3000 == 0 {
		ac.setPaused()
	}
}


func (ac *AtomicCounter) readCounter() uint64 {
	return atomic.LoadUint64(&ac.counter)
}

func (ac *AtomicCounter) resetCounter() {
	atomic.StoreUint64(&ac.counter, 0)
}
