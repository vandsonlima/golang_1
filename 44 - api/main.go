package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/", helloWorld).Methods("GET")
	r.HandleFunc("/{name}", helloWorldWithName).Methods("GET")
	r.HandleFunc("/param/{name}", withParameters).Methods("GET")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		return
	}
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello world!\n")
	if err != nil {
		return
	}
}

func helloWorldWithName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	_, err := fmt.Fprint(w, fmt.Sprintf("Hello world %s", vars["name"]))
	if err != nil {
		return
	}
}

func withParameters(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	qp := r.URL.Query()

	_, err := fmt.Fprint(w, fmt.Sprintf("Hello world %s with paramters %s", vars["name"], qp.Get("id")))
	if err != nil {
		return
	}
}
