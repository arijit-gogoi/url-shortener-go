package controllers

import (
	"html/template"
	"net/http"

	"github.com/arijit-gogoi/url-shortener-go/internal/url"
)

var (
	data         map[string]string
	shortenTmplt *template.Template
)

func init() {
	shortenTmplt = template.Must(template.ParseFiles("internal/views/shorten.html"))
}

func Shorten(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowd", http.StatusMethodNotAllowed)
		return
	}

	// userURL is the URL that the client wants to shorten.
	usersURL := r.FormValue("url")
	if usersURL == "" {
		http.Error(w, "URL not provided", http.StatusBadRequest)
		return
	}

	usersURL = url.Sanitise(usersURL)
	shortURL := url.Shorten(usersURL)

	data = map[string]string{
		"ShortURL": shortURL,
	}

	// Execute the shorten template with data.
	shortenTmplt.Execute(w, data)
}
