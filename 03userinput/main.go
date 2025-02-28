package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	val:="Welcome in user input section"
	fmt.Println(val)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the reading :")
	
	// comma ok syntax or error ok
	
	input, err := reader.ReadString('\n') // read until newline  no try catch in go hence catch the error use ,
	fmt.Println("Thanks for input, ",input);
	fmt.Printf("input  type is  %T \n",input);

	fmt.Printf("error type is  %T \n",err);


}