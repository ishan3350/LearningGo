package main

import (
	"fmt"
	"runtime"
	"sync"
)
var wg sync.WaitGroup

func main(){
	fmt.Println("Hello World!")
	wg.Add(2)
	go displayMsg()
	go displayMsg()
	wg.Wait()
}

func displayMsg(){
	fmt.Println("Hello from DisplayMsg")
	fmt.Println("Number of go routines:",runtime.NumGoroutine())
	wg.Done()
}