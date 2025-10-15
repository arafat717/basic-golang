package main

import "fmt"

// pointer (*ptr) gives us the read value of the referance
func main() {
	mainValue := 40
	fmt.Println("Main Value: ", mainValue)
	var ptr = &mainValue
	fmt.Println("prt value : ", ptr)
	fmt.Println("prt value : ", *ptr)
	*ptr = *ptr + 10
	fmt.Println("New value: ", *ptr)
}
