package main

import (
	"github.com/disiqueira/itibot"
	"github.com/disiqueira/itinerary"
	"net/http"
)

func HeaderHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("You are a Bot! Hi Google Crawler :)"))
}

func main() {
	r := itinerary.NewRouter()

	r.NewPath(HeaderHandler).AddMatcher(itibot.New(true))

	http.ListenAndServe(":8000", r)
}
