package database

import (
	"database/sql"
)

func GetOriginalUrlOfShortUrl(short_url string, db *sql.DB) (string, error) {
	var original_url string
	row := db.QueryRow("SELECT original_url FROM urls WHERE short_url = ?", short_url)
	err := row.Scan(&original_url)

	return original_url, err
}