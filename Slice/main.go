package main

import(
	"fmt"
)

func main(){
	simpleslice()
}

func simpleslice(){

	// it is good practice to use slice instead of arrays

	// creation of slice using composite literal. Slice allows to group together values of the same type
	x := []int{1,2,3,4,5}
	fmt.Println(x)


	strs := []string{"test1","test2","test3","test4","test5"}
	fmt.Println(strs)

	// looping through slice using range

	for key,value := range x{
		fmt.Println(key,value)
	}
}