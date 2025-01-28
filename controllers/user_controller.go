package controllers

import (
	"crud/DTOs"
	"crud/config"
	"crud/models"
	"crud/utils"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func (c *UserController) GetUsers(w http.ResponseWriter) {
	db := config.GetDB()
	if db == nil {
		log.Println("Database connection is nil")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	var users []models.User
	err := db.Select(&users, "SELECT id, username, created_at FROM users")
	if err != nil {
		log.Printf("Error fetching users: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer db.Close()
	json.NewEncoder(w).Encode(users)
}

func (c *UserController) GetUser(id int, w http.ResponseWriter) {
	var user models.User
	db := config.GetDB()
	err := db.Get(&user, `SELECT id, username, created_at FROM users where id = ?`, id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "User not found",
			"message": fmt.Sprintf("No user found with ID %d", id),
		})
		return
	}
	defer db.Close()
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
		db := config.GetDB()
		query := `INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`
		_, err := db.Exec(query, userDTO.Username, userDTO.Password, time.Now())
		if err != nil {
			log.Printf("Error creating user: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		defer db.Close()
		json.NewEncoder(w).Encode(userDTO)
	}

}

func (c *UserController) Login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	token, err := c.authenticate(req.Username, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func (r *UserController) findByUsername(username string) (*models.User, error) {
	var user models.User
	db := config.GetDB()
	err := db.Get(&user, "SELECT * FROM users WHERE username = ?", username)
	defer db.Close()
	return &user, err
}

func (c *UserController) authenticate(username, password string) (string, error) {
	user, err := c.findByUsername(username)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.Id,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})
	secret := os.Getenv("JWT_SECRET")
	return token.SignedString([]byte(secret))
}
