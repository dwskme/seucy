package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	// TODO:Change connectionString to be dynamic
	connectionString := "postgres://root:root@localhost:5432/seucydb?sslmode=disable"
	var err error
	DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	err = DB.Ping()
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		os.Exit(1)
	}
	fmt.Println("Successfully connected to the database!")
	err = ensureTableExists()
	if err != nil {
		panic(err)
	}
}

func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}

func ensureTableExists() error {
	query := `
		CREATE TABLE IF NOT EXISTS users (
			uuid VARCHAR(36) PRIMARY KEY,
			firstname VARCHAR(255),
			lastname VARCHAR(255),
			email VARCHAR(255),
			password VARCHAR(255),
			role VARCHAR(50)
		)
	`
	_, err := DB.Exec(query)
	return err
}
