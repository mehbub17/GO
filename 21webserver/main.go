package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {

	fmt.Println("Web request using get")
	PerformGetRequest()

}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("status code :", response.StatusCode)
	fmt.Println("Content Length : ", response.ContentLength)

	// content,_ :=ioutil.ReadAll(response.Body)--> old way

	var ResponseString strings.Builder // manually not done using string(content) but using library and more powerful
	content, _ := io.ReadAll(response.Body)
	//content is in byte format


	byteCount,_:=ResponseString.Write(content)
	fmt.Println(byteCount) // length of content length
	fmt.Println(ResponseString.String())

	fmt.Println(string(content))

}
