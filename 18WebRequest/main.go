package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://api.creativelabz.org/api/v1/states"

func main() {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	dataByte, err := ioutil.ReadAll(response.Body)
	content := string(dataByte)
	fmt.Println(content)
}
