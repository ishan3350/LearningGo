// print every number from 1 to 10,0000

package main

import (
	"fmt"
)

func main() {

	displaynumbers()
}

func displaynumbers() {
	for i := 1; i <= 10000; i++ {
		fmt.Println(i)
	}
}
