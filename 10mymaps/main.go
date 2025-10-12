package main

import "fmt"

func main() {
	fmt.Println("We are in map of go")
	langauage := make(map[string]string)

	langauage["Js"] = "JavaScript"
	langauage["Py"] = "Phython"
	langauage["Jv"] = "Java"

	fmt.Println(langauage)
	fmt.Println(langauage["Js"])
	// delete(langauage, "Jv")
	fmt.Println(langauage)

	for key, value := range langauage {
		fmt.Printf("For key of %v , value is %v \n  ", key, value)
	}

}
