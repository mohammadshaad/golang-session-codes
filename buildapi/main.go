package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// type Course struct {
// 	id     string
// 	name   string
// 	author string
// 	price  int
// }

func main() {
	fmt.Println("Welcome to build api")

	r := mux.NewRouter()

	r.HandleFunc("/", HomeController).Methods("GET")
	r.HandleFunc("/about", AboutController).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}

func HomeController(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Welcome to my home controller")
	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte("Welcome to my home page"))

}

func AboutController(w http.ResponseWriter, r  *http.Request) {
	fmt.Println("This is my about route")
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Welcome to my about page"))

}

// Controllers/Handlers - The function/logic for route
// Route - API endpoint e.g. /home, /about, /services, etc
// models - database structure
