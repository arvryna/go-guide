package concur

/*
	This Program demostrates changing a shared variable in runtime,
	and protecting it with the help of MutexLock
*/

import (
	"fmt"
	"sync"
)

var (
	counter int
	mu      sync.RWMutex
)

func Increment(wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()
	counter++
	fmt.Println("writing: ", counter)
	mu.Unlock()
}

func Decrementer(wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()
	counter--
	fmt.Println("writing: ", counter)
	mu.Unlock()
}

func ReadCounter(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("read request")
	fmt.Println(">reading: ", counter)
	fmt.Println("read complete")
}

// Dispatch threads with waitgroup
func DispatchThreads() {
	counter = 0
	var reps = 10
	var wg sync.WaitGroup
	for i := 0; i < reps; i++ {
		wg.Add(1)
		go Increment(&wg)

		wg.Add(1)
		go ReadCounter(&wg)
	}
	wg.Wait()
	fmt.Println(counter)
}

// func main() {
// 	start := time.Now()
// 	DispatchThreads()
// 	elapsed := time.Since(start)
// 	log.Printf("Binomial took %s", elapsed)
// }
