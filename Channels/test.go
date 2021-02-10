package main

import "fmt"

var data string

func main() {

	c := make(chan int)
	go foo(c)
	go bar(c)
	fmt.Scan(&data)

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
}
