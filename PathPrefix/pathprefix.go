package main

import (
	"github.com/disiqueira/itinerary"
	"github.com/disiqueira/itipathprefix"
	"net/http"
)

func RateHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("I like movies!!"))
}

func main() {
	r := itinerary.NewRouter()

	r.NewPath(RateHandler).AddMatcher(itipathprefix.New("/movies/"))

	http.ListenAndServe(":8000", r)
}
