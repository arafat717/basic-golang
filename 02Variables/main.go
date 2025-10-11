package main

import "fmt"

const AuthToken string = "alksdf;asdkljfha;sdlfj askdfj;asdf"

func main() {
	fmt.Println("Hello from vaiables")

	var MyName string = "Arafat Jibon"
	fmt.Println(MyName)
	fmt.Printf("This is the type of %T \n", MyName)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("This is the type of %T \n", isLoggedIn)

	var TotalPrice uint8 = 255
	fmt.Println(TotalPrice)
	fmt.Printf("This is the type of %T \n", TotalPrice)

	var HisName uint8
	fmt.Println(HisName)
	fmt.Printf("The Type of his name %T \n", HisName)

	// emplcit type

	var token = "a;lsdkfa;sdfha;sdfj a;sldfkj"
	token = "arafat"
	fmt.Println(token)

	//no var
	countryName := "Bangladesh"
	fmt.Println(countryName)

	countryName = "India"
	fmt.Println(countryName)

	fmt.Println("AuthToken:", AuthToken)

}
