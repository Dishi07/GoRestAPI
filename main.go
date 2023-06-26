package main

import (
	"log"
	"net/http"

	"GoRestAPI/controllers"
)

func main() {
	run()
}

func run() {
	http.HandleFunc("/sign_up", controllers.SignUpHandler)
	http.HandleFunc("/sign_in", controllers.SignInHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
