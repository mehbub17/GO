package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://google.com"

func main() {
	fmt.Println("web request in go language")

	response, err := http.Get(url) // not a copy of the response getting the original response such that we can manipulate the response

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response type %T\n",response)
	fmt.Println(response)

	defer response.Body.Close() // to be done by the requesting user ... mandatory and deferring such that it will run at the end of the code execution // close the connection at the end

	databytes,err :=ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	content := string(databytes) // getting the html skeleton of the page

	fmt.Println(content)

}
