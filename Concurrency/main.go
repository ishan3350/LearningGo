package main

// concurrency example
// goroutine or goroutines

import (
	"fmt"
	"runtime"
	"sync"
)

func init() {
	fmt.Println("This function will execute before main function")

}

var wg sync.WaitGroup

func main() {

	fmt.Println("OS:\t", runtime.GOOS)
	fmt.Println("Architecture:\t", runtime.GOARCH)
	fmt.Println("CPUs:\t", runtime.NumCPU())
	fmt.Println("Goroutines:\t", runtime.NumGoroutine())

	// We will wait for 2 things and program will exit and when compiler finds two wg.Done()
	wg.Add(2)
	go foo()
	go bar()
	fmt.Println("CPUs:\t", runtime.NumCPU())
	fmt.Println("Goroutines:\t", runtime.NumGoroutine())
	wg.Wait()
	// when we use go keyword it creates a new goroutine and program continues to execute and specified function after go keyword will execute on another go routine. Your program may exit once it finishes the current goroutine
}

func foo() {
	for i := 0; i <= 10; i++ {
		fmt.Println("foo: ", i)
	}
	// completing first wg wait
	wg.Done()
}

func bar() {
	for i := 0; i <= 10; i++ {
		fmt.Println("bar: ", i)
	}
	// completing second wg wait so program can exit
	wg.Done()
}
