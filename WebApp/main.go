package main

import (
	"fmt"
	"net/http"
)

// Name as you wish
func handlerFunc(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Guessing game")
}

func main() {

	// Localhost : port 8080
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":8080", nil)

}