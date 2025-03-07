package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("switch")

	rand.Seed(time.Now().UnixNano())

	diceNum := rand.Intn(6) + 1

	fmt.Println(diceNum)

	switch diceNum {

	case 1:
		fmt.Println("value is 1")
	case 2:
		fmt.Println("value is 2")
	case 3:
		fmt.Println("value is 3")
	case 4:
		fmt.Println("value is 4")
		fallthrough // also run the next switch statement
	case 5:
		fmt.Println("value is 5")
	case 6:
		fmt.Println("value is 6")
	default:
		fmt.Println("value not defined")
	}
}
