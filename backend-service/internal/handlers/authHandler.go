package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/dwskme/seucy/backend-service/internal/models"
	services "github.com/dwskme/seucy/backend-service/internal/services"
)

type Credentials struct {
	Identifer string `json:"identifer"`
	Password  string `json:"password"`
}
type AuthHandler struct {
	UserService  *services.UserService
	TokenService *services.TokenService
	AuthService  *services.AuthService
}

func NewAuthHandler(userService *services.UserService, tokenService *services.TokenService, authService *services.AuthService) *AuthHandler {
	return &AuthHandler{UserService: userService, TokenService: tokenService}
}

func (h *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	// TODO: check if any token exists
	token, _ := h.TokenService.ExtractTokenFromHeader(r.Header.Get("Authorization"))
	_, _ = h.TokenService.ValidateToken(token)

	var credentials Credentials
	// Decode Input
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// check if user exists
	userExists, err := h.AuthService.CheckUserExists(credentials.Identifer)
	if err != nil {
		http.Error(w, "Error checking user existence", http.StatusInternalServerError)
		return
	}
	if !userExists {
		http.Error(w, "User does not exist", http.StatusUnauthorized)
		return
	}

	passwordMatch, err := h.AuthService.MatchPassword(credentials.Identifer, credentials.Password)
	if err != nil {
		http.Error(w, "Error checking password", http.StatusInternalServerError)
		return
	}
	if passwordMatch {
		dummyToken := "dummy_access_token"
		w.Header().Set("Content-Type", "application/json")
		response := map[string]string{"access_token": dummyToken}
		json.NewEncoder(w).Encode(response)
	} else {
		http.Error(w, "Incorrect password", http.StatusUnauthorized)
	}
	// TODO: renew the expiry token/ refresh token
}

func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var user *models.User
	// Decode Input
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
	}
	// check valid email or not
	validEmail, msg := h.AuthService.ValidMailAddress(user.Email)
	if !validEmail {
		http.Error(w, msg, http.StatusBadRequest)
		return
	}
	// TODO: validate if all data are filled or not
	// check if user exists
	userExists, err := h.AuthService.CheckUserExists(user.Email)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
	}
	if !userExists {
		http.Error(w, "Email already exists", http.StatusBadRequest)
		return
	}
	// check if username already exists
	userExists, err = h.AuthService.CheckUserExists(user.Username)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
	}
	if !userExists {
		http.Error(w, "Username already exists", http.StatusBadRequest)
		return
	}
	// TODO: redirect to signin after successfull signup
}

func (h *AuthHandler) SignOut(w http.ResponseWriter, r *http.Request) {
	// TODO: clear token /sessions
}
