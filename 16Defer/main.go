package main

import "fmt"

func main() {
	defer fmt.Println("We are going to learn defer!")
	defer fmt.Println("two")
	defer fmt.Println("one")
	defer fmt.Println(" zero")
	fmt.Println("Welcome to defer")
	ourDefer()
}

func ourDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}

// welcome to defer, 4,3,2,1,0 , zero, one, two, we are going to learn defer!
