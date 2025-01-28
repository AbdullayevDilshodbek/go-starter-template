package routes

import (
	"crud/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func Routes(r *mux.Router) {
	r = r.PathPrefix("/api").Subrouter()

	r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		newUserController := controllers.NewUserController()
		newUserController.Login(w, r)
	}).Methods("POST")


	UserRoutes(r)	
}