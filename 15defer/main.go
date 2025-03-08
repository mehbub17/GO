package main

import "fmt"

// defer statements gets executed at the end of the code and also gets executed in reverse order they were called  i mean LIFO order 

func main() {

	fmt.Println("Defer in go language 1")

	defer fmt.Println("Defer statement 1")

	fmt.Println("Defer in go language 2")

	defer fmt.Println("Defer statement 2")

	fmt.Println("Defer in go language 3")

	defer fmt.Println("Defer statement 3")

	fmt.Println("Defer in go language 4")

	defer fmt.Println("Defer statement 4")

	myDefer()
}


func myDefer(){
	for i:=0;i<5;i++{
		defer fmt.Println(i);
	}
}