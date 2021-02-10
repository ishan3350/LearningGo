package main

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

	wg.Add(100)
	for i := 0; i < max; i++ {
		//wg.Add(1)
		go func() {
			//defer wg.Done()
			val := counter
			fmt.Println(val)
			val++
			counter = val
			wg.Done()
			
		}()

	}
	
	fmt.Println("Num Goroutines", runtime.NumGoroutine())
	wg.Wait()

}
