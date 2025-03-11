package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string
	Price    int
	platform string
	Password string
	Tags     []string
}

func main() {

	fmt.Println("making json")
	EncodeJson()
}

func EncodeJson() {

	lcocourses := []course{

		{"reactjs Bootcamp", 200, "code.in", "akjsf", []string{"sfs", "rjs"}},
		{"Mern Bootcamp", 300, "coded.in", "akjsfsfs", []string{"sfffffs", "rrrrrjs"}},
		{"C++ Bootcamp", 400, "codedd.in", "asfskjsf", []string{"sfffs", "rfffjs"}},
	}

	// packaging the data as json data

	finalJson, _ := json.Marshal(lcocourses)

	// fmt.Println(finalJson)
	fmt.Printf("%s\n",finalJson)

	finalJson2,_:=json.MarshalIndent(lcocourses,"","\t")
	fmt.Printf("%s\n",finalJson2)
}
