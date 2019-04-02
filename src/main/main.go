package main

import (
	"fmt"
	"log"
	"net/http"
)

func AllMoviesEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func FindMovieEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func CreateMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func UpdateMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func DeleteMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/movies", AllMoviesEndPoint).Methods("GET")
	mux.HandleFunc("/movies", CreateMovieEndPoint).Methods("POST")
	mux.HandleFunc("/movies", UpdateMovieEndPoint).Methods("PUT")
	mux.HandleFunc("/movies", DeleteMovieEndPoint).Methods("DELETE")
	mux.HandleFunc("/movies/{id}", FindMovieEndpoint).Methods("GET")
	if err := http.ListenAndServe(":3000", mux); err != nil {
		log.Fatal(err)
	}
}