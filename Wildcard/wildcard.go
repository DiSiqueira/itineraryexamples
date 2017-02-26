package main

import (
	"github.com/disiqueira/itinerary"
	"github.com/disiqueira/itiwildcard"
	"net/http"
)

func RateHandler(w http.ResponseWriter, r *http.Request) {
	movie := r.URL.Query().Get("itinerary1")
	msg := "Rate " + movie
	w.Write([]byte(msg))
}

func main() {
	r := itinerary.NewRouter()

	r.NewPath(RateHandler).AddMatcher(itiwildcard.New("/movies/*/rate"))

	http.ListenAndServe(":8000", r)
}
