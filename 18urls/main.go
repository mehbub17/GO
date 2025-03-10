package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=kjsnfbekfek"

func main() {
	fmt.Println("Urls in go")

	result, _ := url.Parse(myUrl)

	fmt.Println(result.Scheme) //https
	fmt.Println(result.Host) //co.dev.3000
	fmt.Println(result.Path) // /learn
	fmt.Println(result.Port()) // 3000
	fmt.Println(result.RawQuery) //  coursename=reactjs&paymentid=kjsnfbekfek basically the params

	qparams := result.Query() // stores all the queryparameters

	fmt.Printf("The type is %T\n",qparams)
	fmt.Println(qparams) // map with key value pair

	fmt.Println(qparams["coursename"])


	for _,val := range qparams{
		fmt.Println("param is",val)
	}

	// making the url

	partsOfUrl := &url.URL{ // passing it as a reference
		Scheme:"https",
		Host:"lco.dev",
		Path:"/toucehs",
		RawQuery:"user=mehbub",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
