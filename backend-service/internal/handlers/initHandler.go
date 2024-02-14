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
	MediaService *s.MediaService
}

func JsonResponse(w http.ResponseWriter, status int, message string, responseData interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if status != http.StatusNoContent {
		w.WriteHeader(status)
		response := map[string]interface{}{
			"message": message,
		}
		response["data"] = responseData
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
		}
	} else {
		w.WriteHeader(status)
	}
}
