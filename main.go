package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello from Scratchpad!"))
}

func showPad(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific scratchpad..."))
}

func createPad(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new scratchpad..."))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/pad", showPad)
	mux.HandleFunc("/pad/create", createPad)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
