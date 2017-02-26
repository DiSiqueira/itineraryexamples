package main

import (
	"github.com/disiqueira/itihost"
	"github.com/disiqueira/itinerary"
	"net/http"
)

func HttpsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("All Requests with HTTPS"))
}

func HttpHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("All Requests with HTTP"))
}

func main() {
	r := itinerary.NewRouter()

	r.NewPath(HttpsHandler).AddMatcher(itihost.New("https"))
	r.NewPath(HttpHandler).AddMatcher(itihost.New("http"))

	http.ListenAndServe(":8000", r)
}
