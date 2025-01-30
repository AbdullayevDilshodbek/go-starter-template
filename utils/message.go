package utils

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, res ErrorResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(res)
}

type ErrorResponse struct {
	Message string `json:"message"`
	Errors  string `json:"errors"`
	Status  int    `json:"status"`
}


		// json.NewEncoder(w).Encode(map[string]interface{}{
		// 	"error":  "invalid credentials",
		// 	"fields": err.Error(),
		// 	"code": http.StatusUnauthorized,
		// })