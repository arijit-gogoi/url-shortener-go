package db

import (
	"database/sql"
)

// CreateTable ensures the URL table exists
func CreateTable(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS urls (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			short_url TEXT NOT NULL,
			original_url TEXT NOT NULL
		); `
	_, err := db.Exec(query)
	return err
}

// StoreURL inserts the new short URL and the original URL into the database.
func StoreURL(db *sql.DB, shortURL, originalURL string) error {
	statement := `INSERT INTO urls (short_url, original_url) VALUES (?, ?)`
	stmnt, err := db.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmnt.Close()
	_, err = stmnt.Exec(shortURL, originalURL)
	return err
}

// GetOriginalURL fetches the original URL using the shortURL string.
func GetOriginalURL(db *sql.DB, shortURL string) (string, error) {
	var originalURL string
	query := `SELECT original_url FROM urls WHERE short_url = ? LIMIT 1`
	err := db.QueryRow(query, shortURL).Scan(&originalURL)
	if err != nil {
		return "", err
	}
	return originalURL, nil
}
