package main

import (
	"github.com/disiqueira/itimethod"
	"github.com/disiqueira/itinerary"
	"net/http"
)

func GetHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is a GET"))
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is a POST"))
}

func main() {
	r := itinerary.NewRouter()

	r.NewPath(GetHandler).AddMatcher(itimethod.New("GET"))
	r.NewPath(PostHandler).AddMatcher(itimethod.New("POST"))

	http.ListenAndServe(":8000", r)
}
