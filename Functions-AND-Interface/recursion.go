package main

import(
	"fmt"
)

func main(){
	
	fmt.Println(factorial(7))

}

func factorial(n int)int{

	// recusrsion is function calls it self
	if n==0{
		return 1
	}else{
		return n* factorial(n-1)
	}

}