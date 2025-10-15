package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("This is switch case here!")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Dice Number: ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("You can open your dice now!")
	case 2:
		fmt.Println("You can move 2!")
	case 3:
		fmt.Println("You can move 3!")
	case 4:
		fmt.Println("You can move 4!")
	case 5:
		fmt.Println("You can move 5!")
	case 6:
		fmt.Println("You can move 6!")
	}
}
