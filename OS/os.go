package main

import (
	"github.com/disiqueira/itinerary"
	"github.com/disiqueira/itios"
	"net/http"
)

func LinuxHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("You are using Linux :) "))
}

func main() {
	r := itinerary.NewRouter()

	r.NewPath(LinuxHandler).AddMatcher(itios.New("Linux"))

	http.ListenAndServe(":8000", r)
}
