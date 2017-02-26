package main

import (
	"github.com/disiqueira/itinerary"
	"github.com/disiqueira/itiquery"
	"net/http"
)

func HeaderHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("You have the righ Query Values!!!"))
}

func main() {
	r := itinerary.NewRouter()

	r.NewPath(HeaderHandler).AddMatcher(itiquery.New("test", "1234"))

	http.ListenAndServe(":8000", r)
}
