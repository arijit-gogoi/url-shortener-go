package main

import (
	"log"
	"net/http"

	"github.com/arijit-gogoi/url-shortener-go/internal/controllers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", controllers.ShowIndex)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
