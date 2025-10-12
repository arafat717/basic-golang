package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("Welcome to slice in golang!")
	ourValue := 50
	fmt.Printf("The type of ourValue is %T \n", ourValue)

	takeInput := bufio.NewReader(os.Stdin)
	fmt.Println("Give me your result in GPA: ")
	input, _ := takeInput.ReadString('\n')
	fmt.Println("The result of mine is: ", input)

	inputInt, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println("integer input err: ", err)
	} else {
		fmt.Println("The integer value is: ", inputInt+15)
	}

	///// time format

	playtime := time.Now()
	fmt.Println("Play time: ", playtime.Format("01-02-2006 Monday 15:04:05"))

	createTime := time.Date(2025, time.November, 20, 12, 0, 0, 0, time.Local)
	fmt.Println("My BirthDay is in :", createTime.Format("01-02-2006 Monday 15:04:05"))

	/// pointer in go

	fmt.Println("My Age: ", inputInt)
	var ptr = &inputInt
	fmt.Println("My Pointer: ", ptr)
	fmt.Println("My Pointer value is: ", *ptr)

	*ptr = *ptr + 1
	fmt.Println("Next year my age will be", *ptr)

}
