package main

import (
	"html/template"
	"log"
	"net/http"
)

var tmplt *template.Template

func init() {
	tmplt = template.Must(template.ParseFiles("internal/views/index.html"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", ShowHomePage)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}

func ShowHomePage(w http.ResponseWriter, r *http.Request) {
	if err := tmplt.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
