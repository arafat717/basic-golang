package main

import (
	"fmt"
)

func main() {
	var myArray [3]string
	myArray[0] = "Arafat"
	myArray[1] = "Arafat"
	myArray[2] = "Arafat"
	fmt.Println("My Array: ", len(myArray))

	var newArray = []int{1, 2, 3, 4, 5}
	fmt.Println("New array: ", newArray)
}
