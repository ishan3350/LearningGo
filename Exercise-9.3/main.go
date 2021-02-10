package main

// Note: This program is creating data race

import (
	"fmt"
	"runtime"
	"sync"
)
var counter int = 0

var wg sync.WaitGroup

func main(){


	wg.Add(100)
	for i := 0; i<100;i++{
		go increamentor()
	}

	wg.Wait()
	fmt.Println(counter)

	
}

func increamentor(){
	counter++
	fmt.Println("Num Goroutine:",runtime.NumGoroutine())
	wg.Done()
}