package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	DB, err := sql.Open("sqlite3", "test.db")

	if err != nil {
		fmt.Println("err", err)
		panic("Could not connect to database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	createTables()
}

func createTables() {
	createEventsTable := `
		CREATE TABLE IF NOT EXISTS events(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			location TEXT NOT NULL,
			dateTime DATETIME NOT NULL,
			user_id INTEGER
		)
	`

	// to execute the above query statement
	_, err := DB.Exec(createEventsTable)
	if err != nil {
		panic("could't create table ")
	}
}
