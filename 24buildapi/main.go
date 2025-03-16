package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

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

	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == "" // will generate courseId manually
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

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one courses")
	w.Header().Set("Content-Type", "application/json")

	// what if the body is empty

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	// what about - {}

	var course Course

	_ = json.NewDecoder(r.Body).Decode(&course) // data will be pushed to course as reference

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Please send some data as no data inside json")
		return
	}

	// generate unique id and convert it to string
	// append course into course

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses,course)
	json.NewEncoder(w).Encode(course)
	return

}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Get one courses")
	w.Header().Set("Content-Type", "application/json")

	// grab the id from request

	params := mux.Vars(r) // params is set of key-val pairs using mux we can just extract them from request

	// loop the course and find matching id,remove add with new again return the response
	for index, course := range courses { // we are not in a database hence we are just looping around to search it in slice

		if course.CourseId == params["id"] { // params is k-v pair

			courses = append(courses[:index],courses[index+1:]... )
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course) // decoding the json
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course) // okay updated
			return 

		}
	}

	// send a response when id not found in data base

	json.NewEncoder(w).Encode("Please enter a valid id for update")
	return
}
