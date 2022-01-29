package concur

import (
	"fmt"
	"sync"
)

var (
	count int
	mu    sync.Mutex
)

// Increament with lock
func increment(wg *sync.WaitGroup) {
	mu.Lock()
	count++
	mu.Unlock()
	wg.Done()
}

// Dispatch threads with waitgroup
func dispatchThreads() {
	var count = 5
	var wg sync.WaitGroup
	wg.Add(count)
	for i := 0; i < count; i++ {
		go increment(&wg)
	}

	fmt.Println(count)
}
