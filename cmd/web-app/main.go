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

	if err := db.CreateTable(sqlitedb); err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", controllers.ShowIndex)
	mux.HandleFunc("/shorten", controllers.Shorten(sqlitedb))

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
