package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/arijit-gogoi/url-shortener-go/internal/controllers"
	"github.com/arijit-gogoi/url-shortener-go/internal/db"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	sqlitedb, err := sql.Open("sqlite3", "db.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer sqlitedb.Close()

	err = db.CreateTable(sqlitedb)
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", controllers.ShowIndex)
	mux.HandleFunc("/shorten", controllers.Shorten(sqlitedb))

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
