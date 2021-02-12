package main

import(
	"fmt"
)

func main(){
	n,err := fmt.Println("Hello World!")

	if err == nil{
		fmt.Println(n)
	}else{
		fmt.Println(err)
	}

	// scanning data and checking for error

	var data string

	_,err = fmt.Scan(&data)
	
	if(err == nil){
		fmt.Println(data)
	}else{
		fmt.Println(err)
	}
}