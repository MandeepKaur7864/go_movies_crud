package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(movies)
}

func main() {
	r := mux.NewRouter()

	// atleast one movie should be there at the starting
	movie := Movie{
		ID:    "1",
		Isbn:  "438222",
		Title: "Movie One",
		Director: &Director{
			Firstname: "John",
			Lastname:  "Doe",
		},
	}
	movies = append(movies, movie)

	// routes
	r.HandleFunc("/movies", getMovies).Methods("GET")

	fmt.Println("Server starting at port: 8000")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}

}
