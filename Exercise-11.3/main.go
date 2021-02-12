package main

import(
	"fmt"
)

type customerror struct{
	info string
}

func (ce customerror) Error() string{
	return fmt.Sprintf("here is the error: %v",ce.info)
}
func main(){
	
	c1 := customerror{
		info: "error1",
	}

	foo(c1)
}

func foo(e error){
	fmt.Println(e)

	//assertion
	fmt.Println(e.(customerror).info)
}