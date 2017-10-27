package main

import (
	"net/http"
)

// Name as you wish
func handlerFunc(w http.ResponseWriter, r *http.Request) {

	//fmt.Fprint(w, "<h1>Guessing game</h1>")
	http.ServeFile(w,r,"guess.html")
}

func main() {

	// Localhost : port 8080
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":8080", nil)

}