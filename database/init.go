package database

import (
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
	"fmt"
	"log"
)


func InitDb(path string) *sql.DB{
	// fmt.Println("hey")
	// return nil
	db, err := sql.Open("sqlite3", path)
	
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	createTableQuery := `
	CREATE TABLE IF NOT EXISTS urls (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        original_url TEXT NOT NULL,
        short_url TEXT NOT NULL UNIQUE,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`

	db.Exec(createTableQuery);

	return db;
}