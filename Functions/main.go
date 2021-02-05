package main

import(
  "fmt"
)

func main(){
  // fmt.Println("Hello World!")

  sayhello()

  displaymessage("Hello World!")

  //We can call function as many times as we want

  displaymessage("Hello again!")

  fmt.Println(add(10,20))

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
