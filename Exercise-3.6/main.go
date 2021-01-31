package main

import(
	"fmt"
)

func main(){
	comparenumber()
}

func comparenumber(){
	a := 10
	b := 20

	if a>b{
		fmt.Println("a is greater than b")
	}else if a<b{
		fmt.Println("a is less than b")
	}else{
		fmt.Println("a is equal to b")
	}
}