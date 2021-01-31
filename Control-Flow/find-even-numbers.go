package main

import (
	"fmt"
)

func main() {
	findevennumbers()

}

func findevennumbers() {
	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Printf("\n%d is even number", i)
		}
	}
}
