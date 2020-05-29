package main

import (
	"fmt"
	"sync"
	"time"
)

var rwlock sync.RWMutex
var mutex sync.Mutex
var x int
var wg sync.WaitGroup

func write() {
	for i := 0; i < 100; i++ {
		// rwlock.Lock()
		mutex.Lock()
		x = x + 1
		time.Sleep(time.Millisecond * 10)
		mutex.Unlock()
		// rwlock.Unlock()
	}

	wg.Done()

}

func read() {
	for i := 0; i < 100; i++ {
		mutex.Lock()
		// rwlock.RLock()
		time.Sleep(time.Millisecond) // 花费1毫秒
		// rwlock.RUnlock()
		mutex.Unlock()
	}

	wg.Done()
}

func main() {
	start := time.Now().UnixNano()
	wg.Add(1)
	go write()

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go read()
	}

	wg.Wait()
	end := time.Now().UnixNano()
	cost := (end - start) / 1000 / 1000
	fmt.Println("cost: ", cost, "ms")

}
