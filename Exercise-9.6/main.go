package main

// program for displaying OS and Architecture
import (
	"fmt"
	"runtime"
)

func main(){
	//fmt.Println("Hello World!")

	fmt.Println("OS: ",runtime.GOOS)
	fmt.Println("Architecture: ",runtime.GOARCH)

}