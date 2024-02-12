package handlers

import (
	"encoding/json"
	"net/http"

	s "github.com/dwskme/seucy/backend-service/internal/services"
)

type Credentials struct {
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
}

type NewHandler struct {
	UserService  *s.UserService
	TokenService *s.TokenService
	AuthService  *s.AuthService
}

func JsonResponse(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	response := map[string]string{"message": message}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
	}
}
