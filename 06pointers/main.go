package main

import "fmt"


func main()  {
	fmt.Println("pointer")


	// var ptr *int

	// fmt.Println("value of ptr is ",ptr); // it will give null

	myNum := 27

	var ptr = &myNum
	fmt.Println("value of pointer is ",ptr) //address
	fmt.Println("value of pointer is ",*ptr)//value
	
}