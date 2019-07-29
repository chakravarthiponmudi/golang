package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func h1(w http.ResponseWriter, req *http.Request) {
	// The "/" pattern matches everything, so we need to check
	// that we're at the root here.
	fmt.Println(*req)
	if req.URL.Path != "/" {
		http.NotFound(w, req)
		return
	}
	fmt.Fprintf(w, "Welcome to the home page!")
}

func h2(w http.ResponseWriter, req *http.Request) {
	// The "/" pattern matches everything, so we need to check
	// that we're at the root here.
	fmt.Println(*req)

	fmt.Fprintf(w, "Welcome to the home page! %s", mux.Vars(req))
}

func setupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", h1)
	mux.HandleFunc("/hello/{username}/id", h2)
}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", h1)
	r.HandleFunc("/hello/{username}/id", h2)
	log.Fatal(http.ListenAndServe(":8080", r))

}
