package main

import (
	"crud/routes"
	"log"
	"net/http"
	"os"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"crud/middleware"
)
func main() {
    LoadEnv()
    r := mux.NewRouter()
	r.Use(middleware.JsonMiddleware)
    routes.Routes(r)

    http.ListenAndServe(":" + os.Getenv("SERVER_PORT"), r)
}

func LoadEnv() {
    err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
