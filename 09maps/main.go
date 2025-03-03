package main

import "fmt"


func main(){
	fmt.Println("Maps");

	lang := make(map[string]string)
	lang["js"]="javascript"
	lang["py"] = "python"

	fmt.Println(lang);
	fmt.Println("JS shorts for",lang["js"])

	//delete

	// delete(lang,"py")
	// fmt.Println(lang)

	for key,val :=range lang{
		fmt.Printf("key %v value is %v\n",key,val);
	}


	// using underscore _ we can skip

	for _,val :=range lang{
		fmt.Printf(" value is %v\n",val);
	}


}