package main

import (
	"net/http"
	"html/template"
	"math/rand"
	"time"
	"strconv"
)

type message struct{
	Message string
}

// random numbers
func random(min, max int) int{
    rand.Seed( time.Now().UnixNano())
    return rand.Intn(max - min) + min
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
	mrand := random(1,20)
	//cookies
	cookie, err := r.Cookie("target");
	if err != nil {
		// Create a cookie instance and set the cookie.
		// You can delete the Expires line (and the time import) to make a session cookie.
		cookie = &http.Cookie{
			Name:    "target",
			Value:   strconv.Itoa(mrand),
			Expires: time.Now().Add(72 * time.Hour),
		}
		http.SetCookie(w, cookie)
	}

	// If we could read it, try to convert its value to an int.
	guess, _:= strconv.Atoi(r.FormValue("guess"))
	cVal, _ := strconv.Atoi(cookie.Value)

	//this checks if the cookie is same as the guess
	if cVal == guess{


	}
	

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