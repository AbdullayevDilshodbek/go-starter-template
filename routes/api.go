package routes

import (
	"crud/controllers"
	"crud/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func Routes(r *mux.Router) {

	r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		newUserController := controllers.NewUserController()
		newUserController.Login(w, r)
	}).Methods("POST")

	apiRoute := r.PathPrefix("/api/v1").Subrouter()
	apiRoute.Use(middleware.AuthMiddleware)
	UserRoutes(apiRoute)
}