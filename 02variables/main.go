package main

import "fmt"

func main(){
	
	var userName string = "Mehbub";
	fmt.Println(userName);
	fmt.Printf("userName is of type : %T \n",userName);

	var isLoggedIn bool= true;
	fmt.Println(isLoggedIn);
	fmt.Printf("isLoggedIn is of type : %T \n",isLoggedIn);

	var smallVal uint8 = 255; //limit is 8 bit or error
	fmt.Println(smallVal);
	fmt.Printf("smallVal is of type : %T \n",smallVal);


	var smallFloat float32 = 255.;// 5 decimal places// float66 is also there and more precise
	fmt.Println(smallFloat);
	fmt.Printf("smallFloat is of type : %T \n",smallFloat);


	//defaults and aliases

}