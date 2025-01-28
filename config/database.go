package config

import (
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func GetDB() *sqlx.DB {

	// Database connection string
	// dcn := "root:123@(localhost:3306)/golang"
	dcn := fmt.Sprintf("%v:%v@(%v:%v)/%v", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	fmt.Println(dcn)
	// Initialize a mysql database connection
	db, err := sqlx.Connect("mysql", dcn)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(2 * time.Minute)

	// Verify the connection to the database is still alive
	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping the database: %v", err)
	}
	return db
}
