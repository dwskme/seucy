package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB(connectionStr string) (*sql.DB, error) {
	var err error
	DB, err = sql.Open("postgres", connectionStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	err = DB.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	fmt.Println("Successfully connected to the database!")

	err = ensureTableExists()
	if err != nil {
		return nil, fmt.Errorf("failed to ensure table exists: %w", err)
	}
	return DB, nil
}

func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}

func ensureTableExists() error {
	query := `
        CREATE TABLE IF NOT EXISTS users (
            uuid      VARCHAR(36) PRIMARY KEY,
            firstname VARCHAR(255),
            lastname  VARCHAR(255),
            email     VARCHAR(255) UNIQUE,
            username  VARCHAR(255) UNIQUE,
            password  VARCHAR(255),
            role      VARCHAR(20)
        );
    `
	_, err := DB.Exec(query)
	if err != nil {
		fmt.Println("Error creating 'users' table:", err)
	}
	return err
}
