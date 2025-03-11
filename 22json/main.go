package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"` // renames the Name parameter to coursename for the client consuming the api
	Price    int
	platform string
	Password string   `json:"-"`              // wont show password to the user consuming the api
	Tags     []string `json:"tags,omitempty"` //if value null then do show it
}

func main() {

	fmt.Println("making json")
	// EncodeJson()
	DecodeJson()
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
	fmt.Printf("%s\n", finalJson)

	finalJson2, _ := json.MarshalIndent(lcocourses, "", "\t")
	fmt.Printf("%s\n", finalJson2)
}

func DecodeJson() {
	//data from web comes in byte format
	jsonData := []byte(`
	
	    {
			"coursename": "reactjs Bootcamp",
			"Price": 200,
			"tags": ["sfs","rjs"]
        }


	`)


	var lcocourse course

	checkValid :=json.Valid(jsonData)

	if checkValid {
		fmt.Println("valid json");
		json.Unmarshal(jsonData,&lcocourse)
		fmt.Printf("%#v\n",lcocourse) // special format
	}else{
		fmt.Println("Invalid Json")
	}


	//some cases where you just want to add data to key value

	var onlineData map[string]interface{} // when you don't know the type of the value
	json.Unmarshal(jsonData,&onlineData)
	fmt.Printf("%#v\n",lcocourse) 

	for k,v :=range onlineData{
		fmt.Printf("key is %v and value is %v and type is %T\n",k,v,v)
	}


}
