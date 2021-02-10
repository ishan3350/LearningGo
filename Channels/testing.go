package main

import "fmt"

var nums chan int

func main() {

	nums = make(chan int)

	go foo()

	for num := range nums {
		fmt.Println(num)
	}

}

func foo() {
	for i := 0; i <= 100; i++ {
		nums <- i

	}
	close(nums)
}
