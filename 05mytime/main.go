package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time Format")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006"))
	fmt.Println(presentTime.Format("01-02-2006 Monday"))
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday")) // only way

	createdDate := time.Date(2020, time.September, 10, 23, 23, 0, 0, time.Local)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))


}
