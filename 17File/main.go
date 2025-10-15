package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "This is my third file in golang!"
	file, err := os.Create("./myThirdFile")
	errorMessage(err)
	newFile, err := io.WriteString(file, content)
	errorMessage(err)
	fmt.Println("New length: ", newFile)
	defer file.Close()
	ReadFile("./myThirdFile")
}

func ReadFile(fileName string) {
	strByte, err := ioutil.ReadFile(fileName)
	errorMessage(err)
	fmt.Println(string(strByte))
}

func errorMessage(err error) {
	if err != nil {
		panic(err)
	}
}
