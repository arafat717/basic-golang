package main

import "fmt"

func main() {
	days := []string{"satarday", "sunday", "wednesday", "thursday", "friday"}
	fmt.Println(days)
	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for key, value := range days {
	// 	fmt.Printf("for key %v , value is %v \n", key, value)
	// }

	myvalue := 1

	for myvalue < 10 {

		if myvalue == 2 {
			goto lco
		}

		if myvalue == 5 {
			myvalue++
			continue
		}

		fmt.Println("My value is here:")

		fmt.Println(myvalue)
		myvalue++
	}

lco:
	fmt.Println("Here we are here!")

}
