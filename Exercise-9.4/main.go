package main

// Note: Solving data race from previous exercise

import (
	"fmt"
	"runtime"
	"sync"
)

var counter int = 0

var wg sync.WaitGroup

var mute sync.Mutex

func main() {

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go increamentor()
	}

	wg.Wait()
	fmt.Println("Counter: ",counter)

}

func increamentor() {
	mute.Lock()
	counter++
	fmt.Println("Num Goroutine:", runtime.NumGoroutine())
	mute.Unlock()
	wg.Done()
}
