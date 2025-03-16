package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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
	json.NewEncoder(w).Encode(courses)
}
