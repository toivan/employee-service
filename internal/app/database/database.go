package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"time"
)

var DB *sql.DB

func Connect() (*sql.DB, error) {
	var err error
	dsn := os.Getenv("DB_DSN")
	maxRetries := 10
	delayBetweenRetries := time.Second * 15
	for i := 0; i < maxRetries; i++ {
		DB, err = sql.Open("mysql", dsn)
		if err != nil {
			log.Printf("Error opening database: %v\n", err)
		} else {
			err = DB.Ping()
			if err == nil {
				log.Println("Successfully connected to the database")
				createTable()
				return DB, nil
			}
			log.Printf("Error pinging database: %v\n", err)
		}

		log.Printf("Retrying in %v seconds...\n", delayBetweenRetries.Seconds())
		time.Sleep(delayBetweenRetries)
	}

	return nil, fmt.Errorf("could not connect to database after %d tries", maxRetries)
}

func createTable() {
	query := `
	CREATE TABLE IF NOT EXISTS employees (
		id INT AUTO_INCREMENT PRIMARY KEY,
		full_name VARCHAR(255),
		phone VARCHAR(20),
		gender VARCHAR(10),
		age INT,
		email VARCHAR(255),
		address VARCHAR(255)
	);`

	if _, err := DB.Exec(query); err != nil {
		log.Fatal(err)
	}
}
