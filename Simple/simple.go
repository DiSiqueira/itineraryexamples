package main

import (
	"github.com/disiqueira/itinerary"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GET Hello World"))
}

func main() {
	r := itinerary.NewRouter()
	r.NewPath(HomeHandler)
	http.ListenAndServe(":8000", r)
}
