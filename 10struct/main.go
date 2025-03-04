package main

import "fmt"

func main() {
	fmt.Println("Structs in go")
	// no inheritance in golang

	userobj := User{"Mehbub", "mehbub@gmail.com", true, 56}
	fmt.Println(userobj)
	fmt.Printf("details is %+v\n",userobj) //+v goes for all the parameters
	fmt.Printf("name is %v\n",userobj.Name)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
