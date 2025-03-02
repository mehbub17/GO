package main

import "fmt"



func main()  {
	fmt.Println("Arrays")

	var arr = [6] int{1,2,3,4,5,6}

	fmt.Println(arr)
	fmt.Println(arr[0])
	fmt.Println(arr[0],arr[1],arr[2])

	var brr = [...]int{3,4,5,6,7}
	fmt.Println(len(brr))

	fmt.Println(brr)

	var crr [10]string

	crr[0] = "mango"
	crr[1] = "apple"
	crr[5]="onion" // will give spaces when less size if applied

	fmt.Println(crr)
	fmt.Println(len(crr)) // give the length initially it was assigned 

	


}