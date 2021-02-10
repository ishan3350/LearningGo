package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	c := make(chan int)
	wg.Add(1)
	go foo(c)
	go bar(c)
	wg.Wait()
	fmt.Println("program is about to exit")

}

func foo(c chan int) {
	for i := 0; i <= 100; i++ {
		c <- i
	}
}
func bar(c chan int) {
	for i := 0; i <= 100; i++ {
		fmt.Println(<-c)
	}
	wg.Done()

}
