package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	ourTExt := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your name:")
	text, _ := ourTExt.ReadString('\n')
	fmt.Println("Welcome Back:", text)
}
