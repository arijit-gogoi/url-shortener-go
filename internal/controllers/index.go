package controllers

import (
	"html/template"
	"net/http"
)

var tmplt *template.Template

func init() {
	tmplt = template.Must(template.ParseFiles("internal/views/index.html"))
}

func ShowIndex(w http.ResponseWriter, r *http.Request) {
	if err := tmplt.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}