package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// creating file and writing data and checking for error

func main(){
	f, err := os.Create("test.txt")

	if err == nil{
		fmt.Println("test.txt created")
		defer f.Close()
		data := strings.NewReader("Hello World!\n")
		io.Copy(f,data)
	}else{
		fmt.Println(err)
		return
	}
}