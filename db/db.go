package db

import (
	"database/sql"
	"fmt"
	// "os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() error {
	connStr := "host=localhost port=5432 user=postgres password=test dbname=todo sslmode=disable" //os.Getenv("DBSTRING")
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	err = DB.Ping()
	if err != nil {
		return err
	}
	fmt.Println("Successfully connected to the database!")

	return nil
}