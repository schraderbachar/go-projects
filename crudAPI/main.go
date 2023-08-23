package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

// every movie will have a director
type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(movies) //encode response from w into json
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //get the id of the resource that we want to remove in our url path
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...) //matching id will append the rest of the data to this spot in the array (it won't exist anymore)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //get the id of the resource that we want to remove in our url path
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item) //return the item
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)   //decode request body and store it as variable 'movie'
	movie.ID = strconv.Itoa(rand.Intn(10000000)) //generate random id

	movies = append(movies, movie) //append new movie to other movies

	json.NewEncoder(w).Encode(movies)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //get the id of the resource that we want to remove in our url path
	for index, item := range movies {
		if item.ID == params["id"] {
			//here we delete the old movie, and make a new one with the same id. Not really what you should do but fine for apps without databases like this one
			movies = append(movies[:index], movies[index+1:]...) //matching id will append the rest of the data to this spot in the array (it won't exist anymore)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			return
		}
	}
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "438227", Title: "Movie one", Director: &Director{Firstname: "John", Lastname: "Doe"}}) //append a few so they are there when start

	movies = append(movies, Movie{ID: "2", Isbn: "45455", Title: "Movie two", Director: &Director{Firstname: "Jane", Lastname: "Doe"}})

	r.HandleFunc("/movies", getMovies).Methods("GET") //when i hit movies route, run getMovies route
	r.HandleFunc("/movie/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
