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

	// slicing the slice

	// displays the whole slice
	fmt.Println(x[:])

	// starts from index position 1 and go upto 3 but index position 3 will be not included
	fmt.Println(x[1:3])
	
	// starts from position 0 and go upto 3 but index position 3 will be not included
	fmt.Println(x[:3])

	// starts from index position 3 and include everything
	fmt.Println(x[3:])

	// accessing elements using for loop (Please don't use range)

	for i:=0;i<len(x);i++{
		fmt.Println(x[i])
	}



}