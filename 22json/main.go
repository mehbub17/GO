package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"` // renames the Name parameter to coursename for the client consuming the api
	Price    int
	platform string
	Password string `json:"-"` // wont show password to the user consuming the api
	Tags     []string `json:"tags,omitempty"`//if value null then do show it
}

func main() {

	fmt.Println("making json")
	EncodeJson()
}

func EncodeJson() {

	lcocourses := []course{

		{"reactjs Bootcamp", 200, "code.in", "akjsf", []string{"sfs", "rjs"}},
		{"Mern Bootcamp", 300, "coded.in", "akjsfsfs", []string{"sfffffs", "rrrrrjs"}},
		{"C++ Bootcamp", 400, "codedd.in", "asfskjsf", []string{}},
	}

	// packaging the data as json data

	finalJson, _ := json.Marshal(lcocourses)

	// fmt.Println(finalJson)
	fmt.Printf("%s\n",finalJson)

	finalJson2,_:=json.MarshalIndent(lcocourses,"","\t")
	fmt.Printf("%s\n",finalJson2)
}
