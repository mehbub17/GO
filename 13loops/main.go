package main

import "fmt"

func main(){
	fmt.Println("Loops")

	for i := 0; i < 4; i++{ 
		fmt.Printf(" iteration number %v\n",i)   
	  } 

	// Infinite loop 
    // for { 
	// 	fmt.Printf("1\n")   
	//   }


	// for loop as while loop 
	i:= -3
    for i < 3 { 
       i += 2
	   fmt.Println(i)
    } 
  fmt.Println(i) 


  days := []string{"a","b","c","d","e"}
  fmt.Println(days)

  for i:=0;i<len(days);i++{
	fmt.Println("value is ",days[i]);
  }

  for i := range days{ // giving the index not the values
	fmt.Println(days[i]);
  }

  for index,day := range days{ // gives both days and values
	fmt.Printf("at index %v values is %v\n",index,day)
  }

  for _,day := range days{ // can skip the index
	fmt.Printf("values is %v\n",day)
  }

  fmt.Println(adder(5,6))
  fmt.Println(SpreadAdder(5,6,5,6))
  fmt.Println(SpreadAdder(5,6,5,5,6,6))

  ans,str := SpreadAddermultiReturn(1,2,3,4,5,6,7,8,9)
  ans2,_ := SpreadAddermultiReturn(4,5,6,7,8,9)

  fmt.Println(ans,str);
  fmt.Println(ans2);

}

func adder(val1 int,val2 int) int{
	return val1 + val2;
}

func SpreadAdder(values...int)int{

	ans := 0

	for _,val := range values{
		ans+=val
	}
	return ans
}
func SpreadAddermultiReturn(values...int)(int,string){

	ans := 0

	for _,val := range values{
		ans+=val
	}
	return ans,"I am mehbub"
}