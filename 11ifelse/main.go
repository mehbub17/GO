package main

import "fmt"

func main() {

	fmt.Println("If Else")

	count := 10

	if(count < 20){
		fmt.Println("Less than 20")
	}else{
		fmt.Println("Greater than 20")
	}

	if count == 10{
		fmt.Println("Value is 10")
	}

	if num := 5;num < 10{
		fmt.Println("Number is less than 10")
	}
}
