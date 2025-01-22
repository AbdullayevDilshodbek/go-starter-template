package controllers

import (
	"crud/DTOs"
	"crud/config"
	"crud/models"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
)

func GetUsers() []models.User {
	users := []models.User{}
	rows, err := config.DB().Query("select id, username, created_at from users")
	if err != nil {
		panic("Database connection error")
	}
	defer rows.Close()
	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.Id, &user.Username, &user.CreatedAt)
		if err != nil {
			panic(err.Error())
		}
		users = append(users, user)
	}
	return users
}

func GetUser(id int) models.User {
	users := []models.User{}
	for i := 0; i < 5; i++ {
		users = append(users, models.User{
			Id:       i + 1,
			Username: "don",
			Password: "123",
		})
	}

	return users[id]
}

func CreateUser(userDTO *DTOs.CreateUserDTO, w http.ResponseWriter) {

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
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error":  "Validation failed",
			"fields": errorResponse,
		})
		return
	}

	query := `INSERT INTO users (username, password, created_at) VALUES (?,?,?)`
	_, err2 := config.DB().Exec(query, userDTO.Username, userDTO.Password, time.Now())
	if err2 != nil {
		panic(err2.Error())
	}
}
