package main

import "fmt"

func main() {
	fmt.Println("Hello from fuctions")
	value := adder(7, 3)
	fmt.Println(value)

	result, text := proAdder(1, 2, 3, 4, 5, 6)
	fmt.Println("Our result: ", result)
	fmt.Println(text)
}

func adder(num1 int, num2 int) int {
	return num1 + num2
}

func proAdder(values ...int) (int, string) {
	result := 0
	for _, val := range values {
		result += val
	}
	return result, "hi there "
}
