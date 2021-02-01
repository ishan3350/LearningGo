package main

import (
	"fmt"
)

func main() {
	//simpleslice()

	// creatingsliceusingmake()

	multidimensionalslice()
}

func simpleslice() {

	// it is good practice to use slice instead of arrays

	// creation of slice using composite literal. Slice allows to group together values of the same type
	x := []int{1, 2, 3, 4, 5}
	fmt.Println(x)

	strs := []string{"test1", "test2", "test3", "test4", "test5"}
	fmt.Println(strs)

	// looping through slice using range

	for key, value := range x {
		fmt.Println(key, value)
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

	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
	}

	// appending to the slice

	x = append(x, 6, 7, 8)
	fmt.Println(x)

	// creating another slice using compsite literal

	y := []int{10, 11, 12, 13, 14, 15}
	fmt.Println(y)

	// appending two slices

	// ... tells compiler to take all elements from the slice and append it
	x = append(x, y...)
	fmt.Println(x)

	// deleting elements from the slice

	// original slice [1 2 3 4 5 6 7 8 10 11 12 13 14 15]
	//deleting elements 2 and 3
	x = append(x[:1], x[3:]...)
	fmt.Println(x)

}

func creatingsliceusingmake() {

	// slice is built on top of array so when the size of slice grows it will create another array and throw the old one out from memory so it requires processing power

	// if you know the size of slice in advance then it's good practice to create a slice using make function

	/*

			https://stackoverflow.com/questions/41668053/cap-vs-len-of-slice-in-golang#:~:text=The%20length%20of%20a%20slice,first%20element%20in%20the%20slice.
			A slice is an abstraction that uses an array under the covers.

		cap tells you the capacity of the underlying array. len tells you how many items are in the array.

		The slice abstraction in Go is very nice since it will resize the underlying array for you, plus in Go arrays cannot be resized so slices are almost always used instead.

		Example:

		s := make([]int, 0, 3)
		for i := 0; i < 5; i++ {
		    s = append(s, i)
		    fmt.Printf("cap %v, len %v, %p\n", cap(s), len(s), s)
		}
		Will output something like this:

		cap 3, len 1, 0x1040e130
		cap 3, len 2, 0x1040e130
		cap 3, len 3, 0x1040e130
		cap 6, len 4, 0x10432220
		cap 6, len 5, 0x10432220
		As you can see once the capacity is met, append will return a new slice with a larger capacity. On the 4th iteration you will notice a larger capacity and a new pointer address.

		Play example

		I realize you did not ask about arrays and append but they are pretty foundational in understanding the slice and the reason for the builtins.

		https://play.golang.org/p/RhhsGgFYI1

	*/

	x := make([]int, 10, 15)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x[0] = 1
	x[1] = 2
	x[2] = 3
	x[3] = 4
	x[4] = 5

	fmt.Println(x)

}

func multidimensionalslice() {


	// Multi dimensional slice
	num1 := []int{1, 2, 3, 4, 5}
	num2 := []int{6, 7, 8, 9, 10}

	data := [][]int{num1, num2}
	fmt.Println(data)

}
