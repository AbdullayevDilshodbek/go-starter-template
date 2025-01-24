package routes

import (
	"crud/DTOs"
	"crud/controllers"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func UserRoutes(r *mux.Router, db *sqlx.DB) {

	newUserController := controllers.NewUserController(db)

	r.HandleFunc("/user/index", func(w http.ResponseWriter, r *http.Request) {
		newUserController.GetUsers(w)
	}).Methods("GET")

	r.HandleFunc("/user/{id}", func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, _ := strconv.Atoi(params["id"])
		newUserController.GetUser(id, w)
	}).Methods("GET")

	r.HandleFunc("/user/create", func(w http.ResponseWriter, r *http.Request) {
		var userDTO DTOs.CreateUserDTO
		_ = json.NewDecoder(r.Body).Decode(&userDTO)
		newUserController.CreateUser(userDTO, w)
	}).Methods("POST")
}
