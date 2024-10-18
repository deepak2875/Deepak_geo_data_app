package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

var DB *sql.DB

// Connect to the database
func Connect() {
	var err error
	DB, err = sql.Open("postgres", "user=postgres password=Ram123 dbname=geodata_app sslmode=disable")
	if err != nil {
		log.Fatal("Database connection error:", err)
	}
	if err = DB.Ping(); err != nil {
		log.Fatal("Unable to reach the database:", err)
	}
}
