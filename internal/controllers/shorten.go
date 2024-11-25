package controllers

import (
	"database/sql"
	"html/template"
	"net/http"

	"github.com/arijit-gogoi/url-shortener-go/internal/db"
	"github.com/arijit-gogoi/url-shortener-go/internal/url"
)

var (
	data         map[string]string
	shortenTmplt *template.Template
)

func init() {
	shortenTmplt = template.Must(template.ParseFiles("internal/views/shorten.html"))
}

func Shorten(sqlitedb *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		err := db.StoreURL(sqlitedb, shortURL, usersURL)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		data = map[string]string{
			"ShortURL": shortURL,
		}

		// Execute the shorten template with data.
		shortenTmplt.Execute(w, data)
	}
}

func Proxy(sqlitedb *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		shortURL := r.URL.Path[1:]
		if shortURL == "" {
			http.Error(w, "URL Not Provided", http.StatusBadRequest)
			return
		}

		originalURL, err := db.GetOriginalURL(sqlitedb, shortURL)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Redirect(w, r, originalURL, http.StatusPermanentRedirect)
	}
}
