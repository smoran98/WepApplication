// GO Web App
// Shane Moran
// G00338607
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
	Guess int
	check bool
	PrintMessage string
}

// Random numbers
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
	
	// 5.cookies
	cookie, err := r.Cookie("target");
	if err != nil {
		// Create a cookie instance & set the cookie.
		// Can delete the Expires line (and the time import) to make a session cookie.
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

	onetotwenty :="Guess a number between 1 and 20"

	mss := &message{Message:onetotwenty,Guess: guess, check : false }

	// Checks if Cookie is same as guess
	if cVal == guess{

		cookie = &http.Cookie{
			Name:    "target",
			Value:   strconv.Itoa(mrand),
			Expires: time.Now().Add(72 * time.Hour),
		}
		http.SetCookie(w, cookie)

		mss.PrintMessage = "You have won"
		mss.check = true;

	}else if guess < cVal {
			mss.PrintMessage = "Too loww"
	}else{
		mss.PrintMessage = "Too Highh"
	}

	t, _ := template.ParseFiles("guess.tmpl")

	t.Execute(w, mss)
}

func main() {
	// Localhost : port 8080
	http.HandleFunc("/", handlerFunc)
	http.HandleFunc("/guess", page2)
	http.ListenAndServe(":8080", nil)

}