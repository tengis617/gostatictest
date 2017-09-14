package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	env := os.Getenv("APP_ENV")

	if env == "production" {
		log.Println("Running production build")
	} else {
		log.Println("Running dev build")
	}
	log.Println("Listening... a bit more")
	http.ListenAndServe(":8080", nil)
}
