package main

import (
	"crud/routes"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)
func main() {

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

    r := mux.NewRouter()
    routes.Routes(r)
	r.Use(JsonMiddleware)

    http.ListenAndServe(":3000", r)
}

func JsonMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        next.ServeHTTP(w, r) // Call the next handler in the chain
    })
}
