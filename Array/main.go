package main

import(
	"fmt"
)

func main(){

	arrayexample()


}

func arrayexample(){

// declaration of array
var x [5]int

x[3] = 42
fmt.Println(x)

//finding size of array length
fmt.Println(len(x))

}