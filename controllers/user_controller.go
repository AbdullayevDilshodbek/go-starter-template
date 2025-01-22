package controllers

import (
	"crud/DTOs"
	"crud/config"
	"crud/models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
)

type UserController struct {
	db *sqlx.DB
}

func NewUserController() *UserController {
	db := config.DB()
	return &UserController{db: db}
}

func (c *UserController) GetUsers(w http.ResponseWriter) {
	var users []models.User
	err := c.db.Select(&users, "SELECT id, username, created_at FROM users")
	if err != nil {
		panic(err.Error())
	}
	json.NewEncoder(w).Encode(users)
}

func (c *UserController) GetUser(id int, w http.ResponseWriter) {
	var user models.User
	err := c.db.Get(&user, `SELECT id, username, created_at FROM users where id = ?`, id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "User not found",
			"message": fmt.Sprintf("No user found with ID %d", id),
		})
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (c *UserController) CreateUser(userDTO DTOs.CreateUserDTO, w http.ResponseWriter) {

	validate := validator.New()
	err := validate.Struct(userDTO)
	if err != nil {
		// Validation failed, handle the error
		validationErrors := err.(validator.ValidationErrors)

		// Prepare the error response
		errorResponse := make(map[string]string)
		for _, fieldError := range validationErrors {
			errorResponse[fieldError.Field()] = fieldError.Tag()
		}

		// Send the JSON response with validation errors
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error":  "Validation failed",
			"fields": errorResponse,
		})
	} else {
		query := `INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`
		_, err := c.db.Exec(query, userDTO.Username, userDTO.Password, time.Now())
		if err != nil {
			panic(err.Error())
		}
		json.NewEncoder(w).Encode(userDTO)
	}

}
