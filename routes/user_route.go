package routes

import (
	"crud/DTOs"
	"crud/controllers"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {
	r.HandleFunc("/user/index", func(w http.ResponseWriter, r *http.Request) {
		users := controllers.GetUsers()
		json.NewEncoder(w).Encode(users)
	}).Methods("GET")

	r.HandleFunc("/user/{id}", func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, _ := strconv.Atoi(params["id"])
		user := controllers.GetUser(id)
		json.NewEncoder(w).Encode(user)
	}).Methods("GET")

	r.HandleFunc("/user/create", func(w http.ResponseWriter, r *http.Request) {
		var userDTO DTOs.CreateUserDTO
		_ = json.NewDecoder(r.Body).Decode(&userDTO)
		controllers.CreateUser(&userDTO, w)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(userDTO)
	}).Methods("POST")
}
