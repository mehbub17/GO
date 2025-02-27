package main

import "fmt"


const Value1 string = "sdhfbksj" // Capital so this is public value can be accessed from other module too in go we don't write export public we just do it making var name starting with Capital Letter

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

	var anotherVar  int // by default 0
	fmt.Println(anotherVar)
	fmt.Printf("smallFloat is of type : %T \n",anotherVar);

	
	// implicit type

	var site = "code.org"  // lexer assigns string type based on the type of the value and later on it will only be allowed as string 
	fmt.Println(site)
	// site = 3 will give an error

	
	// No var style
	//Inside a method this is allowed only // not globally used
	numberOne := 3000 // lexer comes again 
	fmt.Println(numberOne)


	
	fmt.Println(Value1)



}