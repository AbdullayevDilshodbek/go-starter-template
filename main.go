package main

import (
	"crud/routes"
	"net/http"

	"github.com/gorilla/mux"
)
func main() {
    r := mux.NewRouter()
    routes.Routes(r)

    http.ListenAndServe(":3000", r)
}