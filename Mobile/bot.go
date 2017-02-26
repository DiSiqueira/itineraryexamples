package main

import (
	"github.com/disiqueira/itimobile"
	"github.com/disiqueira/itinerary"
	"net/http"
)

func HeaderHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("You are on a Mobile device! :)"))
}

func main() {
	r := itinerary.NewRouter()

	r.NewPath(HeaderHandler).AddMatcher(itimobile.New(true))

	http.ListenAndServe(":8000", r)
}
