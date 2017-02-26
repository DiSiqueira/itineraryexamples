package main

import (
	"github.com/disiqueira/itinerary"
	"log"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GET Hello World"))
}

func main() {
	r := itinerary.NewRouter()
	r.HandleFunc(HomeHandler)
	log.Fatal(http.ListenAndServe(":80", r))
}
