package services

import (
	"errors"
	"fmt"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

type TokenService struct {
	SecretKey []byte
}
type UserClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func NewTokenService(secretkey []byte) *TokenService {
	return &TokenService{SecretKey: secretkey}
}

func (s *TokenService) ExtractTokenFromHeader(bearerToken string) (string, error) {
	if bearerToken == "" {
		return "", errors.New("bearer token not provided")
	}
	tokenParts := strings.Split(bearerToken, " ")
	if len(tokenParts) != 2 {
		return "", errors.New("invalid Bearer token format")
	}
	return tokenParts[1], nil
}

func (s *TokenService) GenerateNewToken(email string) (string, error) {
	claims := UserClaims{
		email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "system",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(s.SecretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (s *TokenService) ValidateToken(tokenstr string) (bool, error) {
	token, err := jwt.ParseWithClaims(tokenstr, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	})
	if err != nil {
		return false, err
	}
	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		fmt.Println(claims.Email, claims.Issuer)
		return true, nil
	}
	return false, errors.New("token is invalid")
}

func (s *TokenService) RefreshToken() {
}

func (s *TokenService) VerifyRefreshToken() {
}
