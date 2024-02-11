package handlers

import (
	"encoding/json"
	"net/http"

	models "github.com/dwskme/seucy/backend-service/internal/models"
	services "github.com/dwskme/seucy/backend-service/internal/services"
)

type Credentials struct {
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
}

type AuthHandler struct {
	UserService  *services.UserService
	TokenService *services.TokenService
	AuthService  *services.AuthService
}

func jsonResponse(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	response := map[string]string{"message": message}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
	}
}

func NewAuthHandler(userService *services.UserService, tokenService *services.TokenService, authService *services.AuthService) *AuthHandler {
	return &AuthHandler{UserService: userService, TokenService: tokenService}
}

func (h *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	var credentials Credentials
	// Decode Input
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		jsonResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	// check if user exists
	userExists, err := h.AuthService.CheckUserExists(credentials.Identifier)
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, "Error checking user existence")
		return
	}
	if !userExists {
		jsonResponse(w, http.StatusUnauthorized, "User does not exist")
		return
	}

	matchPassword, err := h.AuthService.MatchPassword(credentials.Identifier, credentials.Password)
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, "Error checking password")
		return
	}
	if !matchPassword {
		jsonResponse(w, http.StatusUnauthorized, "Incorrect Credentials")
	}
	// TODO: renew the expiry token/ refresh token
}

func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var user *models.User
	// Decode Input
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		jsonResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// check if valid signup request
	validRequest := h.AuthService.CheckValidSignUpRequest(user)
	if !validRequest {
		jsonResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// check valid email or not
	validEmail, msg := h.AuthService.ValidMailAddress(user.Email)
	if !validEmail {
		jsonResponse(w, http.StatusBadRequest, msg)
		return
	}

	// check if email already used
	userExists, err := h.AuthService.CheckUserExists(user.Email)
	if err != nil {
		jsonResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if userExists {
		jsonResponse(w, http.StatusBadRequest, "Email already exists")
		return
	}

	// check if username already used
	userExists, err = h.AuthService.CheckUserExists(user.Username)
	if err != nil {
		jsonResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if userExists {
		jsonResponse(w, http.StatusBadRequest, "Username already exists")
		return
	}

	// create new user in db
	err = h.UserService.CreateUser(*user)
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	// Return success message with 201 Created status
	jsonResponse(w, http.StatusCreated, "User created successfully")
}

func (h *AuthHandler) SignOut(w http.ResponseWriter, r *http.Request) {
	// TODO: clear token /sessions
}
