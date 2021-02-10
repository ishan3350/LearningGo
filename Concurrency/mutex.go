package main
// solving race condition using mutex

// check using below command

// go run -race mutex.go

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Num CPUs:", runtime.NumCPU())
	fmt.Println("Num Goroutines", runtime.NumGoroutine())

	counter := 0

	const max = 100
	var mute sync.Mutex

	wg.Add(100)

	for i := 0; i < max; i++ {
		//wg.Add(1)
		go func() {
			mute.Lock()
			//defer wg.Done()
			fmt.Println(counter)
			counter++
			mute.Unlock()
			wg.Done()

		}()

	}
	wg.Wait()

	fmt.Println("Counter Value:", counter)

	fmt.Println("Num Goroutines", runtime.NumGoroutine())

}
