package main

import (
	"crud/routes"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)
func main() {
    LoadEnv()
    r := mux.NewRouter()
    routes.Routes(r)
	r.Use(JsonMiddleware)

    http.ListenAndServe(":" + os.Getenv("SERVER_PORT"), r)
}

func JsonMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        next.ServeHTTP(w, r) // Call the next handler in the chain
    })
}

func LoadEnv() {
    err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
