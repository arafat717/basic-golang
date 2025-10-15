package main

import "fmt"

func main() {
	fmt.Println("Welcome to here ")
	arafat := MyDetails{"Arafat", "arafat@gmail.com", "Sherpur,Mymensingh,Bangladesh", 19, true}
	fmt.Printf("My name is %v. I am %v years of old. I am from %v. This is my email address %v. \n", arafat.Name, arafat.Age, arafat.Address, arafat.Email)
	fmt.Printf("My whole address is %+v:", arafat)
}

type MyDetails struct {
	Name    string
	Email   string
	Address string
	Age     int
	Status  bool
}
