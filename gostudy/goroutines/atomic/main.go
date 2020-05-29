package main

import (
	"fmt"
	"sync"
	"time"
)

// var rwlock sync.RWMutex
var mutex sync.Mutex
var x int32
var wg sync.WaitGroup

func addMutex() {
	for i := 0; i < 50000; i++ {
		// rwlock.Lock()
		mutex.Lock()
		x = x + 1
		// atomic.AddInt32(&x, 1)
		// time.Sleep(time.Millisecond * 10)
		mutex.Unlock()
		// rwlock.Unlock()
	}

	wg.Done()

}

func main() {
	start := time.Now().UnixNano()
	for i := 0; i < 500; i++ {
		wg.Add(1)
		// go add()
		go addMutex()
	}

	wg.Wait()
	end := time.Now().UnixNano()
	cost := (end - start) / 1000 / 1000
	fmt.Println("x: ", x, "cost: ", cost, "ms")

}
