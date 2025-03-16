package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// creating the model

type Course struct {
	CourseId    string  `json:courseid`
	CourseName  string  `json:coursename`
	CoursePrice int     `json:price`
	Author      *Author `json:author`
}

type Author struct {
	Fullname string `json:"fullname`
	Website  string `json:website`
}

// fake database

var courses []Course

//middleware or helper -- separate files generally

func (c *Course) IsEmpty() bool {

	return c.CourseId == "" && c.CourseName == ""
}

func main() {
	fmt.Println("Building an API")
}

// controllers --> different files generally

// serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1> Welcome in API making</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Get all the courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses) // w is the writer and inside encode whatever is json and will be thrown/given to user
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Get one courses")
	w.Header().Set("Content-Type", "application/json")

	// grab the id from request

	params := mux.Vars(r) // params is set of key-val pairs using mux we can just extract them from request

	// loop the course and find matching id and return the response

	for _, course := range courses {

		if course.CourseId == params["id"] { // params is k-v pair
			json.NewEncoder(w).Encode(course.Author)
			return

		}
	}

	json.NewEncoder(w).Encode("<h2>No course found</h2>")
	return

}
