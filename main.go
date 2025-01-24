package main

import (
	"crud/routes"
	"net/http"
	"crud/config"
	"github.com/gorilla/mux"
)
func main() {

	db := config.DB()

    r := mux.NewRouter()
    routes.Routes(r, db)
	r.Use(JsonMiddleware)

    http.ListenAndServe(":3000", r)
}

func JsonMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        next.ServeHTTP(w, r) // Call the next handler in the chain
    })
}
