package main

import (
	"fmt"
	"sort"
)

func main(){
	fmt.Println("Slices");

	var flist = []string{"apple","mango","bananan"}

	fmt.Printf("type of flist is %T\n",flist)

	flist = append(flist,"coconut","sugarcane")

	fmt.Println(flist)

	var newList = flist[1:3]//start include end exclude
	fmt.Println(newList)


	newList = flist[1:]
	fmt.Println(newList)


	newList = append(newList,newList[1])
	fmt.Println(newList)
	fmt.Println(newList[:3])



	highScore := make([]int,4) // also allocate memory and initialize it with 0
	fmt.Printf("type %T\n",highScore)

	highScore[0]=677;
	highScore[1]=347;
	highScore[2]=277;
	highScore[3]=2277;
	
	fmt.Println(highScore)
	// highScore[4]=2277;// error as out of bound

	highScore = append(highScore, 33,22,444);
	fmt.Println(highScore)

	// as soon as append is used reallocation of memory happens

	sort.Ints(highScore)
	fmt.Println(highScore)
	fmt.Println(sort.IntsAreSorted(highScore))//slice is sorted

	// Removing values from slices based on index

	var courses = []string {"react","js","go","cpp","java"}
	fmt.Println(courses)
	
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...); // slice expect individual element hence ... for spreading

	fmt.Println(courses)

}