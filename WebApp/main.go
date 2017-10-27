package main

import (
	"net/http"
	"html/template"
)

type message struct{
	Message string
}

// Name as you wish
func handlerFunc(w http.ResponseWriter, r *http.Request) {

	//fmt.Fprint(w, "<h1>Guessing game</h1>")
	http.ServeFile(w,r,"guess.html")
}
// Name as you wish
func page2(w http.ResponseWriter, r *http.Request) {

	//fmt.Fprint(w, "<h1>Guessing game</h1>")
	//http.ServeFile(w,r,"guess.tmpl")

	onetotwenty :="Guess a number between 1 and 20"

	t, _ := template.ParseFiles("guess.tmpl")

	t.Execute(w, &message{Message:onetotwenty})
}

func main() {
	// Localhost : port 8080
	http.HandleFunc("/", handlerFunc)
	http.HandleFunc("/guess", page2)
	http.ListenAndServe(":8080", nil)

}