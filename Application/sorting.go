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
}
