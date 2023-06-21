package main

import (
	"log"
	"net/http"

	signUp "GoRestAPI/controllers"
)

func main() {
	run()
}

func run() {
	http.HandleFunc("/sign_up", signUp.SignUpHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
