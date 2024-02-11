package services

import (
	"database/sql"
	"net/mail"

	crypt "github.com/dwskme/seucy/backend-service/internal/utils/crypt"
	"golang.org/x/crypto/bcrypt"
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

func (s *AuthService) MatchPassword(username, password string) (bool, error) {
	var hashedPassword string
	query := "SELECT password FROM users WHERE username = $1"
	err := s.DB.QueryRow(query, username).Scan(&hashedPassword)
	if err != nil {
		return false, err
	}
	err = crypt.CheckPassword(password, hashedPassword)
	switch err {
	case nil:
		// Passwords match
		return true, nil
	case bcrypt.ErrMismatchedHashAndPassword:
		// Incorrect password
		return false, nil
	default:
		// Other bcrypt-related errors
		return false, err
	}
}

func (s *AuthService) ValidMailAddress(address string) (bool, string) {
	_, err := mail.ParseAddress(address)
	return err != nil, "Invalid Email"
}

func (s *AuthService) ValidateStrongPassword(password string) {
	// TODO:use strong password validation
}

func (s *AuthService) CheckeValidSignInRequest() {
	// TODO:check if signin have non empty mail password ..
}

func (s *AuthService) CheckValidSignUpRequest() {
	// TODO:check if email exist, email and passsword is not samepassword etc...
}
