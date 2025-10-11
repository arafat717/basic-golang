package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Give us a rating in 1 to 5 :")
	input, _ := reader.ReadString('\n')
	rating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	fmt.Println("Our rating:", rating)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Insreased Rating:", rating+1)
	}
}
