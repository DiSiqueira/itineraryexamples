package main

import (
	"github.com/disiqueira/itiheader"
	"github.com/disiqueira/itinerary"
	"net/http"
)

func HeaderHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("You came from my Github Page!"))
}

func main() {
	r := itinerary.NewRouter()

	r.NewPath(HeaderHandler).AddMatcher(itiheader.New("Referer", "https://www.github.com/Disiqueira"))

	http.ListenAndServe(":8000", r)
}
