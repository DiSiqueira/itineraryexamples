package main

import (
	"github.com/disiqueira/itiexact"
	"github.com/disiqueira/itinerary"
	"net/http"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is the user route!! "))
}

func main() {
	r := itinerary.NewRouter()

	r.NewPath(UserHandler).AddMatcher(itiexact.New("/user"))

	http.ListenAndServe(":8000", r)
}
