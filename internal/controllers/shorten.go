package controllers

import (
	"html/template"
	"net/http"
	"strings"
)

var data map[string]string
var shortenTmplt *template.Template

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

	// Sanitise the URL.
	usersURL = sanitiseURL(usersURL)

	// TODO: shorten the URL

	// Map the shortened URL to usersURL.
	data = map[string]string{
		"ShortURL": usersURL,
	}

	// Execute the shorten template with data.
	shortenTmplt.Execute(w, data)
}

func sanitiseURL(usersURL string) string {
	var res string
	if !strings.HasPrefix(usersURL, "http://") && !strings.HasPrefix(usersURL, "https://") {
		res = "https://" + usersURL
	}
	return res
}
