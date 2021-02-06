package main

import (
	"fmt"
)

func main() {

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	result := foo(nums...)

	fmt.Println(result)

}

func foo(n ...int) int {

	total := 0
	for num := range n {
		total = total + num

	}
	return total

}
