package main

import (
	"github.com/disiqueira/itihost"
	"github.com/disiqueira/itinerary"
	"net/http"
)

func RateHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is my Host!!"))
}

func main() {
	r := itinerary.NewRouter()

	r.NewPath(RateHandler).AddMatcher(itihost.New("disiqueira.itinerary.test"))

	http.ListenAndServe(":8000", r)
}
