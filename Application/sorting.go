package main

import (
	"fmt"
	"sort"
)

// sorting slice of string

func main() {
	data := []string{"Alpha", "Delta", "Bravo", "Charlie"}
	sort.Strings(data)

	fmt.Println(data)

	numData := []int{9,1,10,2,3,4,5,6,7,8}
	sort.Ints(numData)

	fmt.Println(numData)
}
