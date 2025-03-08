package main

import "fmt"

func main() {
	fmt.Println("Structs in go")
	// no inheritance in golang

	userobj := User{"Mehbub", "mehbub@gmail.com", true, 56}
	fmt.Println(userobj)
	fmt.Printf("details is %+v\n",userobj) //+v goes for all the parameters
	fmt.Printf("name is %v\n",userobj.Name)

	userobj.GetStatus()
	userobj.NewMail() // not changing the original one 
	fmt.Printf("name is %v\n",userobj.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}


func (u User) GetStatus(){

	fmt.Println("Is user active : ",u.Status);
}

func (u User) NewMail(){ // usually passing the copy of the object hence change wont reflect to the original object
	u.Email = "test@email.dev"
	fmt.Println("Email of this user is ",u.Email)
}