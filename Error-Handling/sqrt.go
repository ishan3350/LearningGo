package main

// throwing error in function or returning error
import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Hello World!")

	err, result := sqrt(5)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	err, result = sqrt(-1)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

func sqrt(n int) (error, float64) {
	if n < 0 {
		return errors.New("Sqaure root of negative number resulted in error"), 0
	} else {
		return nil, float64(n * n)
	}
}
