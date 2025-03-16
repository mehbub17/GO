package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	fmt.Println("Some of Go Mod")
	greeter()

	r:=mux.NewRouter()
	r.HandleFunc("/",serveHome).Methods("GET") // passing reference not function itself
	log.Fatal(http.ListenAndServe(":4000",r))

}

func greeter() {
	fmt.Println("Hello there")
}

// w for writing and sending response to the client
// r is a type of pointer if you want to use the parameters of the request r is used i mean reding the request

func serveHome(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("<h1>Wel come go write series</h1>"))
}