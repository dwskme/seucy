package services

import (
	"database/sql"
	"net/mail"

	models "github.com/dwskme/seucy/backend-service/internal/models"
	crypt "github.com/dwskme/seucy/backend-service/internal/utils/crypt"
)

type AuthService struct {
	DB *sql.DB
}

func NewAuthService(db *sql.DB) *AuthService {
	return &AuthService{DB: db}
}

func (s *AuthService) CheckUserExists(identifier string) (bool, error) {
	var count int
	query := "SELECT COUNT(*) FROM users WHERE username = $1 OR email = $2"
	err := s.DB.QueryRow(query, identifier, identifier).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (s *AuthService) MatchPassword(identifier, password string) (bool, error) {
	var hashedPassword string
	query := "SELECT password FROM users WHERE username = $1 OR email = $2"
	err := s.DB.QueryRow(query, identifier, identifier).Scan(&hashedPassword)
	if err != nil {
		return false, err
	}
	return crypt.CheckPassword(password, hashedPassword) == nil, nil
}

func (s *AuthService) ValidMailAddress(address string) (bool, string) {
	_, err := mail.ParseAddress(address)
	if err != nil {
		return false, "Invalid Email: " + err.Error()
	}
	return true, "Valid Email"
}

func (s *AuthService) ValidateStrongPassword(password string) {
	// TODO:use strong password validation
}

func (s *AuthService) CheckeValidSignInRequest() {
	// TODO:check if signin have non empty mail password ..
}

func (s *AuthService) CheckValidSignUpRequest(user *models.User) bool {
	// TODO:add other validaiton.
	return user != nil && user.Email != "" && user.Username != "" && user.Password != "" && user.Firstname != "" && user.Lastname != ""
}
