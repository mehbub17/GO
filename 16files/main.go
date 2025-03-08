package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Handling files in golang")

	content := "This needs to go in the file -- writing in golang"

	file, err := os.Create("./mygofile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	// if err != nil {
	// 	panic(err)
	// }

	checkNilErr(err)

	fmt.Println("file length is ", length)
	defer file.Close() // conventional way like closing the file when the code ends or goes to the end of execution

	//reading
	readFile("./mygofile.txt")
}

func readFile(filename string) {

	databyte, err := ioutil.ReadFile(filename)

	checkNilErr(err) // code reusability

	fmt.Println("Text data inside in file is \n", string(databyte))
}


func checkNilErr(err error){
	if err != nil{
		panic(err)
	}
}