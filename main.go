package main

import (
	"fmt"
	"log"
	"net/http"
)

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ok!")
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	port := "8000"
	http.HandleFunc("/", index)
	http.HandleFunc("/health", health)
	log.Printf("Running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
