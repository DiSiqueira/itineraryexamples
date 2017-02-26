package main

import (
	"github.com/disiqueira/itibasicauth"
	"github.com/disiqueira/itinerary"
	"net/http"
)

func NoAuthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is a UnAuthenticated Request"))
}

func BasicAuthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is a Authenticated Request!!!"))
}

func main() {
	r := itinerary.NewRouter()

	r.NewPath(BasicAuthHandler).AddMatcher(itibasicauth.New("disiqueira", "123"))
	r.NewPath(NoAuthHandler)

	http.ListenAndServe(":8000", r)
}
