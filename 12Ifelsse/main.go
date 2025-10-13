package main

import "fmt"

func main() {
	fmt.Println("Welcome to main")

	var myAge = 19

	if myAge > 18 {
		fmt.Println("I am able to play the game")
	} else if myAge < 18 {
		fmt.Println("I am not able to play the game")
	} else {
		fmt.Println("The age is just 18")
	}

	if countAge := 55; countAge > 20 {
		fmt.Println("the age is bigger than 20")
	} else {
		fmt.Println("the age is less then 20")
	}
}
