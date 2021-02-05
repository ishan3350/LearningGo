package main

import (
	"fmt"
	"strconv"
)

func main(){
  // fmt.Println("Hello World!")

  sayhello()

  displaymessage("Hello World!")

  //We can call function as many times as we want

  displaymessage("Hello again!")

  fmt.Println(add(10,20))

  resultmsg,result := subtract(20,10)

  fmt.Println(resultmsg)
  fmt.Println(result)

  // Variadic function example (We can pass as many arguments as we want)

  addall(10,20,30)

}

func sayhello(){
  fmt.Println("Hello from function")
}

// function with parameter

func displaymessage(msg string){
  fmt.Println(msg)
}

// function returning value

func add(a int, b int) int{
  return a+b
}

// returning multiple values

func subtract(a int, b int)(int, string){
  ans := "Subtraction of b from a: "+strconv.Itoa(a-b)
  return a-b,ans
}

// function with unlimited parameters (Variadic parameter)

func addall(nums...int){

  // variadic paramter is considered as slice of specified type
  result := 0
  for i:= range nums{
    result = result + nums[i]
    fmt.Println(result)
  }
}