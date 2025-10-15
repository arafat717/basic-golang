package main

import (
	"fmt"
	"time"
)

func main() {
	myTime := time.Now()
	fmt.Println("my time:", myTime)
	readAbleTime := myTime.Format("01-02-2006 Monday 15:04:05")
	fmt.Println(readAbleTime)

	createTime := time.Date(2004, time.November, 20, 12, 30, 30, 30, time.Local)
	fmt.Println(createTime.Format("01-02-2006 Monday 15:04:05"))
}
