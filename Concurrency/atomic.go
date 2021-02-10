package main

// solving race condition using mutex

// check using below command

// go run -race mutex.go

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Num CPUs:", runtime.NumCPU())
	fmt.Println("Num Goroutines", runtime.NumGoroutine())

	var counter int64

	const max = 100

	wg.Add(100)

	for i := 0; i < max; i++ {
		//wg.Add(1)
		go func() {
			atomic.AddInt64(&counter, 1)
			//defer wg.Done()
			fmt.Println(atomic.LoadInt64(&counter))
			wg.Done()

		}()

	}
	wg.Wait()

	fmt.Println("Counter Value:", counter)

}
