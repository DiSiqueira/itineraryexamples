package main

import (
	"github.com/disiqueira/itibrowser"
	"github.com/disiqueira/itinerary"
	"net/http"
)

func ChromeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("You are using Chrome! :)"))
}

func IEHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("You are using Internet Explorer :( "))
}

func main() {
	r := itinerary.NewRouter()

	r.NewPath(ChromeHandler).AddMatcher(itibrowser.New("Chrome"))
	r.NewPath(IEHandler).AddMatcher(itibrowser.New("Internet Explorer"))

	http.ListenAndServe(":8000", r)
}
