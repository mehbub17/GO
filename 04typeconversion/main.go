package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Rate Us on a scale of 1 to 10")

	reader := bufio.NewReader(os.Stdin)

	rating, _ := reader.ReadString('\n')

	fmt.Println("Thanks for Rating Us", rating)

	// finaleRating := rating + 1 //error

	// rate,err := strconv.ParseFloat(rating,64);// er need to remove the trailing /n from the rating too

	rate, err := strconv.ParseFloat(strings.TrimSpace(rating), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 : ", rate+1)
	}

}
