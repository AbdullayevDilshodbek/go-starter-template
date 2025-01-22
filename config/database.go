package config

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
)

func DB() *sqlx.DB {
	// Database connection string
	dsn := "root:123@(localhost:3306)/golang"

	// Initialize a mysql database connection
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}

	// Verify the connection to the database is still alive
	err = db.Ping()
	if err != nil {
		panic("Failed to ping the database: " + err.Error())
	}
	return db
}
