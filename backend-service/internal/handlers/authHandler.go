package handlers

import (
	"encoding/json"
	"net/http"

	models "github.com/dwskme/seucy/backend-service/internal/models"
)

func (h *NewHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	var credentials Credentials
	// Decode Input
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		JsonResponse(w, http.StatusBadRequest, "Invalid request payload", nil)
		return
	}
	// check if user exists
	userExists, err := h.AuthService.CheckUserExists(credentials.Identifier)
	if err != nil {
		JsonResponse(w, http.StatusInternalServerError, "Error checking user existence", nil)
		return
	}
	if !userExists {
		JsonResponse(w, http.StatusUnauthorized, "User does not exist", nil)
		return
	}

	matchPassword, err := h.AuthService.MatchPassword(credentials.Identifier, credentials.Password)
	if err != nil {
		JsonResponse(w, http.StatusInternalServerError, "Error checking password", nil)
		return
	}
	if !matchPassword {
		JsonResponse(w, http.StatusUnauthorized, "Incorrect Credentials", nil)
	}
	// TODO: renew the expiry token/ refresh token
}

func (h *NewHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var user *models.User
	// Decode Input
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		JsonResponse(w, http.StatusBadRequest, "Invalid request payload", nil)
		return
	}

	// check if valid signup request
	validRequest := h.AuthService.CheckValidSignUpRequest(user)
	if !validRequest {
		JsonResponse(w, http.StatusBadRequest, "Invalid request payload", nil)
		return
	}

	// check valid email or not
	validEmail, msg := h.AuthService.ValidMailAddress(user.Email)
	if !validEmail {
		JsonResponse(w, http.StatusBadRequest, msg, nil)
		return
	}

	// check if email already used
	userExists, err := h.AuthService.CheckUserExists(user.Email)
	if err != nil {
		JsonResponse(w, http.StatusBadRequest, "Invalid request payload", nil)
		return
	}
	if userExists {
		JsonResponse(w, http.StatusBadRequest, "Email already exists", nil)
		return
	}

	// check if username already used
	userExists, err = h.AuthService.CheckUserExists(user.Username)
	if err != nil {
		JsonResponse(w, http.StatusBadRequest, "Invalid request payload", nil)
		return
	}
	if userExists {
		JsonResponse(w, http.StatusBadRequest, "Username already exists", nil)
		return
	}

	// create new user in db
	err = h.UserService.CreateUser(*user)
	if err != nil {
		JsonResponse(w, http.StatusInternalServerError, "Internal Server Error", nil)
		return
	}

	// Return success message with 201 Created status
	JsonResponse(w, http.StatusCreated, "User created successfully", nil)
}

func (h *NewHandler) Refresh(w http.ResponseWriter, r *http.Request) {
}

func (h *NewHandler) SignOut(w http.ResponseWriter, r *http.Request) {
	h.TokenService.ClearToken(w)
}
