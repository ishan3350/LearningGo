package main

// Note: Solving race condition using atomic

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var counter int64 = 0

var wg sync.WaitGroup

func main() {

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go increamentor()
	}

	wg.Wait()
	fmt.Println(counter)

}

func increamentor() {
	atomic.AddInt64(&counter,1)
	fmt.Println("Num Goroutine:", runtime.NumGoroutine())
	wg.Done()
}
