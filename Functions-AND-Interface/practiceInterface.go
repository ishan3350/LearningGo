package main

import (
	"fmt"
)

func main() {
	displayData(10)
	displayData("Hello World")
}

// it will allow us to pass anytype of value
// In empty interface we can pass any type
func displayData(a interface{}) {
	fmt.Println(a)
}
