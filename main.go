package main

import (
	"log"
	"net/http"

	"GoRestAPI/controllers/signIn"
	"GoRestAPI/controllers/signUp"
)

func main() {
	run()
}

func run() {
	http.HandleFunc("/sign_up", signUp.SignUpHandler)
	http.HandleFunc("/sign_in", signIn.SignInHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
