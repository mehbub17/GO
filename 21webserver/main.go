package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {

	fmt.Println("Web request using get")
	// PerformGetRequest()
	// PerformPostJsonRequest()
	PerformPostFormRequest()

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

	byteCount, _ := ResponseString.Write(content)
	fmt.Println(byteCount) // length of content length
	fmt.Println(ResponseString.String())

	fmt.Println(string(content))

}

func PerformPostJsonRequest() {
	const myurl = "http://localhost:8000/post"

	//fake

	requestBody := strings.NewReader(`
	
	{
		"course":"Go",
		"price":0,
		"platform":"udemy"
	}

	`)

	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	var readBuilder strings.Builder

	content, _ := io.ReadAll(response.Body)

	contentDataByte, _ := readBuilder.Write(content)
	fmt.Println(contentDataByte)
	fmt.Println(readBuilder.String())

}

func PerformPostFormRequest() {
	const myurl = "http://localhost:8000/postform"

	// form data
	data := url.Values{}

	data.Add("firstname", "mehbub")
	data.Add("lastname", "hussain")
	data.Add("email", "hussain@00.com")

	response, err := http.PostForm(myurl, data)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	var ResponseBuilder strings.Builder

	content, _ := io.ReadAll(response.Body)

	length, _ := ResponseBuilder.Write(content)
	fmt.Println(length)

	fmt.Println(ResponseBuilder.String())

}
