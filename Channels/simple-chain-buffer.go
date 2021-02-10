package main

import(
	"fmt"
)

func main(){
	// specifying buffer size of 1
	// c := make(chan int,1)

	// // Here we are putting two values in channel with buffer size of 1 that allows to store only one buffer value and let it sit and not pulling value of so we are not able to store another value 45
	// c <- 40
	// c <- 45
	// fmt.Println(<-c)


	// solving this problem


	ch := make(chan int,1)

	ch <- 40
	fmt.Println(<-ch)
	ch <- 45

	fmt.Println(<-ch)

	// another good example

	ch1 := make(chan int,2)

	ch1 <- 40
	ch1 <- 45

	fmt.Println(<-ch1)
	fmt.Println(<-ch1)

	// Rule of thumb is that if we specify buffer size in argument of make function then we can store that many values without pulling other values. If you need to store more than specified value then first pull old values and then put new values
	

}