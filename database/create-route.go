package database

import (
	"database/sql"
)

func CreateRoute(longURL string, shortURL string, db *sql.DB) error {
	_, err := db.Exec("INSERT INTO urls (original_url, short_url) VALUES (?, ?)", longURL, shortURL);

	return err
}